//go:generate ffjson $GOFILE
package user

import "time"

// User is a user.
type User struct {
	ID      int
	Name    string
	Age     int
	Created time.Time
}
