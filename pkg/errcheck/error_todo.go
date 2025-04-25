package errcheck

import "log"

func LogIfError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %v\n", msg, err)
	}
}

func FatalIfError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v\n", msg, err)
	}
}
