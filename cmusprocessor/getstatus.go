package cmusprocessor

import (
  "fmt"
  "os/exec"
)

func GetStatus() []byte {
  cmd := exec.Command("cmus-remote", "-Q")
  cmusRaw, err := cmd.Output()

  if err != nil {
    if err.Error() == "exit status 1" {
      fmt.Println("Cmus is not running! :)")
    }
    return nil 
  }

  return cmusRaw
}
