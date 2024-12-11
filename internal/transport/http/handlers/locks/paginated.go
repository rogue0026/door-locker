package locks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/internal/storage"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type LockFetcher interface {
	Locks(ctx context.Context, pageNumber int64, recordsOnPage int64) ([]models.Lock, error)
	LocksByRating(ctx context.Context, recordsOnPage int64) ([]models.Lock, error)
}

func Paginated(logger *logrus.Logger, fetcher LockFetcher) http.Handler {
	const fn = "internal.transport.http.handlers.Paginated"
	h := func(w http.ResponseWriter, r *http.Request) {
		pageQuery := r.URL.Query().Get("page")
		pageNumber, err := strconv.ParseInt(pageQuery, 10, 64)
		if err != nil || pageNumber <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("page number must be greater than zero"))
			return
		}

		recordsQuery := r.URL.Query().Get("records")
		recordsOnPage, err := strconv.ParseInt(recordsQuery, 10, 64)
		if err != nil || (recordsOnPage != 5 && recordsOnPage != 10 && recordsOnPage != 20 && recordsOnPage != 50) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("records on page value must be 5, 10, 20 or 50"))
			return
		}
		records, err := fetcher.Locks(r.Context(), pageNumber, recordsOnPage)
		if err != nil {
			if errors.Is(err, storage.ErrRecordsNotFound) {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("records not found"))
			} else {
				logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("internal server error occurred"))
			}
			return
		}
		if len(records) == 0 {
			w.WriteHeader(http.StatusNotFound)
			js, _ := json.Marshal(map[string]interface{}{"error": "нет ни одной записи"})
			_, _ = w.Write(js)
			return
		}
		jsonData, err := json.MarshalIndent(&records, "", "  ")
		if err != nil {
			logger.Error(fmt.Errorf("%s: %w", fn, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}
	return http.HandlerFunc(h)
}
