package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rogue0026/door-locker/internal/storage"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func DoorLockByLimitOffsetHandler(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	const fn = "internal.transport.http.handlers.GetDoorLockByLimitOffset"
	h := func(w http.ResponseWriter, r *http.Request) {
		pageQuery := r.URL.Query().Get("page")
		pageNumber, err := strconv.ParseInt(pageQuery, 10, 64)
		if err != nil || pageNumber <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("page number must be greater than zero"))
			return
		}
		recordsQuery := r.URL.Query().Get("records")
		recordsOnPage, err := strconv.ParseInt(recordsQuery, 10, 64)
		if err != nil || (recordsOnPage != 10 && recordsOnPage != 20 && recordsOnPage != 50 && recordsOnPage != 5) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("records on page value must be 10, 20 or 50"))
			return
		}
		records, err := locksStorage.LocksWithLimitOffset(r.Context(), recordsOnPage, pageNumber)
		if err != nil {
			if errors.Is(err, storage.ErrRecordsNotFound) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("records not found"))
			} else {
				logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error occured"))
			}
			return
		}
		jsonData, err := json.MarshalIndent(&records, "", "  ")
		if err != nil {
			logger.Error(fmt.Errorf("%s: %w", fn, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}
	return http.HandlerFunc(h)
}
