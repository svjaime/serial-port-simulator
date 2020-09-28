package main

import (
	"log"

	"ubiwhere.com/serial-port-simulator/config"
	"ubiwhere.com/serial-port-simulator/message"
	"ubiwhere.com/serial-port-simulator/rest"
	"ubiwhere.com/serial-port-simulator/socat"
	"ubiwhere.com/serial-port-simulator/tcp"
)

//PORT is the default communication port
const PORT string = ":1234"

func main() {

	go func() {
		rest.HandleRequests()
	}()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = socat.Start(cfg.Path, PORT)
	if err != nil {
		log.Fatalf("%v", err)
	}

	conn, err := tcp.Connect(PORT)
	if err != nil {
		log.Fatalf("%v", err)
	}

	message.Listen(conn)
	config.DataFile.Close()

	/*
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			//ticker := time.NewTicker(1 * time.Second)

			connection.Responder(conn)
		}()

		wg.Wait()
	*/

}
