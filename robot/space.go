package robot

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/happilymarrieddad/miner-go/utils"
)

const (
	// Tabs
	stationTabX = 1125
	stationTabY = 196
	miningTabX  = 1037
	miningTabY  = 199

	// Belt
	startingBeltX = 975
	startingBeltY = 238
	endingBeltX   = 1021
	endingBeltY   = 247

	numOfBelts          = 4
	beltSizeIncrementor = 18
)

var currentBelt int

// SelectMiningTab select the mining tab
func SelectMiningTab() {
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(miningTabX, miningTabY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick()
	time.Sleep(time.Millisecond * 250)
}

// SelectStationTab select the mining tab
func SelectStationTab() {
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(stationTabX, stationTabY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick()
	time.Sleep(time.Millisecond * 250)
}

// WarpToBelt warp to the asteroid belt
func WarpToBelt() {

	if currentBelt == numOfBelts {
		currentBelt = 0
	} else {
		currentBelt++
	}

	currentBeltY := startingBeltY + (currentBelt * beltSizeIncrementor)

	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(startingBeltX, currentBeltY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick("right")
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(endingBeltX, endingBeltY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick()
	time.Sleep(time.Millisecond * 250)

	utils.BasicDelay()
}

// WarpToStation warp to the station
func WarpToStation() {

	utils.BasicDelay()
}
