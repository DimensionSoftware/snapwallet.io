package user

import "time"

// User represents a user registered with our system
type User struct {
	ID              string    `firestore:"id"`
	Email           string    `firestore:"email"`
	EmailVerifiedAt time.Time `firestore:"emailVerifiedAt"`
	Phone           string    `firestore:"phone"`
	PhoneVerifiedAt time.Time `firestore:"phoneVerifiedAt"`
	CreatedAt       time.Time `firestore:"createdAt"`
}
