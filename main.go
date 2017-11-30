package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	p "github.com/Daniele122898/GoBot/plugins"
	"github.com/Daniele122898/GoBot/plugins/misc"
	"strings"
	"github.com/Daniele122898/GoBot/helpers/config"
	"github.com/Daniele122898/GoBot/plugins/fun"
	"github.com/Daniele122898/GoBot/plugins/interactions"
	"github.com/Daniele122898/GoBot/helpers/embeds"
	"time"
	"math/rand"
)

var(
	token string
	cmds []p.Command
	prefix = "b!"
)

func main() {
	token = config.Get().Token
	if token == ""{
		fmt.Println("TOKEN CANNOT BE EMPTY!")
		os.Exit(1)
	}
	dg, err := discordgo.New("Bot " +token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		os.Exit(1)
	}
	//Authenticate Wolke API
	interactions.Auth()
	//addallCommandsTocommandList
	initializeCommands()
	//Register the messagecreate func as callback
	// for MessageCreate events
	dg.AddHandler(messageCreate)
	//open a websocket connection to discord and behin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		os.Exit(2)
	}
	// Make the randomness more random
	rand.Seed(time.Now().UTC().UnixNano())
	//Wait here until CTRL-C or ither term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<- sc //blocks
	//cleanly close down the discord session
	dg.Close()
}

func initializeCommands(){
	cmds = []p.Command{
		p.NewPing(),
		misc.NewStats(),
		fun.NewSwag(),
		fun.NewDoor(),
		fun.NewEball(),
		fun.NewDice(),
		fun.NewGoogle(),
		fun.NewRps(),
		interactions.NewInteract(),
		}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	misc.MsgRec()
	//Ignore other bots and self
	if m.Author.Bot {
		return
	}
	//if prefix isn't there, return
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}
	args := strings.Split(m.Content, " ") //extract args
	command := strings.ToLower(args[0])//seperate command from arg list
	args = args[1:] //remove command from arg list
	command = strings.TrimPrefix(command, "b!") //remove prefix
	command = strings.TrimSuffix(command, " ") //remove extra space if its there
	//Try to match command
	f := false
	for _, cmd := range cmds{
		for _, a := range cmd.GetAliases(){
			if a == command {
				//go cmd.Run(command, args, m.Message, s) //run the command
				go func(){
					err:= cmd.Run(command, args, m.Message, s)
					if err != nil {
						//do error handling
						errorHandling(err, s, m)
					}
					misc.CmdEx()
				}()
				f = true
				break
			}
		}
		if f {
			break
		}
	}
}

func errorHandling(err error, session *discordgo.Session, msg *discordgo.MessageCreate){
	switch e := err.(type) {
	case *p.ParameterError:
		//session.ChannelMessageSend(msg.ChannelID, e.Error())
		session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.ERR_COLOR,
			Description: embeds.I_ERR + " "+ e.Error(),
		})
	case *p.PermissionError:
		session.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: embeds.ERR_COLOR,
			Description: embeds.I_ERR + " "+ e.Error(),
		})
	default:
		session.ChannelMessageSend(msg.ChannelID, "Something broke :/\n```\n"+e.Error()+"\n```")
	}
}

