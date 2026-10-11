package leofs_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	)

// LeoFSStorage LeoFS对象存储实现 (S3兼容)
type LeoFSStorage struct {
	*s3_storage.S3Storage
}

func NewLeoFSStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*LeoFSStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &LeoFSStorage{S3Storage: s3Storage}, nil
}

const StorageName = "leofs-storage"

func (x *LeoFSStorage) GetName() string {
	return StorageName
}

// Capabilities 声明存储支持的能力
