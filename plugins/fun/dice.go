package fun

import (
	"github.com/serenity/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/serenity/GoBot/helpers/embeds"
	"math/rand"
	"fmt"
)

type dice struct {
	aliases []string
}

//get  new diceroll command.
func NewDice() plugins.Command{
	return dice{aliases: []string{"roll", "dice", "rolldice"}}
}

func (d dice) GetAliases() []string{
	return d.aliases
}

func (dice) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	num := rand.Intn(6)+1
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color:embeds.DEFAULT_COLOR,
		Title: fmt.Sprintf("%s%d", "🎲 Rolled: ",num),
	})
	return err
}
