package httpx

import (
	"time"

	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/middler"
)

type Requester interface {
	SendJSON(reqOpt *RequestOptions, options ...Option) error
}

type requestMaker struct {
	retryMaxTimes int
	backOffFunc   BackOff
	timeout       time.Duration
}

func New(options ...Option) Requester {
	return new(requestMaker).rebuildRequestMaker(options)
}

func (maker *requestMaker) rebuildRequestMaker(options []Option) *requestMaker {
	request := &requestMaker{
		retryMaxTimes: maker.retryMaxTimes,
		backOffFunc:   maker.backOffFunc,
		timeout:       maker.timeout,
	}
	for _, option := range options {
		option(request)
	}
	return request
}

func (maker *requestMaker) SendJSON(reqOpt *RequestOptions, options ...Option) error {
	// set header json append option
	return maker.DoRequest(reqOpt, options...)
}

func (maker *requestMaker) DoRequest(reqOpt *RequestOptions, options ...Option) error {
	// clone the request with option
	r := maker.rebuildRequestMaker(options)
	client := gout.New().SetMethod(reqOpt.HttpMethod).SetURL(reqOpt.DestinationURL)

	// set timeout and context
	if r.timeout > 0 {
		client = client.SetTimeout(r.timeout)
	}
	if reqOpt.Ctx != nil {
		client = client.WithContext(reqOpt.Ctx)
	}

	// set query
	if reqOpt.Queries != nil {
		client = client.SetQuery(convArray(reqOpt.Queries))
	}

	// set headers
	if reqOpt.Headers != nil {
		client = client.SetHeader(convArray(reqOpt.Headers))
	}

	// set request body
	if reqOpt.RequestBody != nil {
		client = client.SetBody(reqOpt.RequestBody)
	}

	// set user's response middler
	if reqOpt.ResponseMiddler != nil {
		client = client.ResponseUse(middler.WithResponseMiddlerFunc(reqOpt.ResponseMiddler))
	}

	// set response binder only support JSON todo: detect from header content-type
	if reqOpt.BodyBinder != nil {
		client = client.BindJSON(reqOpt.BodyBinder)
	}

	// set retry policy
	retry := client.F().Retry()
	if r.retryMaxTimes > 0 {
		retry = retry.Attempt(r.retryMaxTimes)
	}
	if r.backOffFunc != nil {
		retry = retry.Func(convertToGoutRetryFunc(r.backOffFunc))
	}

	return retry.Do()
}
