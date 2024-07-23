package page_requestor

import (
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type transport_struct struct {
	func_wrapper http.RoundTripper
}

func mal_client() (mal_client *http.Client) {
	mal_client = &http.Client{Transport: transport_struct{func_wrapper: http.DefaultTransport}}

	return mal_client
}

func (the_current_struct transport_struct) RoundTrip(http_request *http.Request) (response *http.Response, err error) {
	const sleep_period time.Duration = 60 * time.Second
	var attempts int = 10

	for range attempts {
		response, err = the_current_struct.func_wrapper.RoundTrip(http_request)
		if err != nil {
			log.Error().Err(err).Msg("Can't get the page")
			return nil, err
		}

		code := response.StatusCode

		if code == 200 || code == 404 {
			return response, nil
		}

		time.Sleep(sleep_period)
	}

	err = errors.New("max retries reaches") // breakout
	response.Body.Close()                   // close the request
	return response, err
}
