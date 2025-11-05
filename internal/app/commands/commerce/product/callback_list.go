package product

import (
	"encoding/json"
	"log"

	"github.com/BohdanIpy/bot_256_demo/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

func (p *ProductComander) paginationKeyboard(currentOffset, currentLimit uint64) tgbotapi.InlineKeyboardMarkup {
	total := uint64(p.productService.GetNumberOfElements())

	var prevKey, nextKey tgbotapi.InlineKeyboardButton

	if currentOffset > 0 {
		var prevOffset uint64
		var prevLimit uint64

		if currentOffset >= currentLimit {
			prevOffset = currentOffset - currentLimit
			prevLimit = currentLimit
		} else {
			prevOffset = 0
			prevLimit = currentOffset
		}

		callbackData := CallbackListData{
			Offset: prevOffset,
			Limit:  prevLimit,
		}
		dataBytes, _ := json.Marshal(callbackData)
		callbackPath := path.CallbackPath{
			Domain:       "commerce",
			Subdomain:    "product",
			CallbackName: "list",
			CallbackData: string(dataBytes),
		}
		// ensure callbackPath.String() stays <= 64 bytes (Telegram limit)
		prevKey = tgbotapi.NewInlineKeyboardButtonData("⬅️ Prev", callbackPath.String())
	}

	if total > currentOffset+currentLimit {
		nextOffset := currentOffset + currentLimit

		var nextLimit uint64 = uint64(DEFAULT_PAGE_SIZE)
		remaining := total - nextOffset
		if remaining < uint64(DEFAULT_PAGE_SIZE) {
			nextLimit = remaining
		}

		callbackData := CallbackListData{
			Offset: nextOffset,
			Limit:  nextLimit,
		}
		dataBytes, _ := json.Marshal(callbackData)
		callbackPath := path.CallbackPath{
			Domain:       "commerce",
			Subdomain:    "product",
			CallbackName: "list",
			CallbackData: string(dataBytes),
		}
		nextKey = tgbotapi.NewInlineKeyboardButtonData("Next ➡️", callbackPath.String())
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	switch {
	case (prevKey.Text != "") && (nextKey.Text != ""):
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(prevKey, nextKey))
	case (prevKey.Text != ""):
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(prevKey))
	case (nextKey.Text != ""):
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(nextKey))
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (p *ProductComander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	// {Offset:4 Limit:4}
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	data, err := p.productService.List(parsedData.Offset, parsedData.Limit)
	if err != nil {
		log.Println(err)
		return
	}
	edit := tgbotapi.NewEditMessageTextAndMarkup(callback.Message.Chat.ID, callback.Message.MessageID, convertListIntoReadable(data), p.paginationKeyboard(parsedData.Offset, parsedData.Limit))

	_, err = p.bot.Send(edit)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
