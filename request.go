package pandapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	acceptHeader      = "Accept"
	authHeader        = "Authorization"
	contentTypeHeader = "Content-Type"

	jsonType = "application/json"
)

var (
	errResponseBodyIsEmpty   = errors.New("response body is empty")
	errUnmarshalToValueError = errors.New("cannot unmarshal into value")
)

func (p *pandapi) serviceRequest(ctx context.Context, method string, url string, headers http.Header, payload []byte, value interface{}) (int, error) {
	var bodyReader io.Reader
	if payload != nil {
		bodyReader = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return 0, errors.Wrap(err, "cannot make request")
	}

	req.Header = p.addDefaultHeaders(headers)

	resp, err := p.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "error making http call")
	}
	defer resp.Body.Close()

	if !isStatusCodeOK(resp.StatusCode) {
		err := errors.Errorf("%d: unknown error", resp.StatusCode)
		return resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "cannot read response body")
		return 0, err
	}
	if len(body) == 0 {
		return resp.StatusCode, errResponseBodyIsEmpty
	}

	err = json.Unmarshal(body, value)
	if err != nil {
		return resp.StatusCode, errUnmarshalToValueError
	}

	return resp.StatusCode, nil
}

func (p *pandapi) addDefaultHeaders(headers http.Header) http.Header {
	if headers == nil {
		headers = http.Header{}
	}

	if _, exists := headers[contentTypeHeader]; !exists {
		headers.Set(contentTypeHeader, jsonType)
	}

	if _, exists := headers[acceptHeader]; !exists {
		headers.Set(acceptHeader, jsonType)
	}

	return headers
}

func isStatusCodeOK(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
