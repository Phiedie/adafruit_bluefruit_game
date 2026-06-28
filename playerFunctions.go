package main

import (
	"image/color"
	"time"
)

func (p *player) scored() {
	p.score++
	leds[p.ledIdx] = color.RGBA{B: 255}
	ws.WriteColors(leds)
	blinkLeds(2, time.Second/2, time.Second/2)
	clearLeds()
	time.Sleep(time.Second / 2)
	var ledIndex int8
	ledIndex = 0
	for i := range p.score {
		ledIndex = p.firstScoreLed + i*p.ledScoreDirection
		leds[ledIndex] = color.RGBA{B: 255}
	}
	ws.WriteColors(leds)
	time.Sleep(time.Second)
	clearLeds()
	return
}

func (p *player) hasWon() bool {
	if p.score == maxScore {
		blinkLeds(5, time.Second/2, time.Second/2)
		return true
	}
	return false
}
