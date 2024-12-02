package accounts

import (
	"context"
	"encoding/json"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Users interface {
	CreateUserAccount(ctx context.Context, userAccount models.Account) error
	DeleteUserAccount(ctx context.Context, userID int64) error
}

func Login(logger *logrus.Logger, users Users) http.Handler {
	const fn = "internal.transport.http.handlers.accounts.Login"
	h := func(w http.ResponseWriter, r *http.Request) {
		account := models.Account{}
		err := json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}

	return http.HandlerFunc(h)
}
