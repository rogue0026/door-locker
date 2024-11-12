package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/internal/storage"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"github.com/sirupsen/logrus"
	"io"
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
			_, _ = w.Write([]byte("page number must be greater than zero"))
			return
		}

		recordsQuery := r.URL.Query().Get("records")
		recordsOnPage, err := strconv.ParseInt(recordsQuery, 10, 64)
		if err != nil || (recordsOnPage != 10 && recordsOnPage != 20 && recordsOnPage != 50 && recordsOnPage != 5) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("records on page value must be 10, 20 or 50"))
			return
		}
		records, err := locksStorage.LocksWithLimitOffset(r.Context(), pageNumber, recordsOnPage)
		if err != nil {
			if errors.Is(err, storage.ErrRecordsNotFound) {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("records not found"))
			} else {
				logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("internal server error occured"))
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

func AddDoorLockHandler(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	const fn = "internal.transport.http.handlers.AddDoorLockHandler"
	h := func(w http.ResponseWriter, r *http.Request) {
		in, err := io.ReadAll(r.Body)
		if err != nil {
			if errors.Is(err, io.EOF) {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("request body is empty"))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		record := models.DoorLock{}
		err = json.Unmarshal(in, &record)
		if err != nil {
			logger.Error(fmt.Errorf("%s: %w", fn, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = locksStorage.SaveLock(r.Context(), record)
		if err != nil {
			logger.Error(fmt.Errorf("%s: %w", fn, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(map[string]interface{}{"message": "entry has been added"})
		if err != nil {
			logger.Error(fmt.Errorf("%s: %w", fn, err))
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(js)
	}
	return http.HandlerFunc(h)
}
