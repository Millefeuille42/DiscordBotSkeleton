package guild

import (
	"fmt"
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
	"github.com/Millefeuille42/DiscordBotSkeleton/utils"
	"strings"
)

// AdminSendSettings Send a guild settings, admin rights are not required for this
func AdminSendSettings(agent sessionAgent.DiscordAgent) {
	if !GuildInitialCheck(agent) {
		return
	}
	settings, err := GuildLoadFile(agent, false, "")
	if err != nil {
		return
	}

	message := fmt.Sprintf(
		"```\n" +
			"Admins:           ",
	)

	for i, admin := range settings.Admins {
		if i == len(settings.Admins)-1 {
			message = fmt.Sprintf("%s@%s\n```", message, utils.GetUser(agent.Session, admin))
			break
		}
		message = fmt.Sprintf("%s@%s, ", message, utils.GetUser(agent.Session, admin))
	}
	utils.SendMessageWithMention(message, "", agent)
}

// AdminSet Add provided admins to the guild
func AdminSet(agent sessionAgent.DiscordAgent) {
	if !GuildInitialCheck(agent) {
		return
	}
	args := agent.Args
	if len(args) <= 1 {
		return
	}
	data, err := GuildLoadFile(agent, false, "")
	if err != nil {
		return
	}

	if !utils.Find(data.Admins, agent.Message.Author.ID) {
		err = utils.SendMessageWrapper(agent.Session, agent.Channel, "You are not an admin")
		return
	}

	for _, user := range args[1:] {
		if !strings.Contains(user, "!") {
			continue
		}
		user = strings.TrimSpace(user)
		user = user[3 : len(user)-1]
		if !utils.Find(data.Admins, user) {
			data.Admins = append(data.Admins, user)
			discordRoleSet(data, user, "admin", agent)
		}
	}
	if guildWriteFile(agent, data) == nil {
		utils.SendMessageWithMention("Successfully added user(s) as admin", "", agent)
	}
}
