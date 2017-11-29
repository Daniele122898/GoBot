package interactions

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/weeb.go/src"
	"github.com/Daniele122898/GoBot/helpers/config"
)

type interact struct{
	aliases []string
}

//get  new interaction command.
func NewInteract() plugins.Command{
	return interact{aliases: []string{"auth", "tags", "types"}}
}

func (i interact) GetAliases() []string{
	return i.aliases
}

func (interact) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	var err error
	switch cmd {
	case "auth":
		err = auth(cmd, args, msg, session)
	case "tags":
		err = tags(cmd, args,msg, session)
	case "types":
		err = types(cmd, args, msg, session)
	}
	return err
}

func auth(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	err := weebgo.Authenticate(config.Get().Weeb)
	if err == nil {
		_, err = session.ChannelMessageSend(msg.ChannelID, "Successfully authenticated")
		return err
	}
	return err
}

func tags(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	td, err := weebgo.GetTags(false)
	if err != nil{
		return err
	}
	out:=""
	for _, s := range td.Tags {
		out += s+", "
	}
	_, err = session.ChannelMessageSend(msg.ChannelID, "```\n"+out+"\n```")
	return err
}

func types(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	td, err := weebgo.GetTypes(false)
	if err != nil{
		return err
	}
	out:=""
	for _, s := range td.Types {
		out += s+", "
	}
	_, err = session.ChannelMessageSend(msg.ChannelID, "```\n"+out+"\n```")
	return err
}
