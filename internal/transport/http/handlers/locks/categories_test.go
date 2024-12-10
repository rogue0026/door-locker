package locks

import (
	"context"
	"errors"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/pkg/logging"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockCategoriesFetcher struct{}

func (mf MockCategoriesFetcher) Categories(_ context.Context) ([]models.Category, error) {
	categories := []models.Category{
		{
			ID:    1,
			Name:  "For home",
			Image: nil,
		},
		{
			ID:    2,
			Name:  "For garage",
			Image: nil,
		},
		{
			ID:    3,
			Name:  "Electric",
			Image: nil,
		},
	}
	return categories, nil
}

type ErrMockCategoriesFetcher struct{}

func (mf ErrMockCategoriesFetcher) Categories(_ context.Context) ([]models.Category, error) {
	return nil, errors.New("sentinel error")
}

func TestCategories(t *testing.T) {
	testLogger := logging.SetupLogger("development", os.Stdout)

	tests := []struct {
		name         string
		expectedCode int
		handler      http.Handler
	}{
		{
			name:         "valid request",
			expectedCode: http.StatusOK,
			handler:      Categories(testLogger, MockCategoriesFetcher{}),
		},
		{
			name:         "error from database",
			expectedCode: http.StatusInternalServerError,
			handler:      Categories(testLogger, ErrMockCategoriesFetcher{}),
		},
	}

	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/door-locks/categories", nil)
			resp := httptest.NewRecorder()
			curTest.handler.ServeHTTP(resp, req)
			if resp.Code != curTest.expectedCode {
				t.Errorf("invalid status code, expected %d, got %d", curTest.expectedCode, resp.Code)
			}
		})
	}
}
