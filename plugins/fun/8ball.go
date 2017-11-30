package fun

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/GoBot/helpers/embeds"
	"math/rand"
)

type eball struct{
	aliases []string
}

var (
	answers = []string {
		"Signs point to yes. ",
		"Yes.",
		"Reply hazy, try again.",
		"Without a doubt. ",
		"My sources say no. ",
		"As I see it, yes. ",
		"You may rely on it.",
		"Concentrate and ask again.",
		"Outlook not so good. ",
		"It is decidedly so.",
		"Better not tell you now.",
		"Very doubtful. ",
		"Yes - definitely. ",
		"It is certain. ",
		"Cannot predict now. ",
		"Most likely. ",
		"Ask again later. ",
		"My reply is no. ",
		"Outlook good. ",
		"Don't count on it.",
	}
)

//get  new 8ball command.
func NewEball() plugins.Command{
	return &eball{aliases: []string{"8b", "8ball"}}
}

func (b eball) GetAliases() []string{
	return b.aliases
}

func (eball) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	if len(args) == 0 {
		return &plugins.ParameterError{"Please add a question..."}
	}

	a := answers[rand.Intn(len(answers))]

	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title: "🎱 " + a,
	})
	return err
}
