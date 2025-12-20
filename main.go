package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/kamilrahmatullin/lab/labs"
	"github.com/kamilrahmatullin/lab/utils"
)

func main() {
	file, err := utils.CreateLogFiles()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	logger := log.New(mw, "", log.Ldate|log.Ltime)
	reader := bufio.NewReader(os.Stdin)

	for {
		logger.Println("Enter the number of the lab assingment (1-5); 0 - Exit:")

		labNumber, err := utils.ReadInt(reader)
		if err != nil {
			logger.Fatal(err)
		}

		switch labNumber {
		case 0:
			logger.Println("Exit...")
			return
		case 1:
			logger.Println("Lab number is", labNumber)
		case 2:
			labs.Run2(logger, reader)
		case 3:
			labs.Run3(logger, reader)
		default:
			logger.Println("We could not find this lab. Bye!")
		}
	}
}
