package utils

import "log"

import "time"

// Log log out the msg(s) to console
func Log(msgs ...interface{}) {
	log.Println(msgs...)
}

// BasicDelay basic delay
func BasicDelay() {
	time.Sleep(2 * time.Second)
}
