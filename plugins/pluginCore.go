package plugins

import "github.com/bwmarrin/discordgo"

type Command interface {
	GetAliases() []string
	Run(string, []string, *discordgo.Message, *discordgo.Session) error
}

type ParameterError struct{
	Msg string
}

type PermissionError struct{
	Msg string
}

func (p *PermissionError) Error() string {
	return p.Msg
}

func (p *ParameterError) Error() string{
	return p.Msg
}
