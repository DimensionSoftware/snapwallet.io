package filemanager

import (
	"bytes"
	"context"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/khoerling/flux/api/lib/db"
	"github.com/khoerling/flux/api/lib/db/models/user"
	"github.com/khoerling/flux/api/lib/db/models/user/file"
	"github.com/khoerling/flux/api/lib/encryption"
	"github.com/khoerling/flux/api/lib/protocol"

	"github.com/lithammer/shortuuid/v3"
)

// Manager manages files
type Manager struct {
	*storage.BucketHandle
	db.Db
	EncryptionManager *encryption.Manager
}

// UploadEncryptedFile handles uploading the encrypted user file to google cloud storage and its metadata to our db
// file metadata is stored in a users' sub collection
func (m Manager) UploadEncryptedFile(ctx context.Context, userID user.ID, req *protocol.UploadFileRequest) (file.ID, error) {
	dekH := encryption.NewDEK()
	dek := encryption.NewEncryptor(dekH)

	now := time.Now()
	fileID := file.ID(shortuuid.New())

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

// GetFile decrypts and returns the file
func (m Manager) GetFile(ctx context.Context, userID user.ID, fileID file.ID) (*file.File, error) {
	md, err := m.Db.GetFileMetadata(ctx, userID, fileID)
	if err != nil {
		return nil, err
	}

	dekH, err := encryption.ParseAndDecryptKeyBytes(md.DataEncryptionKey, m.EncryptionManager.Encryptor)
	if err != nil {
		return nil, err
	}
	dek := encryption.NewEncryptor(dekH)

	obj := m.BucketHandle.Object(string(fileID))

	in, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(in)
	if err != nil {
		return nil, err
	}

	err = in.Close()
	if err != nil {
		return nil, err
	}

	body, err := dek.Decrypt(buf.Bytes(), []byte(userID))
	if err != nil {
		return nil, err
	}

	return &file.File{
		ID:        fileID,
		MimeType:  md.MimeType,
		Size:      md.Size,
		CreatedAt: md.CreatedAt,
		Body:      &body,
	}, nil
}
