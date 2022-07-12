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

// Global Variable
var PagedEmbeds map[string]*PagedEmbed = make(map[string]*PagedEmbed)

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

func DeleteTimeout() {
	for _, embeds := range PagedEmbeds {
		if time.Since(embeds.Message.Timestamp) > 15*time.Minute {
			delete(PagedEmbeds, embeds.Message.ID)
		}
	}
}

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

func (p *PagedEmbed) NextPage() {
	p.Idx = (p.Idx + 1) % len(p.Pages)
}

func (p *PagedEmbed) PreviousPage() {
	if p.Idx == 0 {
		p.Idx = len(p.Pages) - 1
	} else {
		p.Idx = (p.Idx - 1) % len(p.Pages)
	}
}

func (p *PagedEmbed) UpdatePage() {
	p.Session.ChannelMessageEditEmbed(p.Message.ChannelID, p.Message.ID, p.Pages[p.Idx])
}
