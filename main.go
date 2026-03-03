package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type PortResult struct {
	Port  int
	State string
}

func scanPort(ctx context.Context, results chan<- PortResult, ip string, port int, timeout time.Duration) {
	defer close(results) // لكن هذا لكل port، نحتاج channel مختلف

	select {
	case <-ctx.Done():
		return
	default:
	}

	address := net.JoinHostPort(ip, strconv.Itoa(port))
	dialer := &net.Dialer{Timeout: timeout}
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return
	}
	conn.Close()
	results <- PortResult{Port: port, State: "OPEN"}
}

func main() {
	ip := flag.String("ip", "", "Target IP address")
	ports := flag.String("ports", "1-1024", "Port range (e.g., 1-65535, 80,443,22)")
	timeout := flag.Duration("timeout", 3*time.Second, "Connection timeout")
	threads := flag.Int("threads", 100, "Number of concurrent threads")
	flag.Parse()

	if *ip == "" {
		fmt.Println("Usage: go run scanner.go -ip=192.168.1.1 -ports=1-1000 -threads=200")
		return
	}

	portList := parsePorts(*ports)
	if len(portList) == 0 {
		fmt.Println("Invalid port range")
		return
	}

	fmt.Printf("Scanning %s (%d ports) with %d threads...\n", *ip, len(portList), *threads)
	
	openPorts := scanPorts(*ip, portList, *timeout, *threads)
	
	if len(openPorts) > 0 {
		fmt.Printf("\n%d open ports found:\n", len(openPorts))
		for _, port := range openPorts {
			fmt.Printf("Port %d: OPEN\n", port)
		}
	} else {
		fmt.Println("\nNo open ports found.")
	}
}

func parsePorts(portSpec string) []int {
	var ports []int
	specs := strings.Split(portSpec, ",")
	
	for _, spec := range specs {
		spec = strings.TrimSpace(spec)
		if strings.Contains(spec, "-") {
			rangePorts(spec, &ports)
		} else {
			if port, err := strconv.Atoi(spec); err == nil {
				if port >= 1 && port <= 65535 {
					ports = append(ports, port)
				}
			}
		}
	}
	return ports
}

func rangePorts(spec string, ports *[]int) {
	parts := strings.Split(spec, "-")
	if len(parts) != 2 {
		return
	}
	
	start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	
	for port := start; port <= end && port <= 65535; port++ {
		*ports = append(*ports, port)
	}
}

func scanPorts(ip string, portList []int, timeout time.Duration, maxWorkers int) []int {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		sem     = make(chan struct{}, maxWorkers)
		results = make(chan PortResult, len(portList))
		openPorts []int
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Minute)
	)
	defer cancel()

	go func() {
		for result := range results {
			mu.Lock()
			if result.State == "OPEN" {
				openPorts = append(openPorts, result.Port)
			}
			mu.Unlock()
		}
	}()

	for _, port := range portList {
		wg.Add(1)
		sem <- struct{}{} // Acquire semaphore
		
		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore
			
			select {
			case <-ctx.Done():
				return
			default:
			}
			
			address := net.JoinHostPort(ip, strconv.Itoa(p))
			dialer := &net.Dialer{Timeout: timeout}
			conn, err := dialer.DialContext(ctx, "tcp", address)
			if err == nil {
				conn.Close()
				results <- PortResult{Port: p, State: "OPEN"}
			}
		}(port)
	}

	wg.Wait()
	close(results)
	
	sort.Ints(openPorts)
	return openPorts
}
