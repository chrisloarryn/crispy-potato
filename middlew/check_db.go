package middlew

import (
	"net/http"

	"github.com/ccontreras/crispy-potato/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection to DB was lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
