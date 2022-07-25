# Kant

<img align="right" alt="DiscordGo logo" src="docs/img/kant.jpg" width="400">

## About
This is, as you might have guessed it, a discord bot written in Go.  
During high school, I studied philosophy and really enjoyed the ideias of [Immanuel Kant](https://en.wikipedia.org/wiki/Immanuel_Kant), the famous philosopher.  
So I decided to name this project after him.  
I'm doing this mostly to pratice Golang and as a fun side effect recall some of Kant's greatest ideas.

## Features
It's still in the early stages of development so I'm not too sure what I'll do with it, some ideias are:
- Integration with [Fenix API](https://fenixedu.org/dev/api/)
- Integration with [Metro Lisboa API](https://api.metrolisboa.pt/store/)
- Some Kant lore
- Containerize with Docker
- Switch to gRPC to communicate between the bot and the services
- CI/CD with Github Actions
- Host somewhere (maybe Heroku)

## Architecture

I wanted to take this opportunity to learn more about microservices and therefore this architecture is centered around that.  
The bot is used as a frontend for the users in Discord while the business logic is implemented in the backend with multiple microservices.

If you notice any anti-patterns or bad practices, please send me a message or a PR.

## Services File Structure

### Kant
```
.
├── main.go
├── go.mod
├── go.sum
├── config
│   └── config.go
├── discord
│   ├── colors.go
│   ├── emojis.go
│   ├── pagedEmbed.go
│   └── setup.go
├── handlers
│   ├── help.go
│   ├── metro.go
│   └── setup.go
└── metro
    ├── lines.go
    ├── stations.go
    ├── times.go
    └── utils.go
```

### Metro Lisboa
```
.
├── main.go
├── go.mod
├── go.sum
├── handlers
│   ├── handlers.go
│   └── parsers.go
├── metro
│   └── destinations.go
└── server
    └── server.go
```

## Invite link
Use this [link](https://discord.com/oauth2/authorize?client_id=994381773909803050&permissions=8&scope=bot) to invite Kant to your server.

## Installation
Coming soon ...

## Build with
- [discordgo](https://github.com/bwmarrin/discordgo) - bindings for discord api
- [godotenv](github.com/joho/godotenv) - load env variables from a file

## Questions/Suggestions/Bug Reports
Feel free to message me on discord or open an issue/PR on github.
