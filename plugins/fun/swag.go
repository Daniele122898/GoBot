package fun

import (
	"github.com/bwmarrin/discordgo"
	"github.com/serenity/GoBot/plugins"
	"time"
)

type swag struct{
	aliases []string
}

//get  new swag command.
func NewSwag() plugins.Command{
	return swag{aliases: []string{"swag", "lenny"}}
}

func (s swag) GetAliases() []string{
	return s.aliases
}

func (swag) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	switch cmd {
		case "swag":
			return swaggy(msg, session)
		case "lenny":
			return lenny(msg, session)
	}
	return nil
}

func lenny(msg *discordgo.Message, session *discordgo.Session) error{
	_, err := session.ChannelMessageSend(msg.ChannelID, "( ͡° ͜ʖ ͡°)")
	return err
}

func swaggy(msg *discordgo.Message, session *discordgo.Session) error{
	m, err := session.ChannelMessageSend(msg.ChannelID, "( ͡° ͜ʖ ͡°)>⌐■-■")
	if err != nil {
		return err
	}
	time.Sleep(1e+9)
	_, err = session.ChannelMessageEdit(m.ChannelID, m.ID, "( ͡⌐■ ͜ʖ ͡-■)")
	return err
}
