package plugins

import "github.com/bwmarrin/discordgo"

type Command interface {
	GetAliases() []string
	Run(string, []string, *discordgo.Message, *discordgo.Session) (error)
}
