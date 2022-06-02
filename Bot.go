package github.com/Millefeuille42/DiscordBotSkeleton

import (
	"fmt"
	"github.com/Millefeuille42/DiscordBotSkeleton/commands"
	"github.com/Millefeuille42/DiscordBotSkeleton/commands/guild"
	"github.com/Millefeuille42/DiscordBotSkeleton/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

var mainBot BotSkeleton

type BotSkeleton struct {
	FSPrep          func() error
	CommandMapSetup func(*BotSkeleton)
	CommandMap      map[string]commandHandler
	Session         *discordgo.Session
	Prefix          string
	OwnerId         string
}

// prepFileSystem Create required directories
func prepFileSystem() error {
	err := utils.CreateDirIfNotExist("./data")
	if err != nil {
		return err
	}
	err = utils.CreateDirIfNotExist("./data/guilds")
	return err
}

func setupFunctionsMap(skeleton *BotSkeleton) {
	//AdminCommands no args
	skeleton.CommandMap["init"] = guild.GuildInit
	skeleton.CommandMap["params"] = guild.AdminSendSettings
	skeleton.CommandMap["admin"] = guild.AdminSet

	skeleton.CommandMap["help"] = commands.SendHelp
}

// StartBot Starts discord bot
func (bt *BotSkeleton) StartBot() {
	var err error

	bt.Session, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	utils.CheckError(err)
	bt.Session.AddHandler(messageHandler)
	err = bt.Session.Open()
	utils.CheckError(err)
	fmt.Println("Discord bot created")
	if os.Getenv("SEGBOT_IN_PROD") == "" {
		channel, err := bt.Session.UserChannelCreate(bt.OwnerId)
		if err != nil {
			log.Println(err)
			return
		}
		hostname, _ := os.Hostname()
		_ = utils.SendMessageWrapper(bt.Session, channel.ID, "Bot up - "+
			time.Now().Format(time.Stamp)+" - "+hostname)
	}
	if bt.Prefix == "" {
		bt.Prefix = "!"
	}
	utils.SetUpCloseHandler(bt.Session)
}

func New() BotSkeleton {
	mainBot = BotSkeleton{
		FSPrep:          prepFileSystem,
		CommandMapSetup: setupFunctionsMap,
		CommandMap:      make(map[string]commandHandler),
		Session:         nil,
		Prefix:          os.Getenv("BOT_PREFIX"),
		OwnerId:         os.Getenv("BOT_OWNER"),
	}
	return mainBot
}
