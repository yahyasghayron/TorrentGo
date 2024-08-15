package main

import (
	"fmt"
	bencode "github.com/IncSW/go-bencode"
)

func main() {
	fmt.Println("Test Bencode encoding and decoding")

	// Sample data to encode
	data := map[string]interface{}{
		"foo":  1,
		"bar":  "baz",
		"list": []interface{}{1, 2, 3},
	}

  encodedData, err := bencode.Marshal(data)
  if err != nil {
    fmt.Println("Error during encoding:", err)
  }

	fmt.Println("Encoded data:", string(encodedData))

	// Decode the data
	decodedData, err := bencode.Unmarshal([]byte(encodedData))
	if err != nil {
		fmt.Println("Error during decoding:", err)
		return
	}

	// Print the decoded data
	fmt.Printf("Decoded data: %#v\n", decodedData)
}

