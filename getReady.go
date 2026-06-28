package main

import (
	"image/color"
	"time"
)

func getReady() {
	for {
		if player1.button.Get() {
			leds[player1.ledIdx] = player1.playerColor
		} else {
			leds[player1.ledIdx] = color.RGBA{R: 0}
		}
		if player2.button.Get() {
			leds[player2.ledIdx] = player2.playerColor
		} else {
			leds[player2.ledIdx] = color.RGBA{R: 0}
		}
		ws.WriteColors(leds)
		if player1.button.Get() && player2.button.Get() {
			time.Sleep(time.Millisecond * 200)
			leds[player1.ledIdx] = color.RGBA{}
			leds[player2.ledIdx] = color.RGBA{}
			time.Sleep(time.Second)
			return
		}

		tick()
	}
}
