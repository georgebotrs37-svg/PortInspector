🛡️ PortInspector Pro
A high-performance, feature-rich network port scanner built with **Go**. 

**PortInspector** is designed for speed and reliability, utilizing Go's advanced concurrency patterns to scan thousands of ports in seconds.

 ✨ Advanced Features
- **Flexible Port Specification:** Scan specific ports (`80,443`), ranges (`1-1024`), or a mix of both.
- **Controlled Concurrency:** Implements a **Semaphore pattern** to manage worker threads, preventing system resource exhaustion.
- **Smart Timeouts:** Fully customizable connection timeouts and global scan contexts.
- **Professional CLI:** Uses standard flags for easy integration into security workflows.
- **Ordered Results:** Automatically sorts open ports for clear, readable output.

 🚀 Technical Highlights
- **Goroutines & Channels:** For asynchronous, non-blocking network I/O.
- **Sync Package:** Uses `sync.WaitGroup` and `sync.Mutex` for safe data handling across threads.
- **Context API:** Implements `context.WithTimeout` for graceful scan termination and resource cleanup.

 🛠️ Installation
Ensure you have [Go](https://go.dev/dl/) installed:

```bash
git clone https://github.com/YOUR_USERNAME/PortInspector.git
cd PortInspector
go build -o scanner.exe main.go
💻 Usage
Run the scanner with various flags to customize your scan:
Flag	Description	Default
-ip	Target IP address (Required)	""
-ports	Ports to scan (e.g., 80,443,1-100)	1-1024
-threads	Number of concurrent workers	100
-timeout	Connection timeout duration	3s
Examples:
Basic Scan (First 1024 ports):
code
Bash
go run main.go -ip 1.1.1.1
Aggressive Scan (High threads & specific range):
code
Bash
go run main.go -ip 8.8.8.8 -ports 1-65535 -threads 1000 -timeout 1s
Specific Ports Scan:
code
Bash
go run main.go -ip 192.168.1.1 -ports 21,22,80,443,3306
📚 What I Practiced in This Project
Advanced Concurrency Control using Semaphores.
Complex String Parsing for port ranges and lists.
Using Context for lifecycle management of network connections.
Building a professional CLI tool with the flag package.
Disclaimer: This tool is for educational and authorized testing purposes only. Use responsibly.
code
