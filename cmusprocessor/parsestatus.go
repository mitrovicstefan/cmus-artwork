package cmusprocessor

import (
  "errors"
  "strings"
  "os"
)

func ParseStatus(cmus_input string) (string, error) {
 	lines := strings.Split(cmus_input, "\n")
	var filePath string

  // Process all tags here
	for _, line := range lines {
    if strings.HasPrefix(line, "file ") {
			filePath = line[5:]
		}
	}

  var ext string
  var filePathWithoutExt string

  // Validate data here
	if filePath != "" {
		lastSlashIndex := strings.LastIndex(filePath, "/")
		if lastSlashIndex != -1 {
			filePathWithoutExt = filePath[:lastSlashIndex+1] + "cover"
      foundExt, err := findCoverExtension(filePathWithoutExt)
      if err != nil {
        return "", errors.New("No file 1")
      }
      ext = foundExt
		}
	} else {
    return "", errors.New("No file 2")
  }

	return filePathWithoutExt + ext, nil
}

func findCoverExtension(pathWithoutExtension string) (string, error) {
  extensions := [7]string{".jpg", ".jpeg", ".png", ".bmp", ".webp", ".tiff", ".tif"}

  for _, ext := range extensions {
    if fileExists(pathWithoutExtension + ext) {
      return ext, nil
    }
  }

  return "", errors.New("No file found :)")
}

func fileExists(filename string) bool {
   info, err := os.Stat(filename)
   if os.IsNotExist(err) {
      return false
   }
   return !info.IsDir()
}
