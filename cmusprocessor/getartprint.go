package cmusprocessor

import (
  "fmt"
  "os"
  "log"
  "github.com/TheZoraiz/ascii-image-converter/aic_package"
  "golang.org/x/term"
)

func GetArtPrint() string {
  cmusRaw := GetStatus()

  _, termHeight, err := term.GetSize(int(os.Stdin.Fd()))
  if err != nil {
    log.Fatal(err)
  }

  coverPath, err := ParseStatus(string(cmusRaw))
  if err != nil {
    return "¯\\_(ツ)_/¯\n"
  }

  flags := aic_package.DefaultFlags() 
  flags.Colored = true
  flags.Height = termHeight
  flags.CustomMap = "█"

  // Create the image converter	
  asciiArt, err := aic_package.Convert(coverPath, flags)
  if err != nil {
	  fmt.Println(err)
  }

  return asciiArt
}

