package locks

import (
	"context"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/pkg/logging"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

/*
	type LockFetcher interface {
		Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error)
		LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error)
	}
*/

type MockPopularFetcher struct{}

func (mf MockPopularFetcher) Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	return nil, nil
}

func (mf MockPopularFetcher) LocksByRating(ctx context.Context, _ int64) ([]models.Lock, error) {
	return []models.Lock{
		{
			Title:                "TestLock",
			Images:               []string{"test_image1.png", "test_image2.png"},
			Price:                12000,
			SalePrice:            10100,
			Equipment:            "Lock, fingerprint module, installation guide",
			Colors:               []string{"gold", "white", "black"},
			Description:          "Standard lock",
			Category:             "For home",
			CardMemory:           200,
			Material:             []string{"metal"},
			HasMobileApplication: true,
			PowerSupply:          12,
			Size:                 "120 x 220 x 10",
			Weight:               2000,
			DoorType:             []string{"garage"},
			DoorThicknessMin:     100,
			DoorThicknessMax:     200,
			Rating:               4.23,
			Quantity:             100,
		},
		{
			Title:                "TestLock1",
			Images:               []string{"test_image1.png", "test_image2.png"},
			Price:                13000,
			SalePrice:            12100,
			Equipment:            "Lock, fingerprint module, installation guide",
			Colors:               []string{"gold", "white", "black"},
			Description:          "Standard lock",
			Category:             "For home",
			CardMemory:           200,
			Material:             []string{"metal"},
			HasMobileApplication: true,
			PowerSupply:          12,
			Size:                 "120 x 220 x 10",
			Weight:               2000,
			DoorType:             []string{"garage"},
			DoorThicknessMin:     100,
			DoorThicknessMax:     200,
			Rating:               4.23,
			Quantity:             100,
		},
	}, nil
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
		{
			name:         "bad records parameter",
			numOfRecords: -100,
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			query := url.Values{}
			query.Set("records", strconv.Itoa(curTest.numOfRecords))
			req := httptest.NewRequest(http.MethodGet, "/api/door-locks/popular?"+query.Encode(), nil)
			resp := httptest.NewRecorder()
			handler := Popular(logging.SetupLogger("development", os.Stdout), MockPopularFetcher{})
			handler.ServeHTTP(resp, req)
			if resp.Code != curTest.expectedCode {
				t.Errorf("invalid response status code, expected %d, got %d", curTest.expectedCode, resp.Code)
			}
		})
	}
}
