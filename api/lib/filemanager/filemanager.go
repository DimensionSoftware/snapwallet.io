package filemanager

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/protocol"
	"github.com/rs/xid"
)

// Manager manages files
type Manager struct {
	*storage.BucketHandle
	*db.Db
	EncryptionManager *encryption.Manager
}

// UploadEncryptedFile handles uploading the encrypted user file to google cloud storage and its metadata to our db
// file metadata is stored in a users' sub collection
func (m Manager) UploadEncryptedFile(ctx context.Context, userID user.ID, req *protocol.UploadFileRequest) (file.ID, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	now := time.Now()
	fileID := file.ID(xid.New().String())

	obj := m.BucketHandle.Object(string(fileID))
	out := obj.NewWriter(ctx)
	out.ContentType = "application/octet-stream"

	ciphertext, err := dek.Encrypt(req.Body, []byte(userID))
	if err != nil {
		return "", err
	}

	n, err := out.Write(ciphertext)
	if err != nil {
		return "", err
	}
	log.Println("bytes out to gcs:", n)

	err = out.Close()
	if err != nil {
		return "", err
	}

	md := file.Metadata{
		ID:                fileID,
		DataEncryptionKey: *encryption.GetEncryptedKeyBytes(dekH, m.EncryptionManager.Encryptor),
		MimeType:          req.MimeType,
		Size:              req.Size,
		CreatedAt:         now,
	}

	err = m.Db.SaveFileMetadata(ctx, userID, &md)
	if err != nil {
		return "", err
	}

	return fileID, nil
}
