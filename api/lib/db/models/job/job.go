package job

type Kind string

const (
	KindCreateWyreAccountForUser Kind = "CREATE_WYRE_ACCOUNT_FOR_USER"
)

type Status string

const (
	StatusQueued Status = "QUEUED"
	StatusDone   Status = "DONE"
)

type Job struct {
	ID         string   `json:"id"`
	Kind       Kind     `json:"kind"`
	Status     Status   `json:"status"`
	RelatedIDs []string `json:"relatedIDs"`
	CreatedAt  int64    `json:"createdAt"`
	UpdatedAt  int64    `json:"updatedAt"`
}
