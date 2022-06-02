package guild

import (
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
)

func discordRoleSetLoad(id, role string, agent sessionAgent.DiscordAgent) error {
	roleId := "none"

	data, err := guildLoadFile(agent, false, "")
	if err != nil {
		return err
	}
	if id == "" {
		id = agent.Message.Author.ID
	}

	if role == "admin" {
		roleId = data.AdminRole
	}

	if roleId != "none" {
		_ = agent.Session.GuildMemberRoleAdd(agent.Message.GuildID, id, roleId)
	}
	return nil
}

func discordRoleSet(data GuildData, id, role string, agent sessionAgent.DiscordAgent) {
	roleId := "none"

	if id == "" {
		id = agent.Message.Author.ID
	}

	if role == "admin" {
		roleId = data.AdminRole
	}

	if roleId != "none" {
		_ = agent.Session.GuildMemberRoleAdd(agent.Message.GuildID, id, roleId)
	}
}
