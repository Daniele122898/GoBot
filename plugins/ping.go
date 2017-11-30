package plugins

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type ping struct {
	aliases []string
}

//get  new ping command.
func NewPing() Command{
	return 	&ping{aliases: []string{"ping"}}
}

func (p ping) GetAliases() []string{
	return p.aliases
}

func (ping) Run(cmd string, args []string, msg *discordgo.Message, s *discordgo.Session) (error){
	start := time.Now()
	m, err := s.ChannelMessageSend(msg.ChannelID, "Pong!")
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageEdit(msg.ChannelID, m.ID, m.Content+" ("+time.Since(start).String()+") :ping_pong:")
	return err
}
