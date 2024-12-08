package locks

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rogue0026/door-locker/internal/models"
	"github.com/sirupsen/logrus"
)

type CategoryFetcher interface {
	Categories(ctx context.Context) ([]models.Category, error)
}

func Categories(logger *logrus.Logger, locks CategoryFetcher) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		categories, err := locks.Categories(r.Context())
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
