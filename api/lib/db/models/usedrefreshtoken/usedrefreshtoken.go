package usedrefreshtoken

import "time"

// UsedRefreshToken ... used to enforce one time use of Refresh token (RTR: Refresh Token Rotation)
type UsedRefreshToken struct {
	ID        string    `firestore:"id"`        // jti of the refresh token (id)
	Subject   string    `firestore:"subject"`   // for auditing
	IssuedAt  time.Time `firestore:"issuedAt"`  // for auditing
	ExpiresAt time.Time `firestore:"expiresAt"` // for auditing & if we know when it expires then we can clean it up once its no long granting access
	UsedAt    time.Time `firestore:"usedAt"`    // for auditing
}
