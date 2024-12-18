package images

import (
	"net/http"
)

func ImageByName() http.Handler {
	return http.FileServer(http.Dir("images"))
}
