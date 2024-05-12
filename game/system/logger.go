package system

import "log"

type Logger struct {
	level string
}

func (g *Logger) log(msg string) {
	log.Fatal(msg)
}
