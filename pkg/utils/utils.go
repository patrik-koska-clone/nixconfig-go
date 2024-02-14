package utils

import (
	"encoding/base64"
)

func DecodeFromB64(contents []string) ([]string, error) {
	var decodedContents []string

	for _, b64Content := range contents {
		decodedBytes, err := base64.StdEncoding.DecodeString(b64Content)
		if err != nil {
			return decodedContents, err
		}

		decodedString := string(decodedBytes)
		decodedContents = append(decodedContents, decodedString)
	}

	return decodedContents, nil

}
