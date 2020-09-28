package message

import (
	"fmt"
	"io"
	"log"
	"time"
)

//Listen handles incoming communication
func Listen(rw io.ReadWriter) {

	for {

		buff := make([]byte, 128)
		n, err := rw.Read(buff)

		if err != nil {
			if err != io.EOF {
				log.Printf("Failed to read: %v", err)
				continue
			}
			//log.Println("EOF")
			continue
		}

		log.Printf("Read %v bytes.\n", n)
		fmt.Printf("MSG READ: %# 02X\n", buff[:n])

		resp, err := getResponse(buff[:n])
		if err != nil {
			log.Printf("error fetching response: %v\n", err)
			continue
		}

		time.Sleep(100 * time.Millisecond)

		n, err = rw.Write(resp)
		if err != nil {
			log.Printf("Failed to write back: %v", err)
			continue
		}

		log.Printf("Wrote %v bytes.\n", n)
		fmt.Printf("MSG WROTE: %# 02X\n", resp)
	}
}
