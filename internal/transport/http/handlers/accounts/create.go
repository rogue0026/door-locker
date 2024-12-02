package accounts

import (
	"context"
	"encoding/json"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AccountSaver interface {
	Save(ctx context.Context, account models.Account) error
}

func Create(logger *logrus.Logger, accounts AccountSaver) http.Handler {
	const fn = "internal.transport.http.handlers.accounts.Save"
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
		err = accounts.Save(r.Context(), userAccount)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	return http.HandlerFunc(f)
}
