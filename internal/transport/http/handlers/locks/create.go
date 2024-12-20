package locks

import (
	"context"
	"encoding/json"
	"github.com/rogue0026/door-locker/internal/models"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type LockSaver interface {
	Save(ctx context.Context, lock models.Lock) error
}

func Create(logger *logrus.Logger, locks LockSaver) http.Handler {
	const fn = "internal.transport.http.handlers.locks.Create"
	h := func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		lock := models.Lock{}
		err = json.Unmarshal(reqBody, &lock)
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = locks.Save(r.Context(), lock)
		if err != nil {
			logger.Errorf("%s: %s", fn, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		js, _ := json.Marshal(map[string]interface{}{"message": "entry has been added"})
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(js)
	}
	return http.HandlerFunc(h)
}
