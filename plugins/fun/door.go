package fun

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/GoBot/helpers/embeds"
)

type door struct {
	aliases []string
}

//get  new door command.
func NewDoor() plugins.Command{
	return	&door{aliases: []string{"door"}}
}

func (d door) GetAliases() []string{
	return d.aliases
}

func (door) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	men := msg.Mentions
	if len(men) != 1 {
		return &plugins.ParameterError{Msg:"Please add exactly ONE @mention!"}
	}

	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Description: men[0].Username+"#"+men[0].Discriminator+" :point_right: :door: ",
	})

	return err
}
