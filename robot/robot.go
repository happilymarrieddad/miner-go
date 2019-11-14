package robot

import "github.com/go-vgo/robotgo"

// KeyTap - tap a key
func KeyTap(letter string, arr ...[]interface{}) {
	robotgo.KeyTap(letter, arr)
}
