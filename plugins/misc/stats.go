package misc

import (
	"github.com/bwmarrin/discordgo"
	"github.com/serenity/GoBot/plugins"
	"runtime"
	"strconv"
	"github.com/serenity/GoBot/helpers/embeds"
	"github.com/dustin/go-humanize"
)

type stats struct{
	aliases []string
}

//get  new stats command.
func NewStats() plugins.Command{
	p := stats{aliases: []string{"sys", "info", "stats"}}
	return p
}

func (p stats) GetAliases() []string{
	return p.aliases
}

func (stats) Run(cmd string, args []string, msg *discordgo.Message, session *discordgo.Session) (error){
	//count users, guilds and channels
	users := make(map[string]string)
	channels := 0
	guilds := session.State.Guilds
	for _, g := range guilds {
		channels += len(g.Channels)

		for _, u := range g.Members {
			users[u.User.ID] = u.User.Username
		}
	}

	//Ram stats
	var ram runtime.MemStats
	runtime.ReadMemStats(&ram)

	//Get uptime
	//TODO UPTIME


	_, err := session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
		Color: embeds.DEFAULT_COLOR,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: session.State.User.AvatarURL("1024"),
		},
		Fields: []*discordgo.MessageEmbedField{
			//System
			{Name: "GO Version", Value: runtime.Version(), Inline: true},
			{Name: "Bot Version", Value:"1.0.0-alpha.1", Inline:true},
			//Bot
			{
				Name: "Used RAM",
				Value: humanize.Bytes(ram.HeapAlloc) + "/" + humanize.Bytes(ram.HeapSys),
				Inline:true,
			},
			{Name: "Collected Garbage", Value:humanize.Bytes(ram.TotalAlloc), Inline:true},
			{Name: "Running Coroutines", Value:strconv.Itoa(runtime.NumGoroutine()), Inline:true},
			//Discord
			{Name: "Connected Servers", Value:strconv.Itoa(len(guilds)), Inline:true},
			{Name: "Watching Channels", Value:strconv.Itoa(channels), Inline:true},
			{Name: "Users with access", Value:strconv.Itoa(len(users)), Inline:true},
		},
	})
	return err
}


