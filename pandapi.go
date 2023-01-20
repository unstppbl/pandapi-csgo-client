package pandapi

import (
	"context"
	"time"

	"github.com/unstppbl/pandapi-csgo-client/models"
)

type PandapiClient interface {
	GetRunningMatches(ctx context.Context) (matches []models.Match, err error)
	GetUpcomingMatches(ctx context.Context, period time.Duration) (matches []models.Match, err error)
}
