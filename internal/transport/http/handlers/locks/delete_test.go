package locks

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rogue0026/door-locker/pkg/logging"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockRemover struct{}

func (mr MockRemover) Delete(_ context.Context, _ int64) error {
	return nil
}

type ErrMockRemover struct{}

func (mr ErrMockRemover) Delete(_ context.Context, _ int64) error {
	return errors.New("sentinel error")
}

func TestDelete(t *testing.T) {
	testLogger := logging.SetupLogger("development", os.Stdout)
	tests := []struct {
		name         string
		partNumber   string
		handler      http.Handler
		expectedCode int
	}{
		{
			name:         "valid request",
			partNumber:   "1",
			handler:      Delete(testLogger, MockRemover{}),
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid part number",
			partNumber:   "asdfasdf",
			handler:      Delete(testLogger, MockRemover{}),
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "error from database",
			partNumber:   "123",
			handler:      Delete(testLogger, ErrMockRemover{}),
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/door-locks/%s", curTest.partNumber), nil)
			resp := httptest.NewRecorder()
			r := chi.NewRouter()
			r.Handle("/api/door-locks/{PartNumber}", curTest.handler)
			r.ServeHTTP(resp, req)
			if resp.Code != curTest.expectedCode {
				t.Errorf("%s, invalid status code, got %d, expected %d", curTest.name, resp.Code, curTest.expectedCode)
			}
		})

	}
}
