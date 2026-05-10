package store
import (
	"context" 
	"database/sql"
)
type Usersstore struct {
db *sql.DB
}
func (s *Usersstore) Create(ctx context.Context) error {
	// Implementation for creating a user
	return nil
}
