package main

import (
	"image/color"
	"time"
)

func startGeme() {
	var startLed int8
	var ticksBetweenSwitch int8
	var tickCount int8
	var acitveLed int8
	var gameColor color.RGBA
	for {
		gameColor = color.RGBA{R: 128, G: 128, B: 128}
		if player1.hasWon() || player2.hasWon() {
			return
		}
		clearLeds()
		ticksBetweenSwitch = 50 - (player1.score*5 + player2.score*5)
		if ticksBetweenSwitch < 20 {
			ticksBetweenSwitch = 20
		}
		println(ticksBetweenSwitch)
		tickCount = 0
		acitveLed = reverseLedIdx(0, ledLen)
		if playerLastScored == 1 {
			acitveLed = reverseLedIdx(5, ledLen)
		}
		startLed = acitveLed
		leds[acitveLed] = gameColor
		ws.WriteColors(leds)
		for {
			if ticksBetweenSwitch == tickCount {
				tickCount = 0
				acitveLed = switchToNextActiveLed(acitveLed, gameColor)
				if acitveLed == player1.ledIdx && player1.validButtonPush {
					player2.scored()
					break
				}
				if acitveLed == player1.ledIdx-1 && !player1.hasHit {
					player2.scored()
					break
				} else {
					player1.hasHit = false
				}

				if acitveLed == player2.ledIdx && player2.validButtonPush {
					player1.scored()
					break
				}
				if acitveLed == player2.ledIdx-1 && !player2.hasHit {
					player1.scored()
					break
				} else {
					player2.hasHit = false
				}
				if acitveLed == startLed && ticksBetweenSwitch > 10 {
					ticksBetweenSwitch -= 5
					println(ticksBetweenSwitch)
				}
			}
			if acitveLed == player1.ledIdx {
				if player1.button.Get() {
					player1.hasHit = true
					player1.validButtonPush = true
					leds[player1.ledIdx] = player1.playerColor
					gameColor = player1.playerColor
				}
			} else {
				if player1.button.Get() {
					if player1.validButtonPush == false {
						player2.scored()
						break
					}
				} else {
					player1.validButtonPush = false
				}
			}
			if acitveLed == player2.ledIdx {
				if player2.button.Get() {
					player2.hasHit = true
					player2.validButtonPush = true
					leds[player2.ledIdx] = player2.playerColor
					gameColor = player2.playerColor
				}
			} else {
				if player2.button.Get() {
					if player2.validButtonPush == false {
						player1.scored()
						break
					}
				} else {
					player2.validButtonPush = false
				}
			}

			ws.WriteColors(leds)
			tick()
			tickCount++

		}

	}
}

func switchToNextActiveLed(curActiveLed int8, c color.RGBA) (nowActiveLed int8) {
	leds[curActiveLed] = color.RGBA{}
	nowActiveLed = curActiveLed - 1
	if nowActiveLed == -1 {
		nowActiveLed = ledLen - 1
	}
	leds[nowActiveLed] = c
	ws.WriteColors(leds)
	return nowActiveLed
}

func showScore() {
	player1.setScoreInLeds()
	player2.setScoreInLeds()
	ws.WriteColors(leds)
	time.Sleep(time.Second * 2)
}
