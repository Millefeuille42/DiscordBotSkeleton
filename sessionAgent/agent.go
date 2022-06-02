package sessionAgent

import "github.com/bwmarrin/discordgo"

// DiscordAgent Contains discord's session and message structs, and the guild's command channel
type DiscordAgent struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Channel string

	Content string
	Command string
	Args    []string
}
