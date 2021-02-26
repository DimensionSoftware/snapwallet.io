package user

import "time"

// User represents a user registered with our system
type User struct {
	ID        string    `firebase:"id"`
	Email     string    `firebase:"email"`
	Phone     string    `firebase:"phone"`
	CreatedAt time.Time `firebase:"createdAt"`
}
