package tg

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GroupMap = map[string]int64

type Bot struct {
	API      *tgbotapi.BotAPI
	groups   GroupMap
	Replys   chan *tgbotapi.Message
	Commands BotCommands
}

func NewBot(tg_token string, groups GroupMap, commands BotCommands) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(tg_token)
	if err != nil {
		return nil, err
	}
	botAPI.Debug = true
	botAPI.Request(commands.Cfg())
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	bot := &Bot{
		API:    botAPI,
		groups: groups,
		Replys: make(chan *tgbotapi.Message),
	}

	return bot, nil
}

func (bot *Bot) MessageListener() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.API.GetUpdatesChan(u)

	for update := range updates {
		message := update.Message
		// ignore any non-Message Updates
		if message == nil {
			continue
		}

		if message.ReplyToMessage != nil {
			fmt.Println("Add to chan:", message.ReplyToMessage)
			bot.Replys <- message
			continue
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "")
		text, err := bot.Commands.Run(message.Command())
		if err != nil {
			continue
		}
		msg.Text = text
		bot.API.Send(msg)
	}
}

func (bot *Bot) SendGroupMessage(groupName, text string) (tgbotapi.Message, error) {
	group := bot.groups[groupName]
	msg := tgbotapi.NewMessage(group, text)
	return bot.API.Send(msg)
}

func (bot *Bot) EditGroupMessage(groupName string, messageID int, text string) {
	group := bot.groups[groupName]
	msg := tgbotapi.NewEditMessageText(group, messageID, text)
	message, _ := bot.API.Send(msg)
	log.Printf("Edited message %d in group %s", message.MessageID, groupName)
}

func (bot *Bot) SendPoll(groupName, text string, options ...string) (tgbotapi.Message, error) {
	group := bot.groups[groupName]
	poll := tgbotapi.NewPoll(group, text, options...)
	poll.IsAnonymous = false
	return bot.API.Send(poll)
}
