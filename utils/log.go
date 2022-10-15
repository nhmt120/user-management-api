package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func WriteLog(log_file string) {
	// var file *os.File
	path := "Logs/" + log_file
	if _, err := os.Stat(path); err == nil {
		file, err0 := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		fmt.Println(err0)
		log.SetOutput(file)
	} else if errors.Is(err, os.ErrNotExist) {
		if err0 := os.Mkdir("Logs", os.ModePerm); err0 != nil {
			fmt.Printf("Cannot create directory: %s.", path)
			fmt.Println(err0)
		}

		file, err1 := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

		if err1 != nil {
			fmt.Println(err1.Error())
			panic("Cannot open log file.")
		} else {
			log.SetOutput(file)
		}
	}

	// file, _ = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	// return file
}
