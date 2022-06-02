package github

import (
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
	"github.com/Millefeuille42/DiscordBotSkeleton/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

.com/Millefeuille42/DiscordBotSkeleton

import (
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
	"github.com/Millefeuille42/DiscordBotSkeleton/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type commandHandler func(agent sessionAgent.DiscordAgent)

func commandRouter(agent sessionAgent.DiscordAgent) {
	agent.Content = strings.Replace(agent.Message.Content, mainBot.Prefix, "", 1)
	splitBuffer := utils.CleanSplit(agent.Content, ' ')
	if len(splitBuffer) < 1 {
		return
	}
	agent.Command = splitBuffer[0]
	agent.Args = splitBuffer
	if fc, ok := mainBot.CommandMap[agent.Command]; ok {
		fc(agent)
	}
}

// messageHandler Discord bot message handler
func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	botID, _ := session.User("@me")
	agent := sessionAgent.DiscordAgent{
		Session: session,
		Message: message,
	}
	agent.Channel = agent.Message.ChannelID

	if message.Author.ID == botID.ID || !strings.HasPrefix(message.Content, mainBot.Prefix) {
		return
	}
	commandRouter(agent)
}
