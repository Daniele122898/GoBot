package interactions

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/weeb.go/src"
)

type interact struct{
	aliases []string
}

//get  new interaction command.
func NewInteract() plugins.Command{
	return interact{aliases: []string{"pat"}}
}

func (i interact) GetAliases() []string{
	return i.aliases
}

func (interact) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	err := weebgo.Authenticate("")
	return err
}
