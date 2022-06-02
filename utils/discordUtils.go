package utils

import (
	"fmt"
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func SendMessageWrapper(session *discordgo.Session, channel, message string) error {
	if len(message) > 1950 {
		_, err := session.ChannelFileSend(channel, "text.txt", strings.NewReader(message))
		return err
	}
	_, err := session.ChannelMessageSend(channel, message)
	return err
}

// SendMessageWithMention Sends a discord message prepending a mention + \n to the message, if id == "", id becomes the message author
func SendMessageWithMention(message, id string, agent sessionAgent.DiscordAgent) {
	var err error

	if len(message) > 1950 {
		_, err = agent.Session.ChannelFileSend(agent.Channel, "text.txt", strings.NewReader(message))
	}

	if agent.Message != nil && agent.Message.ChannelID == agent.Channel {
		_, err = agent.Session.ChannelMessageSendReply(agent.Channel, message, agent.Message.Reference())
	} else {
		if id == "" {
			id = agent.Message.Author.ID
		}
		err = SendMessageWrapper(agent.Session, agent.Channel, fmt.Sprintf("<@%s>\n%s", id, message))
	}

	if err != nil {
		LogError(err)
	}
}

// GetUser Returns associated user of provided id
func GetUser(session *discordgo.Session, id string) string {
	ret, err := session.User(id)
	if err != nil {
		return ""
	}
	return ret.Username
}

// GetChannelName Returns associated channel name of provided id
func GetChannelName(session *discordgo.Session, id string) string {
	ret, _ := session.Channel(id)
	return ret.Name
}

// LogErrorToChan Sends plain error to command channel
func LogErrorToChan(agent sessionAgent.DiscordAgent, err error) {
	if err == nil {
		return
	}
	LogError(err)
	_ = SendMessageWrapper(agent.Session, agent.Channel,
		fmt.Sprintf("An Error Occured, Please Try Again Later {%s}", err.Error()))
}
