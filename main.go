package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	p "github.com/serenity/GoBot/plugins"
	"github.com/serenity/GoBot/plugins/misc"
	"strings"
	"github.com/serenity/GoBot/helpers/config"
	"github.com/serenity/GoBot/plugins/fun"
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
		}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
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
	command = strings.TrimLeft(command, "b!") //remove prefix

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
					}
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

