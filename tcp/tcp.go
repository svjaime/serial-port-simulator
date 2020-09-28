//Package tcp represents the app communication interface
package tcp

import (
	"fmt"
	"net"
)

//Connect connects tcp client to specified port
func Connect(port string) (net.Conn, error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		return nil, fmt.Errorf("Error resolving tcp address: %v", err)
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("Error dialing tcp: %v", err)
	}

	err = conn.SetLinger(0)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	fmt.Printf("Connected on port %v.\n", port)
	return conn, nil
}
