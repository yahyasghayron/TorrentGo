package main

import (
    "fmt"
    "net"
    "time"
)

// ConnectToUDPTracker connects to a UDP tracker and sends a request.
func ConnectToUDPTracker(trackerAddress string, infoHash string, peerID string) {
    // Resolve UDP address
    addr, err := net.ResolveUDPAddr("udp", trackerAddress)
    if err != nil {
        fmt.Println("Error resolving address:", err)
        return
    }

    // Create UDP connection
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        fmt.Println("Error creating UDP connection:", err)
        return
    }
    defer conn.Close()

    // Prepare request (this is a simplified example)
    request := []byte("info_hash=" + infoHash + "&peer_id=" + peerID + "&port=6881")
    _, err = conn.Write(request)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }

    // Set a timeout for reading the response
    conn.SetReadDeadline(time.Now().Add(10 * time.Second))

    // Read response
    response := make([]byte, 1024)
    n, err := conn.Read(response)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    // Print response
    fmt.Println("Tracker response:", string(response[:n]))
}

