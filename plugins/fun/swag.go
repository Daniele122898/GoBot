package fun

import (
	"github.com/bwmarrin/discordgo"
	"github.com/serenity/GoBot/plugins"
	"time"
)

type swag struct{
	aliases []string
}

//get  new stats command.
func NewSwag() plugins.Command{
	p := swag{aliases: []string{"swag"}}
	return p
}

func (s swag) GetAliases() []string{
	return s.aliases
}

func (swag) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	m, err := session.ChannelMessageSend(msg.ChannelID, "( ͡° ͜ʖ ͡°)>⌐■-■")
	if err != nil {
		return err
	}
	time.Sleep(1e+9)
	_, err = session.ChannelMessageEdit(m.ChannelID, m.ID, "( ͡⌐■ ͜ʖ ͡-■)")
	return err
}
