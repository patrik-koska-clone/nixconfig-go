package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DownloadURLResponse struct {
	Items []struct {
		GitUrl  string `json:"git_url"`
		HTMLUrl string `json:"html_url"`
	} `json:"items"`
}

type ContentURLResponse struct {
	Content string
}

func GetNixConfigURLs(URL string, token string) ([]string, []string, error) {
	var (
		dlur         DownloadURLResponse
		downloadURLs []string
		htmlURLs     []string
		req          *http.Request
		err          error
		client       *http.Client
		resp         *http.Response
		bodyBytes    []byte
	)

	req, err = http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return downloadURLs, htmlURLs, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	client = &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return downloadURLs, htmlURLs, err
	}
	defer resp.Body.Close()

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return downloadURLs, htmlURLs, err
	}

	err = json.Unmarshal(bodyBytes, &dlur)
	if err != nil {
		return downloadURLs, htmlURLs, err
	}

	for _, dlu := range dlur.Items {
		downloadURLs = append(downloadURLs, dlu.GitUrl)
		htmlURLs = append(htmlURLs, dlu.HTMLUrl)
	}

	return downloadURLs, htmlURLs, nil
}

func GetNixConfigContents(downloadURLs []string, token string) ([]string, error) {
	var (
		contents      []string
		contentResp   ContentURLResponse
		maxConcurrent int
		semaphore     chan struct{}
		errChan       chan error
		cancel        context.CancelFunc
		ctx           context.Context
		req           *http.Request
		err           error
		client        *http.Client
		resp          *http.Response
		bodyBytes     []byte
		finalErr      error
	)

	maxConcurrent = 5
	semaphore = make(chan struct{}, maxConcurrent)

	contents = []string{}
	errChan = make(chan error) // Channel for collecting errors

	for _, url := range downloadURLs {
		semaphore <- struct{}{} // Acquire a semaphore slot
		go func(url string) {
			defer func() { <-semaphore }() // Release slot on completion

			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel() // Ensure timeout goroutines also release a token

			req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				err = fmt.Errorf("error creating new request\n%v", err)
				errChan <- err
				return
			}

			req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

			client = &http.Client{}

			resp, err = client.Do(req)
			if err != nil {
				err = fmt.Errorf("error sending HTTP request\n%v", err)
				errChan <- err
				return
			}
			defer resp.Body.Close()

			bodyBytes, err = io.ReadAll(resp.Body)
			if err != nil {
				err = fmt.Errorf("error reading response body\n%v", err)
				errChan <- err
				return
			}

			err = json.Unmarshal(bodyBytes, &contentResp)
			if err != nil {
				err = fmt.Errorf("error parsing json to struct\n%v", err)
				errChan <- err
				return
			}

			contents = append(contents, contentResp.Content)
		}(url) // Pass the URL to the goroutine's closure
	}

	// Block and wait for all goroutines to complete
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}

	for i := 0; i < len(downloadURLs); i++ {
		select {
		case err := <-errChan:
			finalErr = fmt.Errorf("error during http requests: %w", err)
		default:
			// No error received
		}
	}

	return contents, finalErr
}
