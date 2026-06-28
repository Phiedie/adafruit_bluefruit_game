package main

import (
	"image/color"
	"time"
)

func clearLeds() {
	for i := range leds {
		leds[i] = color.RGBA{}
	}
	ws.WriteColors(leds)
}

func blinkLeds(blinks uint8, durationOn time.Duration, durationOff time.Duration) {
	for range blinks {
		ws.SetBrightness(0)
		ws.WriteColors(leds)
		time.Sleep(durationOff)
		ws.SetBrightness(ledBrightness)
		ws.WriteColors(leds)
		time.Sleep(durationOn)
	}
}
