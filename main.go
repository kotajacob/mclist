package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
)

const usage = "Usage: mclist <hostname>[:port]"

type status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
}

func main() {
	addr := getAddr()
	fmt.Printf("%s:\n", addr)
	resp, delay, err := bot.PingAndList(addr)
	if err != nil {
		log.Fatalf("ping failed: %v\n", err)
	}

	var s status
	err = json.Unmarshal(resp, &s)
	if err != nil {
		log.Fatalf("error unmarshaling response: %v\n", err)
	}

	fmt.Print(s)
	fmt.Println("Delay:", delay)
}

func getAddr() string {
	if len(os.Args) < 2 {
		fmt.Println("error no server address:", usage)
		os.Exit(1)
	}
	addr := os.Args[1]
	if !strings.ContainsRune(addr, ':') {
		addr += ":25565"
	}
	return addr
}

func (s status) String() string {
	var b strings.Builder
	fmt.Fprintln(&b, "Server:", s.Version.Name)
	fmt.Fprintln(&b, "Protocol:", s.Version.Protocol)
	fmt.Fprintln(&b, "Description:", s.Description)
	fmt.Fprintf(&b, "Players: %d/%d\n", s.Players.Online, s.Players.Max)
	for _, v := range s.Players.Sample {
		fmt.Fprintf(&b, "- [%s] %v\n", v.Name, v.ID)
	}
	return b.String()
}
