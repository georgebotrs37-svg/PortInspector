# 🛡️ PortInspector
A high-performance, concurrent TCP port scanner built with **Go**.

## 🚀 Overview
**PortInspector** is a lightweight command-line tool designed to scan a range of network ports (1-1024) on a target IP address. Unlike traditional sequential scanners, this tool leverages Go's powerful concurrency model to perform scans in seconds.

## ✨ Key Features
- **Parallel Scanning:** Utilizes **Goroutines** to scan multiple ports simultaneously, significantly reducing execution time.
- **Synchronization:** Uses `sync.WaitGroup` to ensure the main program waits for all scanning processes to complete before exiting.
- **Smart Timeout:** Implements `net.DialTimeout` (1 second) to prevent the scanner from hanging on filtered or unresponsive ports.
- **CLI Ready:** Accepts the target IP address directly as a command-line argument.

## 🛠️ How It Works (Technical Details)
The core logic resides in the `scanPort` function. For every port in the range:
1. A new **Goroutine** is spawned.
2. It attempts a TCP connection using `net.DialTimeout`.
3. If the connection is successful, it identifies the port as **OPEN**.
4. It ensures the connection is properly closed (`conn.Close()`) to manage system resources efficiently.

## 💻 Usage
Make sure you have [Go](https://golang.org/dl/) installed on your machine.

1. Clone the repository or download `main.go`.
2. Run the following command in your terminal:

```bash
go run main.go <TARGET_IP>
