package images

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	basePath = "./images/"
)

type ImageFetcher interface{}

func ImageByName(logger *logrus.Logger) http.Handler {
	const fn = "internal.transport.http.handlers.images.ImageByName"
	h := func(w http.ResponseWriter, r *http.Request) {
		imgName := chi.URLParam(r, "ImageName")
		file, err := os.Open(filepath.Join(basePath, imgName))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				js, _ := json.Marshal(map[string]interface{}{"err": "file not found"})
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write(js)
				return
			}
			logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(encoded, data)
		js, _ := json.MarshalIndent(map[string]interface{}{"image": string(encoded)}, "", "   ")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(js)
	}
	return http.HandlerFunc(h)
}
