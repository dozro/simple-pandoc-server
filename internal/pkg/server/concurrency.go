package server

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func conc_fromCache(ctx context.Context, data []byte) ([]byte, error) {
	log.Debug("starting concurrent cache lookup")
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	// Run the actual work in a goroutine
	go func() {
		out, err := fromCache(data)
		if err != nil {
			log.Debugf("cache lookup failed: %s", err)
			errChan <- err
			return
		}
		resultChan <- out
	}()
	// Wait for completion or cancellation
	select {
	case <-ctx.Done():
		log.Debug("canceling cache lookup")
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	case res := <-resultChan:
		return res, nil
	}
}

func concurrentCacheLookupAndRendering(ctx context.Context, data []byte, convOp func([]byte) ([]byte, error)) ([]byte, error) {
	if convOp == nil {
		return nil, fmt.Errorf("concurrentCacheLookupAndRendering requires concurrency function")
	}
	const numWork = 2
	resultChan := make(chan conc_result, numWork)
	ctx1, cancelConvert := context.WithCancel(ctx)
	defer cancelConvert()
	ctx2, cancelLookup := context.WithCancel(ctx)
	defer cancelLookup()

	go func() {
		data, err := conc_conversionWrapperFunc(ctx1, data, convOp)
		resultChan <- conc_result{data: data, err: err}
	}()
	go func() {
		data, err := conc_fromCache(ctx2, data)
		resultChan <- conc_result{data: data, err: err}
	}()
	var firstErr error

	for i := 0; i < numWork; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-resultChan:
			if res.err == nil && res.data != nil {
				log.Debug("concurrentCacheLookupAndRendering finished successfully")
				cancelConvert()
				cancelLookup()
				go toCache(data, res.data)
				return res.data, res.err
			}
			if firstErr == nil {
				log.Error(res.err)
				firstErr = res.err
			}
		}
	}
	return nil, fmt.Errorf("both tasks failed, first error: %v", firstErr)
}

type conc_result struct {
	data []byte
	err  error
}

func conc_conversionWrapperFunc(ctx context.Context, data []byte, convOp func([]byte) ([]byte, error)) ([]byte, error) {
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	go func() {
		out, err := convOp(data)
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- out
	}()
	// Wait for completion or cancellation
	select {
	case <-ctx.Done():
		log.Debug("canceling concurrent cache lookup")
		return nil, ctx.Err()
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		return nil, err
	}
}
