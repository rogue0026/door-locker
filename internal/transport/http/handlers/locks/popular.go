package locks

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Popular(logger *logrus.Logger, locks LockFetcher) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		recordsQuery := r.URL.Query().Get("records")
		numOfRecords, err := strconv.ParseInt(recordsQuery, 10, 64)
		if err != nil || numOfRecords < 1 {
			w.WriteHeader(http.StatusBadRequest)
			js, _ := json.Marshal(map[string]interface{}{"error": "укажите параметр records"})
			_, _ = w.Write(js)
			return
		}
		data, err := locks.LocksByRating(r.Context(), numOfRecords)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonData, err := json.MarshalIndent(&data, "", "   ")
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
