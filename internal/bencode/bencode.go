package bencode

import (
	"fmt"
	bencode "github.com/IncSW/go-bencode"
)

func Encode(data interface{}) string {
	encodedData, err := bencode.Marshal(data)
	if err != nil {
		fmt.Println("Error during encoding:", err)
	}
	fmt.Println("Encoded data:", string(encodedData))
	return string(encodedData)
}

func Decode(data []byte) interface{} {
	decodedData, err := bencode.Unmarshal(data)
	if err != nil {
		fmt.Println("Error during decoding:", err)
		return nil
	}
	fmt.Printf("Decoded data: %#v\n", decodedData)

	return decodedData
}
