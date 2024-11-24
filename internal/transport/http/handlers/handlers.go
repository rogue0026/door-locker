package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/rogue0026/door-locker/internal/models"
	"github.com/rogue0026/door-locker/internal/storage"
	"github.com/rogue0026/door-locker/internal/storage/postgres"
	"github.com/sirupsen/logrus"
)

func DoorLockByLimitOffsetHandler(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	const fn = "internal.transport.http.handlers.DoorLockByLimitOffsetHandler"
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
		records, err := locksStorage.LocksLimitOffset(r.Context(), pageNumber, recordsOnPage)
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

func PopularDoorLocks(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	// http:hostname:port/api/door-locks/popular?records=records_number
	const fn = "internal.transport.http.handlers.PopularDoorLocks"
	h := func(w http.ResponseWriter, r *http.Request) {
		recordsQuery := r.URL.Query().Get("records")
		numOfRecords, err := strconv.ParseInt(recordsQuery, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		records, err := locksStorage.LocksOrderedByRating(r.Context(), numOfRecords)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonData, err := json.MarshalIndent(&records, "", "   ")
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}
	return http.HandlerFunc(h)
}

func DoorLocksCategories(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	// http://hostname:port/api/door-locks/categories
	const fn = "internal.transport.http.handlers.DoorLocksCategories"
	h := func(w http.ResponseWriter, r *http.Request) {
		categories, err := locksStorage.LocksCategories(r.Context())
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonData, err := json.MarshalIndent(&categories, "", "   ")
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}
	return http.HandlerFunc(h)
}

func AddDoorLock(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	const fn = "internal.transport.http.handlers.AddDoorLock"
	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
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
		err = json.Unmarshal(reqBody, &record)
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = locksStorage.SaveLock(r.Context(), record)
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(map[string]interface{}{"message": "entry has been added"})
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(js)
	}
	return http.HandlerFunc(h)
}

func DeleteDoorLock(logger *logrus.Logger, locksStorage *postgres.Storage) http.Handler {
	const fn = "internal.transport.http.handlers.DeleteDoorLock"
	type DeleteRequest struct {
		PartNumber string `json:"part_number"`
	}
	h := func(w http.ResponseWriter, r *http.Request) {
		delReq := DeleteRequest{}
		err := json.NewDecoder(r.Body).Decode(&delReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = locksStorage.DeleteLockByPartNumber(r.Context(), delReq.PartNumber)
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		// we are ignoring error, because it's never raise here
		js, _ := json.Marshal(map[string]string{"message": "entry has been deleted"})
		_, _ = w.Write(js)
	}
	return http.HandlerFunc(h)
}

func RegisterAccount(logger *logrus.Logger, storage *postgres.Storage) http.Handler {
	const fn = "internal.transport.http.handlers.RegisterAccount"
	f := func(w http.ResponseWriter, r *http.Request) {
		userAccount := models.Account{}
		err := json.NewDecoder(r.Body).Decode(&userAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		parsedBirthDate, err := time.Parse("02.01.2006", userAccount.BirthDate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userAccount.BirthDate = parsedBirthDate.Format("2006-01-02")
		err = storage.CreateUserAccount(r.Context(), userAccount)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	return http.HandlerFunc(f)
}

func DeleteAccount(logger *logrus.Logger, storage *postgres.Storage) http.Handler {
	type DeleteRequest struct {
		UserID int64 `json:"user_id"`
	}
	const fn = "internal.transport.http.handlers.DeleteAccount"
	f := func(w http.ResponseWriter, r *http.Request) {
		delRequest := DeleteRequest{}
		err := json.NewDecoder(r.Body).Decode(&delRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = storage.DeleteUserAccount(r.Context(), delRequest.UserID)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	return http.HandlerFunc(f)
}
