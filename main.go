package main

func main() {
    trackerAddress := "tracker.openbittorrent.com:80" // UDP address
    infoHash := "d2474e86c95b19b8bcfdb92bc12c9d44667cfa36"
    peerID := "ABCDEFGHIJKLMNOPQRST"

    ConnectToUDPTracker(trackerAddress, infoHash, peerID)
}

