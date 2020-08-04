package main

import (
	"fmt"
	"github.com/cfi2017/sheep-bot/internal/commands"
	"github.com/cfi2017/sheep-bot/internal/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/dgcobra"
)

var token = viper.GetString("token")
var session *discordgo.Session
var handler *dgcobra.Handler

func main() {
	util.InitialiseFlags()
	util.InitialiseConfig()
	db, err := util.InitialisePersistence()
	if err != nil {
		panic(err)
	}

	util.SetDatabase(db)

	if token == "" {
		log.Fatal("missing token")
	}

	session, err = discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	// create and setup new handler
	handler = dgcobra.NewHandler(session)
	// add simple global prefix
	handler.AddPrefix("s!")
	handler.AddPrefix("sh!")
	handler.AddPrefix("sheep!")
	handler.AddPrefix("baa!")
	// set command factory
	handler.RootFactory = commands.RootCmdFactory
	// register new handler with discordgo
	handler.Start()

	// add ready handler to add bot mention when ID is available
	session.AddHandlerOnce(onReady)

	// open session
	err = session.Open()
	if err != nil {
		panic(err)
	}

	// cleanup
	log.Println("Bot is running. Press CTRL-C to exit.")
	waitForSig() // wait for termination signal
	err = session.Close()
	if err != nil {
		panic(err)
	}
}

func onReady(_ *discordgo.Session, ready *discordgo.Ready) {
	// register bot mention as new global prefix
	handler.AddPrefix(fmt.Sprintf("<@!%s> ", ready.User.ID))
}

func waitForSig() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
