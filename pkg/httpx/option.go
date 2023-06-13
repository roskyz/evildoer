package httpx

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/guonaihong/gout/middler"
)

type Option func(r *requestMaker)

func WithBackOff(f BackOff) Option {
	return func(r *requestMaker) {
		r.backOffFunc = f
	}
}

func WithRetryTimes(maxTimes int) Option {
	return func(r *requestMaker) {
		r.retryMaxTimes = maxTimes
	}
}

func WithTimeout(duration time.Duration) Option {
	return func(r *requestMaker) {
		r.timeout = duration
	}
}

type RequestResult struct {
	StateCode int
	Header    http.Header
	BodyBytes []byte
}

func (r *RequestResult) Use() middler.ResponseMiddler {
	return middler.WithResponseMiddlerFunc(func(response *http.Response) (err error) {
		r.StateCode = response.StatusCode
		r.Header = response.Header
		if r.BodyBytes, err = io.ReadAll(response.Body); err != nil {
			return err
		}
		if err = response.Body.Close(); err != nil {
			return err
		}
		response.Body = io.NopCloser(bytes.NewBuffer(r.BodyBytes))
		return
	})
}
