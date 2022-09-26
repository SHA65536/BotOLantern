package main

import (
	"botolantern"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot, err := botolantern.MakeHandler(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	err = bot.Start()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	err = bot.Stop()
	if err != nil {
		panic(err)
	}
}
