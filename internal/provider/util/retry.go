package util

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/config"
)

func RetryForTooManyRequests[T any](
	f func() (T, *http.Response, error),
) (T, *http.Response, error) {
	var result T
	var err error
	var httpResponse *http.Response

	for i := 0; i < config.MaxRetriesForTooManyRequests; i++ {
		result, httpResponse, err = f()

		if httpResponse == nil || httpResponse.StatusCode != http.StatusTooManyRequests {
			return result, httpResponse, err
		}

		waitForSeconds := httpResponse.Header.Get("X-Ratelimit-Reset")
		if waitForSeconds == "" {
			return result, httpResponse, errors.Join(
				err,
				errors.New("header X-Ratelimit-Reset is missing or contains non valid value"),
			)
		}

		rateLimitedFor, _ := strconv.Atoi(httpResponse.Header.Get("X-Ratelimit-Reset"))
		time.Sleep(time.Duration(rateLimitedFor) * time.Second)
	}

	return result, httpResponse, errors.Join(
		err,
		errors.New("number of retries exceeded"),
	)
}
