package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type PagedEmbed struct {
	Idx     int
	Pages   []*discordgo.MessageEmbed
	Message *discordgo.Message
	Session *discordgo.Session
}

// TODO refactor
// This uses a global hash of embeds which is not ideal

// Global Variable to store all paginated embeds
var PagedEmbeds map[string]*PagedEmbed = make(map[string]*PagedEmbed)

// Create a paginated embed and save to global hash
func CreatePageEmbed(s *discordgo.Session, pages []*discordgo.MessageEmbed, m *discordgo.Message) {
	PagedEmbeds[m.ID] = &PagedEmbed{
		Idx:     0,
		Pages:   pages,
		Message: m,
		Session: s,
	}

	s.MessageReactionAdd(m.ChannelID, m.ID, NavLeft)
	s.MessageReactionAdd(m.ChannelID, m.ID, NavRight)
}

// Delete old paginated embeds
func DeleteTimeout() {
	for _, embeds := range PagedEmbeds {
		if time.Since(embeds.Message.Timestamp) > 15*time.Minute {
			delete(PagedEmbeds, embeds.Message.ID)
		}
	}
}

// Switch to previous/next page based on reaction
func (p *PagedEmbed) SwitchPage(r *discordgo.MessageReactionAdd) {
	switch r.Emoji.Name {
	case NavRight:
		p.NextPage()
	case NavLeft:
		p.PreviousPage()
	}
	p.UpdatePage()

	p.Session.MessageReactionRemove(p.Message.ChannelID, p.Message.ID, r.Emoji.Name, r.UserID)
}

// Switch to next page
func (p *PagedEmbed) NextPage() {
	p.Idx = (p.Idx + 1) % len(p.Pages)
}

// Switch to previous page
func (p *PagedEmbed) PreviousPage() {
	if p.Idx == 0 {
		p.Idx = len(p.Pages) - 1
	} else {
		p.Idx = (p.Idx - 1) % len(p.Pages)
	}
}

// Update embed with current page
func (p *PagedEmbed) UpdatePage() {
	p.Session.ChannelMessageEditEmbed(p.Message.ChannelID, p.Message.ID, p.Pages[p.Idx])
}
