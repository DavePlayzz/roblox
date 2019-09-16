package main

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	dgo *discordgo.Session
)

func init() {
	key, ok := os.LookupEnv("KEY")
	if !ok {
		errors.Fatalln("Missing Discord API Key: Set env var $KEY")
	}

	var err error
	dgo, err = discordgo.New("Bot " + key)
	if err != nil {
		errors.Fatalln(err)
	}

	err = dgo.Open()
	if err != nil {
		errors.Fatalln(err)
	}
}

func main() {
	// hook into presence update
	var err error
	err = dgo.UpdateStatus(0, "Logging...")

	dgo.AddHandler(func(ses *discordgo.Session, pre *discordgo.PresenceUpdate) {
		gam := pre.Game
		if gam == nil {
			return
		}

		usr := pre.User
		if usr.Bot {
			return
		}

		nam := gam.Name
		sta := gam.State
		uid := usr.ID

		// print them for now...
		fmt.Printf("UID: %s\nGame Name: %s\nGame State: %s", uid, nam, sta)
	})
}
