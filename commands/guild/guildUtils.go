package guild

import (
	"encoding/json"
	"fmt"
	"github.com/Millefeuille42/DiscordBotSkeleton/sessionAgent"
	"github.com/Millefeuille42/DiscordBotSkeleton/utils"
	"io/ioutil"
	"os"
)

// GuildLoadFile Returns guild data from file
func GuildLoadFile(agent sessionAgent.DiscordAgent, silent bool, id string) (GuildData, error) {
	data := GuildData{}

	if id == "" {
		id = agent.Message.GuildID
	}
	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/guilds/%s.json", id))
	if err != nil {
		if !silent {
			utils.LogErrorToChan(agent, err)
		}
		return GuildData{}, err
	}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		if !silent {
			utils.LogErrorToChan(agent, err)
		}
		return GuildData{}, err
	}

	return data, nil
}

// guildWriteFile Writes guild data to file
func guildWriteFile(agent sessionAgent.DiscordAgent, data GuildData) error {
	jsonGuild, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		utils.LogErrorToChan(agent, err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/guilds/%s.json", data.GuildID), jsonGuild, 0677)
	if err != nil {
		utils.LogErrorToChan(agent, err)
		return err
	}
	return nil
}

// GuildInitialCheck Required before guild related actions, checks if guild exists
func GuildInitialCheck(agent sessionAgent.DiscordAgent) bool {
	_, err := os.Stat(fmt.Sprintf("./data/guilds/%s.json", agent.Message.GuildID))
	if !os.IsNotExist(err) {
		return true
	}

	utils.SendMessageWithMention("This guild doesn't exist, create it with the "+"init command", "", agent)
	return false
}
