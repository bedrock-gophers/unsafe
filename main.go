package main

import (
	"github.com/bedrock-gophers/unsafe/unsafe"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"log"
	"log/slog"
)

func main() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := server.DefaultConfig().Config(slog.Default())
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()
	for p := range srv.Accept() {
		accept(p)
	}
}

func accept(p *player.Player) {
	unsafe.Rotate(p, 90, 90)
}
