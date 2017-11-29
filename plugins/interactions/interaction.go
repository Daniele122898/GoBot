package interactions

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/weeb.go/src"
	"github.com/Daniele122898/GoBot/helpers/config"
	"github.com/Daniele122898/weeb.go/src/net"
	"github.com/Daniele122898/GoBot/helpers/embeds"
)

type interact struct{
	aliases []string
}

//get  new interaction command.
func NewInteract() plugins.Command{
	return interact{aliases: []string{"auth", "tags", "types", "pat"}}
}

func (i interact) GetAliases() []string{
	return i.aliases
}

func (interact) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	var err error
	switch cmd {
	case "auth":
		err = auth(msg, session)
	case "tags":
		err = tags(msg, session)
	case "types":
		err = types(msg, session)
	case "pat":
		err = pat(msg, session)
	}
	return err
}

func pat(msg *discordgo.Message, session *discordgo.Session) error{
	d, err := weebgo.GetRandomImage("pat", nil, net.GIF, net.FALSE, false)
	if err!=nil{
		return err
	}
	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Image: &discordgo.MessageEmbedImage{
			URL:d.Url,
		},
	})
	return err
}

func auth(msg *discordgo.Message, session *discordgo.Session) (error){
	err := weebgo.Authenticate(config.Get().Weeb)
	if err == nil {
		_, err = session.ChannelMessageSend(msg.ChannelID, "Successfully authenticated")
		return err
	}
	return err
}

func tags(msg *discordgo.Message, session *discordgo.Session) (error){
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

func types(msg *discordgo.Message, session *discordgo.Session) (error){
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
