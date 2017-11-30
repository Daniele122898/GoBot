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
	"github.com/Daniele122898/weeb.go/src/data"
	"fmt"
)

type interact struct{
	aliases []string
}

//get  new interaction command.
func NewInteract() plugins.Command{
	//i := &interact{aliases: []string{"tags", "types", "pat", "hug", "kiss", "slap", "highfive", "cuddle", "lick", "poke"}}
	i := &interact{}
	d, err := getTypes()
	if err != nil{
		fmt.Errorf("%s", "ERROR GETTING TYPES FROM WEEB.SH. ONLY USING DEFAULT TYPES")
		i.aliases = []string{"tags", "types", "pat", "hug", "kiss", "slap", "highfive", "cuddle", "lick", "poke"}
		return i
	}
	//fmt.Println("Got these weeb.sh types: ", d.Types)
	//i.aliases = append(i.aliases,d.Types...)
	i.aliases = d.Types
	i.aliases = append(i.aliases, "tags", "types")
	return i
}

func (i interact) GetAliases() []string{
	return i.aliases
}

func (interact) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) error {
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
	case "highfive":
		err = highfive(msg, session)
	case "cuddle":
		err = cuddle(msg, session)
	case "lick":
		err = lick(msg, session)
	case "poke":
		err = poke(msg, session)
	default:
		err = defaultInter(cmd, msg, session)
	}
	return err
}

func defaultInter(cmd string, msg *discordgo.Message, session *discordgo.Session) error{
	d, err := weebgo.GetRandomImage(cmd, nil, net.ANY, net.FALSE, false)
	if err!=nil{
		session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.ERR_COLOR,
			Title: embeds.I_ERR+ " Failed to get image from API :/",
		})
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

func getTypes() (*data.TypesData, error){
	d, err := weebgo.GetTypes(false)
	return d, err
}

func poke(msg *discordgo.Message, session *discordgo.Session) error {
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)

	return inter("lick", "licked","(♥ω♥ ) ~♪", m,msg,session)
}

func lick(msg *discordgo.Message, session *discordgo.Session) error{
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)
	//check for self cuddle
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.DEFAULT_COLOR,
			Title:helpers.GiveUserAndDiscrim(m[0])+" why would you lick yourself?!?!? (̂ ˃̥̥̥ ˑ̫ ˂̥̥̥ )̂ ",
			Image: &discordgo.MessageEmbedImage{
				URL:"https://i.imgur.com/sfp9gih.gif",
			},
		})
		return err
	}
	return inter("lick", "licked","(♥ω♥ ) ~♪", m,msg,session)
}

func inter(typ, inter, emoji string, m []*discordgo.User, msg *discordgo.Message, session *discordgo.Session) error {
	d, err := weebgo.GetRandomImage(typ, nil, net.GIF, net.FALSE, false)
	if err!=nil{
		session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.ERR_COLOR,
			Title: embeds.I_ERR+ " Failed to get image from API :/",
		})
		return err
	}

	//Get list of patted users
	us :=""
	for _, u := range m {
		us+= helpers.GiveUserAndDiscrim(u)+", "
	}
	us = strings.TrimRight(us, ", ")

	if len(us) > 220{
		us = us[:220]+"..."
	}

	_, err = session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Title: helpers.GiveUserAndDiscrim(msg.Author)+" "+inter+" "+us+ " "+emoji,
		Image: &discordgo.MessageEmbedImage{
			URL:d.Url,
		},
	})
	return err
}

func cuddle(msg *discordgo.Message, session *discordgo.Session) error {
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)
	//check for self cuddle
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.DEFAULT_COLOR,
			Title:helpers.GiveUserAndDiscrim(m[0])+" don't cuddle yourself ;-; At least take this pillow (̂ ˃̥̥̥ ˑ̫ ˂̥̥̥ )̂ ",
			Image: &discordgo.MessageEmbedImage{
				URL:"http://i.imgur.com/CM0of.gif",
			},
		})
		return err
	}
	return inter("cuddle", "cuddled","o(≧∇≦o)", m,msg,session)
}

func highfive(msg *discordgo.Message, session *discordgo.Session) error {
	if msg.Mentions == nil || len(msg.Mentions) == 0{
		return &plugins.ParameterError{"Mention at least one user!"}
	}

	m := helpers.RemoveDuplicateUsers(msg.Mentions)

	//check for self high5
	if len(m) == 1 && m[0].ID == msg.Author.ID {
		_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.DEFAULT_COLOR,
			Title:helpers.GiveUserAndDiscrim(m[0])+" no friends to high five? (̂ ˃̥̥̥ ˑ̫ ˂̥̥̥ )̂ ",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i.imgur.com/L5CEoYE.gif",
			},
		})
		return err
	}
	return inter("highfive", "high-fived","°˖✧◝(⁰▿⁰)◜✧˖°", m,msg,session)
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

	return inter("kiss", "kissed","(✿ ♥‿♥)♥", m,msg,session)
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

	return inter("slap", "slapped","(ᗒᗩᗕ)՞", m,msg,session)
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

	return inter("pat", "patted","｡◕ ‿ ◕｡", m,msg,session)
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

	return inter("hug", "hugged","°˖✧◝(⁰▿⁰)◜✧˖°", m,msg,session)
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
	_, err = session.ChannelMessageSend(msg.ChannelID, "```\n"+out[:len(out)-2]+"\n```")
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
	_, err = session.ChannelMessageSend(msg.ChannelID, "Available Tags are:\n```\n"+out[:len(out)-2]+"\n```")
	return err
}
