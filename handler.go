package main

import (
	"fmt"
	"os"

	tele "github.com/3JoB/telebot"
)

func Handler(b *tele.Bot) {
	b.Handle(tele.OnDocument, ReadFileInfo)
}

func ReadFileInfo(c tele.Context) error {
	c.Bot().Download(&c.Message().Document.File, c.Message().Document.FileName)
	info, _ := os.Stat(c.Message().Document.FileName)
	fmt.Printf("Size: %v byte \n, ModTime: %v \n", info.Size(), info.ModTime().String())
	return nil
}