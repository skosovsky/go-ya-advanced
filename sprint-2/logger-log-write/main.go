package main

import (
	"log"
	"os"
)

func main() {
	logger()
	loggerNew()
	loggerFlags()
}

func logger() {
	file, err := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!") // 2024/05/16 08:09:30 Logging to a file in Go!
}

func loggerNew() {
	flog, err := os.OpenFile("./server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()

	myLog := log.New(flog, "serv ", log.LstdFlags|log.Lshortfile)
	myLog.Println("Start server")  // serv 2024/05/16 08:06:07 main.go:34: Start server
	myLog.Println("Finish server") // serv 2024/05/16 08:06:07 main.go:35: Finish server
}

func loggerFlags() {
	file, err := os.OpenFile("./client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Logging to a file in Flags") // 2024/05/16 08:09:30 main.go:50: Logging to a file in Flags
}
