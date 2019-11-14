package tasks

import (
	"github.com/happilymarrieddad/miner-go/robot"

	"time"
)

// MineInSystem mine in the current system the player is in
func MineInSystem() {
	var running bool

	for {
		// Undock from station
		robot.UndockFromStation()

		// Warp to belt
		robot.WarpToBelt()

		// Close distance to asteroids
		robot.CloseDistanceToAsteroids()

		// Loop
		maxIncr := 20
		for i := 0; i < maxIncr; i++ {
			// Target 2 asteroids
			robot.SelectAndTargetClosestAsteroids()

			// Set miner on each one
			robot.SetMinerOnEachAsteroid()

			// Delay for 55 seconds
			time.Sleep(time.Second * 55)

			// Stop miners and untarget
			robot.StopMinersAndUntarget()

			// Check if hold is full
			if robot.CheckIfHoldIsFull() {
				i = maxIncr
			}
		}

		// Warp station
		robot.WarpToStation()

		// Unload cargo to station
		robot.UnloadCargoToStation()

		if !running {
			break
		}
	}

}
