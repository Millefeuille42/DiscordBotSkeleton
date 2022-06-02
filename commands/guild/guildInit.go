package guild

import (
	"fmt"
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
	"github.com/Millefeuille42/DiscordBotSkeleton/utils"
	"github.com/bwmarrin/discordgo"
	"os"
)

// getOrCreateRole Internal, Creates or get appropriate role
func getOrCreateRole(name string, roles *[]*discordgo.Role, agent sessionAgent.DiscordAgent) (*discordgo.Role, error) {
	var role *discordgo.Role
	skip := false

	for _, rl := range *roles {
		if rl.Name == name {
			skip = true
			role = rl
			break
		}
	}
	if !skip {
		role, err := agent.Session.GuildRoleCreate(agent.Message.GuildID)
		if err != nil {
			return nil, err
		}
		role, err = agent.Session.GuildRoleEdit(
			agent.Message.GuildID, role.ID,
			name, role.Color, false, role.Permissions, true,
		)
		if err != nil {
			return nil, err
		}
	}

	if role == nil {
		return nil, os.ErrInvalid
	}
	return role, nil
}

// createRoles Internal, Creates or get appropriate roles, and associate them to data
func createRoles(agent sessionAgent.DiscordAgent, data *GuildData) error {
	names := []string{
		"SegBot - Admin",
	}
	roles, err := agent.Session.GuildRoles(agent.Message.GuildID) // Set roles list here so not queried every time
	checkRoles := err == nil

	for _, name := range names {
		var role *discordgo.Role

		if checkRoles {
			role, err = getOrCreateRole(name, &roles, agent) // Pass roles as pointer reason is, as above
			if err != nil {
				return err
			}
		}
		switch name {
		case "SegBot - Admin":
			data.AdminRole = role.ID
		}
	}
	return nil
}

// createData Internal, creates and returns data file
func createData(agent sessionAgent.DiscordAgent) GuildData {
	data := GuildData{
		GuildID:   agent.Message.GuildID,
		Admins:    append(make([]string, 0), agent.Message.Author.ID),
		AdminRole: "None",
	}
	if createRoles(agent, &data) != nil {
		_ = utils.SendMessageWrapper(agent.Session, agent.Channel,
			"Failed to create roles, you'll have to create and configure the missing ones")
	}
	return data
}

// writeData Internal, checks if guild registered, if not registers guild
func writeData(agent sessionAgent.DiscordAgent, data GuildData) error {
	path := fmt.Sprintf("./data/guilds/%s.json", agent.Message.GuildID)

	exists, err := utils.CreateFileIfNotExist(path)
	if err != nil {
		utils.LogErrorToChan(agent, err)
		return err
	}
	if exists {
		utils.SendMessageWithMention("This Guild is already registered!", "", agent)
		return os.ErrExist
	}
	if guildWriteFile(agent, data) != nil {
		return err
	}
	return nil
}

// GuildInit Create guild's data file
func GuildInit(agent sessionAgent.DiscordAgent) {
	data := createData(agent)
	if data.GuildID == "" {
		return
	}
	if writeData(agent, data) != nil {
		return
	}
	discordRoleSet(data, "", "admin", agent)
	utils.SendMessageWithMention("Guild registered successfully!", "", agent)
}
