package pandapi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/unstppbl/pandapi-csgo-client/models"
)

const (
	transportDialerTimeout       = time.Second * 40
	transportTlsHandshakeTimeout = time.Second * 15
	transportMaxIdleConns        = 100
	transportMaxConnsPerHost     = 100
	transportMaxIdleConnsPerHost = 5

	ISO8601 = "2006-01-02T15:04:05"
)

type pandapi struct {
	baseUrl string
	token   string
	*http.Client
}

func NewPandapiClient(baseURL string, token string, httpTimeout time.Duration) PandapiClient {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: transportDialerTimeout,
		}).DialContext,
		TLSHandshakeTimeout: transportTlsHandshakeTimeout,
		MaxIdleConns:        transportMaxIdleConns,
		MaxConnsPerHost:     transportMaxConnsPerHost,
		MaxIdleConnsPerHost: transportMaxIdleConnsPerHost,
	}
	cl := &http.Client{
		Timeout:   httpTimeout,
		Transport: transport,
	}
	return &pandapi{
		baseUrl: baseURL,
		token:   token,
		Client:  cl,
	}
}

func (p *pandapi) GetRunningMatches(ctx context.Context) (matches []models.Match, err error) {
	matches = []models.Match{}

	reqUrl := fmt.Sprintf("%s/csgo/matches/running?token=%s", p.baseUrl, p.token)

	_, err = p.serviceRequest(ctx, http.MethodGet, reqUrl, nil, nil, &matches)
	if err != nil {
		return matches, err
	}

	return matches, nil
}

func (p *pandapi) GetUpcomingMatches(ctx context.Context, period time.Duration) (matches []models.Match, err error) {
	matches = []models.Match{}

	fromTime := time.Now().UTC().Format(ISO8601)
	toTime := time.Now().Add(period).UTC().Format(ISO8601)

	reqUrl := fmt.Sprintf("%s/csgo/matches/upcoming?token=%s&range[begin_at]=%sZ,%sZ", p.baseUrl, p.token, fromTime, toTime)

	_, err = p.serviceRequest(ctx, http.MethodGet, reqUrl, nil, nil, &matches)
	if err != nil {
		return matches, err
	}

	return matches, nil
}
