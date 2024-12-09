package locks

import (
	"context"
	"github.com/rogue0026/door-locker/internal/models"
	"net/http"
	"testing"
)

/*
	type LockFetcher interface {
		Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error)
		LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error)
	}
*/
type MockFetcher struct{}

func (mf MockFetcher) Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
}

func (mf MockFetcher) LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error) {
}

func TestPopular(t *testing.T) {
	tests := []struct {
		name         string
		numOfRecords int
		expectedCode int
	}{
		{
			name:         "valid request",
			numOfRecords: 10,
			expectedCode: http.StatusOK,
		},
	}
}
