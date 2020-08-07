package v1

import (
	"net/http"

	"github.com/cqbqdd11519/apiserver-test/internal/utils"
)

func Authorize(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.TLS == nil || len(req.TLS.PeerCertificates) == 0 {
			_ = utils.RespondError(w, http.StatusBadRequest, "is not https or there is no peer certificate")
			return
		}

		h.ServeHTTP(w, req)
	})
}
