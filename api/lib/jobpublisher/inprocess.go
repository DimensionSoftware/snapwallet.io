package jobpublisher

import (
	"github.com/khoerling/flux/api/lib/db"
)

type InProcessPublisher struct {
	*db.Db
}
