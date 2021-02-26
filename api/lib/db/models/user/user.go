package user

import "time"

// User represents a user registered with our system
type User struct {
	ID        string    `firestore:"id"`
	Email     string    `firestore:"email"`
	Phone     string    `firestore:"phone"`
	CreatedAt time.Time `firestore:"createdAt"`
}
