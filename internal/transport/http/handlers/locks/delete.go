package locks

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LockRemover interface {
	Delete(ctx context.Context, partNumber string) error
}

type DeleteRequest struct {
	PartNumber string `json:"part_number"`
}

func Delete(logger *logrus.Logger, locks LockRemover) http.Handler {
	const fn = "internal.transport.http.handlers.Remove"
	h := func(w http.ResponseWriter, r *http.Request) {
		delReq := DeleteRequest{}
		err := json.NewDecoder(r.Body).Decode(&delReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = locks.Delete(r.Context(), delReq.PartNumber)
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
