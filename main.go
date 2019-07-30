package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/buzzer"

	"github.com/lilrara/closeencounters/speaker"
	"github.com/lilrara/closeencounters/visor"
)

func main() {
	visor := visor.Visor()
	speaker := speaker.Speaker()

	//for {

	//visor.Off()

	visor.LED[0] = color.RGBA{R: 0x01, G: 0x01, B: 0x01}
	visor.LED[1] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}

	for i := range visor.LED[2:] {
		visor.LED[i+2] = color.RGBA{R: 0x01, G: 0x01, B: 0x01}
	}
	visor.Show()

	//	time.Sleep(100 * time.Millisecond)
	//}
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.D4, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.E4, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.C4, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.C3, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.G3, buzzer.Eighth)

	/* ime.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.E5, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.A5, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.D5, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.F3, buzzer.Eighth)
	time.Sleep(100 * time.Millisecond)
	speaker.Tone(buzzer.C4, buzzer.Eighth)
	*/

	for {
		time.Sleep(200 * time.Millisecond)

	}

}
