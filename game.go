package main

import (
	"image/color"
)

func startGeme() {
	for {
		if player1.hasWon() || player2.hasWon() {
			return
		}
		clearLeds()
		var ticksBetweenSwitch uint8 = 50
		var tickCount uint8 = 0
		var acitveLed int8 = ledLen - 1
		leds[acitveLed] = color.RGBA{R: 255, B: 128}
		ws.WriteColors(leds)
		for {
			if ticksBetweenSwitch == tickCount {
				tickCount = 0
				acitveLed = switchToNextActiveLed(acitveLed)
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
				if acitveLed == 0 && ticksBetweenSwitch > 10 {
					ticksBetweenSwitch -= 5
				}
			}
			if acitveLed == player1.ledIdx {
				if player1.button.Get() {
					player1.hasHit = true
					player1.validButtonPush = true
					leds[player1.ledIdx] = color.RGBA{G: 255}
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
					leds[player2.ledIdx] = color.RGBA{G: 255}
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

func switchToNextActiveLed(curActiveLed int8) (nowActiveLed int8) {
	leds[curActiveLed] = color.RGBA{}
	nowActiveLed = curActiveLed - 1
	if nowActiveLed == -1 {
		nowActiveLed = ledLen - 1
	}
	leds[nowActiveLed] = color.RGBA{R: 255, B: 128}
	ws.WriteColors(leds)
	return nowActiveLed
}
