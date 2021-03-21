package job

type Job struct {
	ID         string
	Kind       string
	Status     string
	RelatedIDs []string
}
