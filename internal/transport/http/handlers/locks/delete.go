package locks

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type LockRemover interface {
	Delete(ctx context.Context, partNumber int64) error
}

func Delete(logger *logrus.Logger, locks LockRemover) http.Handler {
	const fn = "internal.transport.http.handlers.Remove"
	h := func(w http.ResponseWriter, r *http.Request) {
		partNumber, err := strconv.ParseInt(chi.URLParam(r, "PartNumber"), 10, 64)
		if err != nil {
			js, _ := json.Marshal(map[string]interface{}{"error": "invalid part number"})
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(js)
			return
		}
		err = locks.Delete(r.Context(), partNumber)
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
