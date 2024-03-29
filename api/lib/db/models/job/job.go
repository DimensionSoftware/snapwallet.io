package job

type Kind string

const (
	KindUpdateWyreAccountForUser        Kind = "UPDATE_WYRE_ACCOUNT_FOR_USER"
	KindCreateWyrePaymentMethodsForUser Kind = "CREATE_WYRE_PAYMENT_METHODS_FOR_USER"
)

type Status string

const (
	StatusQueued Status = "QUEUED"
	StatusDone   Status = "DONE"
)

type Job struct {
	ID         string   `json:"id" firestore:"id"`
	Kind       Kind     `json:"kind" firestore:"kind"`
	Status     Status   `json:"status" firestore:"status"`
	RelatedIDs []string `json:"relatedIDs" firestore:"relatedIDs"`
	CreatedAt  int64    `json:"createdAt" firestore:"createdAt"`
	UpdatedAt  int64    `json:"updatedAt" firestore:"updatedAt"`
}
