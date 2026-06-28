package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const ledNeoPin = machine.WS2812
const ledLen int8 = 10
const maxScore int8 = 5
const ledBrightness = 30

type player struct {
	button            machine.Pin
	ledIdx            int8
	score             int8
	validButtonPush   bool
	hasHit            bool
	firstScoreLed     int8
	ledScoreDirection int8
}

var player1 = player{
	button:            machine.BUTTONA,
	ledIdx:            reverseLedIdx(7, ledLen),
	firstScoreLed:     reverseLedIdx(ledLen-1, ledLen),
	ledScoreDirection: 1,
}
var player2 = player{
	button:            machine.BUTTONB,
	ledIdx:            reverseLedIdx(2, ledLen),
	firstScoreLed:     reverseLedIdx(0, ledLen),
	ledScoreDirection: -1,
}
var ws = ws2812.New(ledNeoPin)
var leds = make([]color.RGBA, 10)

func main() {
	initIO()
	getReady()
	startGeme()
}

func reverseLedIdx(idx int8, ledsLen int8) int8 {
	return ledsLen - 1 - idx
}

func initIO() {
	player1.button.Configure(machine.PinConfig{machine.PinInputPulldown})
	player2.button.Configure(machine.PinConfig{machine.PinInputPulldown})
	ledNeoPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws.SetBrightness(ledBrightness)
}

func tick() {
	time.Sleep(time.Millisecond * 10)
}
