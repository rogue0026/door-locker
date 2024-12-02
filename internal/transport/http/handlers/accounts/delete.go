package accounts

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AccountRemover interface {
	Remove(ctx context.Context, userID int64) error
}

type DeleteRequest struct {
	UserID int64 `json:"user_id"`
}

func Delete(logger *logrus.Logger, accounts AccountRemover) http.Handler {
	const fn = "internal.transport.http.handlers.Remove"
	f := func(w http.ResponseWriter, r *http.Request) {
		delRequest := DeleteRequest{}
		err := json.NewDecoder(r.Body).Decode(&delRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = accounts.Remove(r.Context(), delRequest.UserID)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	return http.HandlerFunc(f)
}
