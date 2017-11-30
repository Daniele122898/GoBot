package fun

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"strings"
	"github.com/Daniele122898/GoBot/helpers/embeds"
)

type rps struct {
	aliases []string
}

//get  new rps command.
func NewRps() plugins.Command{
	return 	&rps{aliases: []string{"rps", "rockpaperscissor"}}
}

func (r rps) GetAliases() []string {
	return r.aliases
}

func (rps) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	if len(args) != 1{
		return &plugins.ParameterError{"Please only choose one of the three options:\n`rock`, `paper`, `scissor`"}
	}
	var bot string
	switch strings.ToLower(args[0]) {
		case "rock", "rocks", "stone":
			bot = "paper"
		case "paper", "papers":
			bot = "scissor"
		case "scissor", "scissors":
			bot = "rock"
		default:
			return &plugins.ParameterError{"Please only choose one of the three options:\n`rock`, `paper`, `scissor`"}
	}
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Description: ":frowning: You lost! Bot chose: `"+bot+"`",
	})
	return err
}
