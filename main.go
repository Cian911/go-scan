package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

const protocol = "tcp"

func worker(ports, results chan int, address string) {
	for p := range ports {
		conn, err := net.Dial(protocol, fmt.Sprintf("%s:%d", address, p))
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func createFlags() (address *string, portBufferNum *int, portRange *int) {
	address = flag.String("address", "scanme.nmap.org", "The address you wish to scan.")
	portBufferNum = flag.Int("buffers", 100, "Number of ports you want to buffer ahead of time for faster scanning. Default is 100.")
	portRange = flag.Int("range", 100, "How many ports do you want to scan? Default is 100.")
	flag.Parse()
	return
}

func main() {
	address, portBufferNum, portRange := createFlags()

	fmt.Printf("Scanning %s up to port %d", *address, *portRange)

	ports := make(chan int, *portBufferNum)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, *address)
	}

	go func() {
		for i := 1; i <= *portRange; i++ {
			ports <- i
		}
	}()

	for i := 0; i < *portRange; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("Port %d open.\n", port)
	}
}
