package locks

import (
	"context"
	"fmt"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/pkg/logging"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

/*
type LockFetcher interface {
	Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error)
	LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error)
}
*/

type NormalMockFetcher struct{}

func (mf NormalMockFetcher) Locks(_ context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	locks := []models.Lock{
		{
			PartNumber:           1,
			Title:                "Test Lock 1",
			Images:               nil,
			Price:                1200,
			SalePrice:            1000,
			Equipment:            "test equipment",
			Colors:               []string{"red", "green", "blue"},
			Description:          "test description",
			Category:             "test category",
			CardMemory:           233,
			Material:             []string{"metal", "carbon"},
			HasMobileApplication: false,
			PowerSupply:          12,
			Size:                 "233 x 12 x 101",
			Weight:               4300,
			DoorType:             []string{"type1", "type2"},
			DoorThicknessMin:     230,
			DoorThicknessMax:     430,
			Rating:               12.2324,
			Quantity:             101,
		},
		{
			PartNumber:           2,
			Title:                "Test Lock 2",
			Images:               nil,
			Price:                1200,
			SalePrice:            1000,
			Equipment:            "test equipment",
			Colors:               []string{"red", "green", "blue"},
			Description:          "test description",
			Category:             "test category",
			CardMemory:           233,
			Material:             []string{"metal", "carbon"},
			HasMobileApplication: false,
			PowerSupply:          12,
			Size:                 "233 x 12 x 101",
			Weight:               4300,
			DoorType:             []string{"type1", "type2"},
			DoorThicknessMin:     230,
			DoorThicknessMax:     430,
			Rating:               12.2324,
			Quantity:             101,
		},
		{
			PartNumber:           3,
			Title:                "Test Lock 3",
			Images:               nil,
			Price:                1200,
			SalePrice:            1000,
			Equipment:            "test equipment",
			Colors:               []string{"red", "green", "blue"},
			Description:          "test description",
			Category:             "test category",
			CardMemory:           233,
			Material:             []string{"metal", "carbon"},
			HasMobileApplication: false,
			PowerSupply:          12,
			Size:                 "233 x 12 x 101",
			Weight:               4300,
			DoorType:             []string{"type1", "type2"},
			DoorThicknessMin:     230,
			DoorThicknessMax:     430,
			Rating:               12.2324,
			Quantity:             101,
		},
	}
	return locks, nil
}

func (mf NormalMockFetcher) LocksByRating(_ context.Context, _ int64) ([]models.Lock, error) {
	return nil, nil
}

type EmptyMockFetcher struct{}

func (mf EmptyMockFetcher) Locks(_ context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error) {
	empty := make([]models.Lock, 0)
	return empty, nil
}

func (mf EmptyMockFetcher) LocksByRating(_ context.Context, _ int64) ([]models.Lock, error) {
	return nil, nil
}

func TestPaginated(t *testing.T) {
	v := url.Values{}
	v.Add("page", "1")
	v.Add("records", "5")
	testLogger := logging.SetupLogger("development", os.Stdout)
	mockFetcher := NormalMockFetcher{}

	paginatedHandler := Paginated(testLogger, mockFetcher)
	testReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/door-locks?%s", v.Encode()), nil)
	resp := httptest.NewRecorder()
	paginatedHandler.ServeHTTP(resp, testReq)
	if resp.Code != http.StatusOK {
		t.Errorf("invalid status code, have %d, want %d", resp.Code, http.StatusOK)
	}
}
