package guild

// GuildData Contains the GuildID, a list of Admins and the guild Settings
type GuildData struct {
	GuildID  string
	Admins   []string
	AdminRole      string
	Locked   bool
}
