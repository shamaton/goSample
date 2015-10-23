package model

import (
	"log"
)

type Base struct {
	text string
}

func (b *Base) SetText(text string) {
	b.text = text
}

func (b *Base) TestPrint() {
	log.Println(b.text)
}
