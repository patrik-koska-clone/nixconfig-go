package collector

import (
	"fmt"
	"log"
	"os"
)

func PutFilesToDirectory(decodedContents []string,
	htmlURLs []string,
	configType string,
	outputDir string,
	apiPageNumber int) error {

	err := os.MkdirAll(fmt.Sprintf("%s/page-%d", outputDir, apiPageNumber), 0755)
	if err != nil {
		return err
	} else {
		log.Println("directory created successfully!")
	}

	for i, fileContent := range decodedContents {

		err := os.WriteFile(fmt.Sprintf("%s/page-%d/%d-%s",
			outputDir,
			apiPageNumber,
			i,
			configType),
			[]byte(fileContent),
			0644)

		if err != nil {
			return err
		}

		nixFileName := fmt.Sprintf("%s/page-%d/%d-%s",
			outputDir,
			apiPageNumber,
			i,
			configType,
		)

		urlToAppend := fmt.Sprintf("\n# appended url: %s\n", htmlURLs[i])

		file, err := os.OpenFile(nixFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := file.WriteString(urlToAppend); err != nil {
			return err
		}
	}
	return nil

}
