package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func scanPort(wg *sync.WaitGroup, ip string, port int) {
	defer wg.Done()

	address := ip + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Printf("Port %d is OPEN\n", port)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <IP>")
		return
	}

	ip := os.Args[1]
	var wg sync.WaitGroup

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go scanPort(&wg, ip, port)
	}

	wg.Wait()
	fmt.Println("Scan completed.")
}
