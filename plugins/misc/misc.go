package misc

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/GoBot/helpers/embeds"
)

type misc struct{
	aliases []string
}

//get  new Misc command.
func NewMisc() plugins.Command{
	return &misc{aliases: []string{"git","gitlab", "github","invite", "inv","about", "support"}}
}

func (m misc) GetAliases() []string{
	return m.aliases
}

func (misc) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	var err error
	switch cmd {
	case "git","gitlab", "github":
		err = git(msg, session)
	case "invite", "inv":
		err = invite(msg, session)
	case "about":
		err = about(msg, session)
	case "support":
		err = support(msg, session)
	}
	return err
}

func support(msg *discordgo.Message, session *discordgo.Session) error{
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title:"Click here for Support Server",
		URL: "https://discordapp.com/invite/Pah4yj5",
	})
	return err
}

func about(msg *discordgo.Message, session *discordgo.Session) error{
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title:"ℹ️ About",
		URL: "https://github.com/Daniele122898/GoBot",
		Description:"Hei there (｡･ω･)ﾉﾞ\nI was created by Serenity#0783. You can find him [here](https://discordapp.com/invite/Pah4yj5)",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:"How was i created?",
				Value:"I was written in Go using the Discord.Go wrapper.\n" +
					"For more info use `b!sys`\n" +
					"Or visit [my Github page](https://github.com/Daniele122898/GoBot)",
				Inline:false,
			},
			{
				Name:"About me",
				Value:"I'm a cute little bunny uwu\nPat me please :>",
				Inline:false,
				},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: session.State.User.AvatarURL("1024"),
		},
	})
	return err
}

func git(msg *discordgo.Message, session *discordgo.Session) error{
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title:"Click here for Github",
		URL: "https://github.com/Daniele122898/GoBot",
	})
	return err
}

func invite(msg *discordgo.Message, session *discordgo.Session) error{
	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title:"Click here for Invite",
		Description: "Bunny needs all the perms if you intend to use all of its features. Unchecking certain perms will inhibit some of Bunnys' functions",
		URL: "https://discordapp.com/oauth2/authorize?client_id=379003892001931265&scope=bot&permissions=305523831",
	})
	return err
}
