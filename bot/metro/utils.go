package metro

import "github.com/DiogoSantoss/kant-bot/bot/discord"

func colorForLine(line string) int {
	switch line {
	case "Amarela":
		return discord.Yellow
	case "Azul":
		return discord.Blue
	case "Verde":
		return discord.Green
	case "Vermelha":
		return discord.Red
	default:
		return discord.White
	}
}
