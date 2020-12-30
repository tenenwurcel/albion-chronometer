package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/inancgumus/screen"
	hook "github.com/robotn/gohook"
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Print("--- Aperte + para começar ---")
	hook.Register(hook.KeyDown, []string{"+"}, func(e hook.Event) {
		f, err := os.Open("beep-07.mp3")
		if err != nil {
			log.Fatal(err)
		}

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()

		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

		screen.Clear()
		counter := 1

		for counter <= 90 {
			screen.MoveTopLeft()

			fmt.Printf("================> %d <================", counter)

			time.Sleep(time.Second)

			counter++
		}

		speaker.Play(streamer)
		time.Sleep(1 * time.Second)

		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println("--- Aperte + para começar ---")
	})

	s := hook.Start()
	<-hook.Process(s)
}
