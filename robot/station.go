package robot

import (
	"github.com/go-vgo/robotgo"
	"github.com/happilymarrieddad/miner-go/utils"
	"time"
)

const (
	// Station Positions
	stationUndockX = 1280
	stationUndockY = 198

	// Hold Positions
	shipOreHoldMenuX  = 555
	shipOreHoldMenuY  = 600
	shipOreHoldStartX = 661
	shipOreHoldStartY = 572
	shipOreHoldEndX   = 707
	shipOreHoldEndY   = 581
	shipOreDragStartX = 707
	shipOreDragStartY = 601
	shipOreDragEndX   = 555
	shipOreDragEndY   = 650
)

// UndockFromStation undock from the station
func UndockFromStation() {

	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(stationUndockX, stationUndockY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick()

	utils.BasicDelay()
}

// SelectAllCargo select all cargo
func SelectAllCargo() {
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(shipOreHoldMenuX, shipOreHoldMenuY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick()
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(shipOreHoldStartX, shipOreHoldStartY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick("right")
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(shipOreHoldEndX, shipOreHoldEndY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseClick()
	time.Sleep(time.Millisecond * 250)
}

// DragSelectedCargoToStation drag selected cargo to station
func DragSelectedCargoToStation() {
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(shipOreDragStartX, shipOreDragStartY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseToggle("down", "left")
	time.Sleep(time.Millisecond * 250)
	robotgo.MoveSmooth(shipOreDragEndX, shipOreDragEndY, 0.1, 1.0)
	time.Sleep(time.Millisecond * 250)
	robotgo.MouseToggle("up", "left")
	time.Sleep(time.Millisecond * 250)
}

// UnloadCargoToStation unload the cargo to the station
func UnloadCargoToStation() {

	SelectAllCargo()
	DragSelectedCargoToStation()

	utils.BasicDelay()
}
