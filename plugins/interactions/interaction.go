package interactions

import (
	"github.com/Daniele122898/GoBot/plugins"
	"github.com/bwmarrin/discordgo"
	"github.com/Daniele122898/weeb.go/src"
	"github.com/Daniele122898/GoBot/helpers/config"
	"github.com/Daniele122898/weeb.go/src/net"
	"github.com/Daniele122898/GoBot/helpers/embeds"
	"github.com/Daniele122898/GoBot/helpers"
	"strings"
)

type interact struct{
	aliases []string
}

//get  new interaction command.
func NewInteract() plugins.Command{
	return interact{aliases: []string{"tags", "types", "pat", "hug", "kiss", "slap"}}
}

func (i interact) GetAliases() []string{
	return i.aliases
}

func (interact) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	var err error
	switch cmd {
	case "tags":
		err = tags(msg, session)
	case "types":
		err = types(msg, session)
	case "pat":
		err = pat(msg, session)
	case "hug":
		err = hug(msg, session)
	case "kiss":
		err = kiss(msg, session)
	case "slap":
		err = slap(msg, session)
	}
	return err
}

func kiss(msg *discordgo.Message, session *discordgo.Session) error{
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)

	//check for self kiss
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.WARN_COLOR,
			Title:embeds.I_WARN +" " + helpers.GiveUserAndDiscrim(m[0])+" you may pat yourself or hug a pillow but kissing yourself is too much (๑•﹏•)",
		})
		return err
	}

	d, err := weebgo.GetRandomImage("kiss", nil, net.GIF, net.FALSE, false)
	if err!=nil{
		return err
	}

	//Get list of patted users
	patted :=""
	for _, u := range m {
		patted+= helpers.GiveUserAndDiscrim(u)+", "
	}
	patted = strings.TrimRight(patted, ", ")

	if len(patted) > 220{
		patted = patted[:220]+"..."
	}

	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title: helpers.GiveUserAndDiscrim(msg.Author)+" kissed "+patted+ "  (✿ ♥‿♥)♥",
		Image: &discordgo.MessageEmbedImage{
			URL:d.Url,
		},
	})
	return err
}

func slap(msg *discordgo.Message, session *discordgo.Session) error{
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)

	//check for self slap
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.WARN_COLOR,
			Title:embeds.I_WARN +" " + helpers.GiveUserAndDiscrim(m[0])+" why would you slap yourself... Are you okay? 〣( ºΔº )〣",
			Image: &discordgo.MessageEmbedImage{
				URL:"https://media.giphy.com/media/Okk9cb1dvtMxq/giphy.gif",
			},
		})
		return err
	}

	d, err := weebgo.GetRandomImage("slap", nil, net.GIF, net.FALSE, false)
	if err!=nil{
		return err
	}

	//Get list of patted users
	patted :=""
	for _, u := range m {
		patted+= helpers.GiveUserAndDiscrim(u)+", "
	}
	patted = strings.TrimRight(patted, ", ")

	if len(patted) > 220{
		patted = patted[:220]+"..."
	}

	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title: helpers.GiveUserAndDiscrim(msg.Author)+" slaps "+patted+ "  (ᗒᗩᗕ)՞",
		Image: &discordgo.MessageEmbedImage{
			URL:d.Url,
		},
	})
	return err
}

func pat(msg *discordgo.Message, session *discordgo.Session) error{
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)

	//check for selfpat
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.DEFAULT_COLOR,
			Title:helpers.GiveUserAndDiscrim(m[0])+", why are you patting yourself? Are you okay? ｡ﾟ･（>﹏<）･ﾟ｡",
			Image: &discordgo.MessageEmbedImage{
				URL:"https://i.imgur.com/QFtH3Gl.gif",
			},
		})
		return err
	}

	d, err := weebgo.GetRandomImage("pat", nil, net.GIF, net.FALSE, false)
	if err!=nil{
		return err
	}

	//Get list of patted users
	patted :=""
	for _, u := range m {
		patted+= helpers.GiveUserAndDiscrim(u)+", "
	}
	patted = strings.TrimRight(patted, ", ")

	if len(patted) > 220{
		patted = patted[:220]+"..."
	}

	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title: helpers.GiveUserAndDiscrim(msg.Author)+" pats "+patted+ "  ｡◕ ‿ ◕｡",
		Image: &discordgo.MessageEmbedImage{
			URL:d.Url,
		},
	})
	return err
}

func hug(msg *discordgo.Message, session *discordgo.Session) error {
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)

	//check for selfpat
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.DEFAULT_COLOR,
			Title:helpers.GiveUserAndDiscrim(m[0])+" don't hug yourself ;-; At least take this pillow (̂ ˃̥̥̥ ˑ̫ ˂̥̥̥ )̂ ",
			Image: &discordgo.MessageEmbedImage{
				URL:"http://i.imgur.com/CM0of.gif",
			},
		})
		return err
	}

	d, err := weebgo.GetRandomImage("hug", nil, net.GIF, net.FALSE, false)
	if err!=nil{
		return err
	}

	//Get list of hugged users
	hugged :=""
	for _, u := range m {
		hugged+= helpers.GiveUserAndDiscrim(u)+", "
	}
	hugged = strings.TrimRight(hugged, ", ")

	if len(hugged) > 220{
		hugged= hugged[:220]+"..."
	}

	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title: helpers.GiveUserAndDiscrim(msg.Author)+" hugs "+hugged+ "  °˖✧◝(⁰▿⁰)◜✧˖°",
		Image: &discordgo.MessageEmbedImage{
			URL:d.Url,
		},
	})
	return err
}

func Auth() (error){
	err := weebgo.Authenticate(config.Get().Weeb)
	if err == nil {
		return err
	}
	return err
}

func tags(msg *discordgo.Message, session *discordgo.Session) (error){
	//check if owner
	if msg.Author.ID != config.Get().Owner {
		return &plugins.PermissionError{"This command may only be used by the owner."}
	}
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
	//check if owner
	if msg.Author.ID != config.Get().Owner {
		return &plugins.PermissionError{"This command may only be used by the owner."}
	}
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
