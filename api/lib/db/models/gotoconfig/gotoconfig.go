package gotoconfig

// ID hashed config, for deduplication
type ID string

// short id for compact / high ec qr codes
type ShortID string

type Config struct {
	ID      ID          `firestore:"id"`
	ShortID ShortID     `firestore:"shortID"`
	Config  interface{} `firestore:"config"`
}

type SnapWidgetConfig struct {
	AppName string             `json:"appName,omitempty" firestore:"appName,omitempty"`
	Wallets []SnapWidgetWallet `json:"wallets,omitempty" firestore:"wallets,omitempty"`
	Intent  string             `json:"intent,omitempty" firestore:"intent,omitempty"`
	Focus   bool               `json:"focus,omitempty" firestore:"focus,omitempty"`
	Theme   map[string]string  `json:"theme,omitempty" firestore:"theme,omitempty"`
	Product *SnapWidgetProduct `json:"product,omitempty" firestore:"product,omitempty"`
}

type SnapWidgetWallet struct {
	Asset   string `json:"asset,omitempty" firestore:"asset,omitempty"`
	Address string `json:"address,omitempty" firestore:"address,omitempty"`
}

type SnapWidgetProduct struct {
	ImageURL           string  `json:"imageURL,omitempty" firestore:"imageURL,omitempty"`
	VideoURL           string  `json:"videoURL,omitempty" firestore:"videoURL,omitempty"`
	DestinationAmount  float64 `json:"destinationAmount,omitempty" firestore:"destinationAmount,omitempty"`
	DestinationTicker  string  `json:"destinationTicker,omitempty" firestore:"destinationTicker,omitempty"`
	DestinationAddress string  `json:"destinationAddress,omitempty" firestore:"destinationAddress,omitempty"`
	Title              string  `json:"title,omitempty" firestore:"title,omitempty"`
	Author             string  `json:"author,omitempty" firestore:"author,omitempty"`
}
