package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

/*
	avast/retry-go default configurations:
		attempts			10
		delay				100 ms
		max jitter			100 ms
		on retry			func(n uint, err error) {}
		retry if 			IsRecoverable
		delay type			CombineDelay(BackOffDelay, RandomDelay)
		last error only		false
*/

const (
	existingUrl    = "http://example.com"
	nonExistingUrl = "http://non.example.com"
)

func main() {
	// getInfo(existingUrl)
	getInfo(nonExistingUrl)
}

func getInfo(url string) {
	fmt.Printf("Contacting url %s at %v\n", url, time.Now())

	var body []byte
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)

			if err == nil {
				defer func() {
					err := resp.Body.Close()
					if err != nil {
						panic(err)
					}
				}()

				body, err = ioutil.ReadAll(resp.Body)
			}

			return err
		},
		// options
		retry.Attempts(5),
		retry.Delay(time.Millisecond*500),
		retry.DelayType(retry.FixedDelay),
		retry.OnRetry(func(n uint, err error) {
			// WARN: this function is executed only if main function returns error
			fmt.Printf("Attempt # %d at %v: %s\n", n, time.Now(), err)
		}),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
