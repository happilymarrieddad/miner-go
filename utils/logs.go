package utils

import "log"

// Log log out the msg(s) to console
func Log(msgs ...interface{}) {
	log.Println(msgs...)
}
