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
	err := config.Setup()
	if err != nil {
		fmt.Println("Failed to start bot.")
		return
	}

	// Create and configure a new Discord session
	dS, err := discord.Setup()
	if err != nil {
		fmt.Println("Failed to start bot.")
		return
	}
	
	handlers.Setup(dS)
	err = discord.Start(dS)
	if err != nil {
		fmt.Println("Failed to start bot.")
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Stop(dS)
}
