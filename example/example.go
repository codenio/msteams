package main

import (
	"fmt"

	"github.com/codenio/msteams"
)

func main() {
	fmt.Println("Hellow World")
	m := msteams.New("https://outlook.office.com/webhook/210d3e8b-2938-47cc-8ef4-26318f0b8104@e84aae58-fb19-400e-b9ee-2609d69c396a/IncomingWebhook/724cbbb5a00d47eeb9c22c74e1564b95/62e197a2-c4b4-479a-8c0b-361ff990d63a")
	m.Card = m.NewMessageCard("Hi", "Testing OpenURI", "this is my summary")
	m.Card.SetThemeColor("2DC72D")
	m.Card.AddSections("foo", "bar", true)
	m.Card.CreateOpenURIAction("Test Open URI", []msteams.URI{{
		OS:  msteams.Windows,
		URI: "https://www.google.com",
	},
	})

	m.SendMessage()
}
