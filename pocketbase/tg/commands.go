package tg

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotCommands []BotCommand

type BotCommand struct {
	Command     string
	Description string
	Handler     func() string
}

func (bcs *BotCommands) Cfg() tgbotapi.SetMyCommandsConfig {
	var commands []tgbotapi.BotCommand
	for _, bc := range *bcs {
		commands = append(commands, tgbotapi.BotCommand{
			Command:     bc.Command,
			Description: bc.Description,
		})
	}
	return tgbotapi.NewSetMyCommands(commands...)
}

func (bcs *BotCommands) Run(command string) (string, error) {
	command = "/" + command
	for _, bc := range *bcs {
		if bc.Command == command {
			return bc.Handler(), nil
		}
	}
	return "", errors.New("command not found")
}
