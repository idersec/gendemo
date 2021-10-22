package biz

import "log"

func catchErr(logInfo string, err error) {
	if err != nil {
		log.Printf(logInfo+": %s\n", err)
	}
}
