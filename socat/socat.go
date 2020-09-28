//Package socat handles socat server launch
package socat

import (
	"fmt"
	"os/exec"
	"time"
)

//Start launches socat server
func Start(filePath, port string) error {

	arguments := []string{"pty,link=" + filePath + ",rawer", "tcp-listen" + port}

	cmd := exec.Command("socat", arguments...)
	//socat pty,link=ttySP0,rawer tcp-listen:1234

	fmt.Printf("Running socat command: %v\n", cmd.String())

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting socat: %v", err)
	}

	time.Sleep(100 * time.Millisecond)
	return nil
}
