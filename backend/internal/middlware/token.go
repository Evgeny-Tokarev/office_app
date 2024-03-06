package middlware

import (
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/services/userservice"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func TokenMiddleware(us *userservice.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && (r.URL.Path == "/user" || r.URL.Path == "/user/login") {
				next.ServeHTTP(w, r)
				return
			}
			fields := strings.Fields(r.Header.Get("Authorization"))
			if len(fields) < 2 {
				util.SendTranscribedError(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != "bearer" {
				util.SendTranscribedError(w, fmt.Sprintf("unsupported authorization type %s", authorizationType), http.StatusUnauthorized)
				return
			}
			_, err := us.TokenMaker.VerifyToken(fields[1])
			if err != nil {
				fmt.Println("Error parsing token:", err)
				util.SendTranscribedError(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
