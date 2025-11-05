package product

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (p *ProductComander) Help(inputMessage *tgbotapi.Message) {

	helpText := `
/help__commerce__product — Show this help message
/list__commerce__product — List all entities
/list__commerce__product {"offset": offset, "limit": limit} — Paginated list
/get__commerce__product <id> — Get entity by ID
/delete__commerce__product <id> — Delete entity by ID
/new__commerce__product {"Title": "someTitle"} — Create new entity
/edit__commerce__product {"id": id, "Title": "newTitle"} — Edit existing entity`

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpText)

	_, err := p.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
