package locks

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/pkg/logging"
)

type MockFetcher struct {
}

func (mf MockFetcher) Locks(_ context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	return mf.NonEmptyResult()
}

func (mf MockFetcher) LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error) {
	return nil, nil
}

func (mf MockFetcher) EmptyResult() ([]models.Lock, error) {
	return []models.Lock{}, nil
}

func (mf MockFetcher) NonEmptyResult() ([]models.Lock, error) {
	locks := []models.Lock{
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
	}
	return locks, nil
}

type EmptyFetcher struct{}

func (ef EmptyFetcher) LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error) {
	return nil, nil
}

func (ef EmptyFetcher) Locks(_ context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	return []models.Lock{}, nil
}

func TestPaginated(t *testing.T) {
	testLogger := logging.SetupLogger("development", os.Stdout)
	mockFetcher := MockFetcher{}
	tests := []struct {
		name         string
		handler      http.Handler
		pages        []int
		records      []int
		codeExpected int
	}{
		{
			name:         "valid page and records",
			handler:      Paginated(testLogger, mockFetcher),
			pages:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			records:      []int{5, 10, 20, 50},
			codeExpected: http.StatusOK,
		},
		{
			name:         "invalid page",
			handler:      Paginated(testLogger, mockFetcher),
			pages:        []int{0, -1, -2, -3, -1000},
			records:      []int{5, 10, 20, 50},
			codeExpected: http.StatusBadRequest,
		},
		{
			name:         "invalid records number",
			handler:      Paginated(testLogger, mockFetcher),
			pages:        []int{1, 2, 3, 4, 5},
			records:      []int{-5, -10, -20, -50},
			codeExpected: http.StatusBadRequest,
		},
		{
			name:         "empty result from database",
			handler:      Paginated(testLogger, EmptyFetcher{}),
			pages:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			records:      []int{5, 10, 20, 50},
			codeExpected: http.StatusNotFound,
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			for _, pageNum := range curTest.pages {
				for _, recordsNum := range curTest.records {
					queryString := url.Values{}
					queryString.Set("page", strconv.Itoa(pageNum))
					queryString.Set("records", strconv.Itoa(recordsNum))
					testRequest := httptest.NewRequest(http.MethodGet, "/api/door-locks?"+queryString.Encode(), nil)
					resp := httptest.NewRecorder()
					curTest.handler.ServeHTTP(resp, testRequest)
					if resp.Code != curTest.codeExpected {
						t.Errorf("%s: invalid status code, got %d, expect: %d", curTest.name, resp.Code, curTest.codeExpected)
					}
				}
			}
		})
	}
}
