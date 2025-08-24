package middleware

import (
	"net/http"

	"github.com/ccontreras/crispy-potato/internal/adapters/secondary/database/mongodb"
)

// DatabaseMiddleware handles database connection checking
type DatabaseMiddleware struct {
	conn *mongodb.Connection
}

// NewDatabaseMiddleware creates a new DatabaseMiddleware instance
func NewDatabaseMiddleware(conn *mongodb.Connection) *DatabaseMiddleware {
	return &DatabaseMiddleware{
		conn: conn,
	}
}

// CheckConnection ensures database connection is available
func (m *DatabaseMiddleware) CheckConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !m.conn.IsConnected() {
			http.Error(w, "Database connection lost", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
