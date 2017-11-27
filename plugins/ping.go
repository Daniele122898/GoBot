package plugins

import "github.com/bwmarrin/discordgo"

type ping struct {
	aliases []string
}

//get  new ping command.
func NewPing() Command{
	p := ping{aliases: []string{"ping"}}
	return p
}

func (p ping) GetAliases() []string{
	return p.aliases
}

func (ping) Run(cmd string, args []string, msg *discordgo.Message, s *discordgo.Session) (error){
	_, err := s.ChannelMessageSend(msg.ChannelID, "Pong!")
	return err
}
