package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/DiogoSantoss/kant-bot/bot/config"
	"github.com/DiogoSantoss/kant-bot/bot/discord"
	"github.com/DiogoSantoss/kant-bot/bot/handlers"
)

func main() {

	// Setup env variables and dependency injection
	config.Setup()

	// Create and configure a new Discord session
	dS := discord.Setup()
	handlers.Setup(dS)
	discord.Start(dS)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Stop(dS)
}
