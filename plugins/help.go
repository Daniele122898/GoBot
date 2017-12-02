package plugins

import (
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/GoBot/helpers/embeds"
)

type help struct{
	aliases []string
}

var(
	commands []Command
)

//get  new help command.
func NewHelp() Command{
	return 	&help{aliases: []string{"help", "h"}}
}

func (h help) GetAliases() []string{
	return h.aliases
}

func RegisterCommands(cmds []Command){
	commands = cmds
}

func (help) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	out := ""
	for _, cmd := range commands{
		for _, a := range cmd.GetAliases(){
			out += "`"+a+"`, "
		}
	}
	_, err:= session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title:"Available Commands",
		Description: out,
	})
	return err
}

