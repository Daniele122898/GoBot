package fun

import (
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/GoBot/plugins"
	"net/url"
	"strings"
	"github.com/Daniele122898/GoBot/helpers/embeds"
)

type google struct{
	aliases []string
}

//get  new google command.
func NewGoogle() plugins.Command{
	return 	&google{aliases: []string{"google"}}
}

func (g google) GetAliases() []string {
	return g.aliases
}

func (google) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	if len(args) == 0{
		return &plugins.ParameterError{"Add something to search for!"}
	}
	search := url.PathEscape(strings.Join(args, " "))
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color:embeds.DEFAULT_COLOR,
		Title: "Click for Search results",
		URL: "https://lmgtfy.com/?q="+search,
	})
	return err
}
