package utils

import (
  "fmt"
  "time"
  "github.com/mitrovicstefan/cmus-artwork/cmusprocessor"
)


func UpdateTerminal() {
	for {
    asciiArt := cmusprocessor.GetArtPrint()
		ClearTerminal()
    fmt.Printf("%v", asciiArt)
		time.Sleep(10 * time.Second)
	}
}
