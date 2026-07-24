package handlers

import (
	"log"
)

func Catch(s string, t string, err error) {
	if err != nil {
		switch t {
		case "fatal":
			msg := "error in: " + s + ", error: " + err.Error()
			log.Fatalln(msg)
		case "warn":
			msg := "error in: " + s + ", error: " + err.Error()
			log.Println(msg)
		case "panic":
			msg := "error in: " + s + ", error: " + err.Error()
			log.Panicln(msg)

		}

	}

}
