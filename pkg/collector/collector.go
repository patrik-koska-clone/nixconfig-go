package collector

import (
	"fmt"
	"log"
	"os"
)

func PutFilesToDirectory(decodedContents []string, configType string, outputDir string) error {
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	} else {
		log.Println("directory created successfully!")
	}

	for i, fileContent := range decodedContents {

		err := os.WriteFile(fmt.Sprintf("%s/%d-%s",
			outputDir,
			i,
			configType),
			[]byte(fileContent),
			0644)

		if err != nil {
			return err
		}
	}
	return nil
}
