package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/exp/slog"
)

var bucketName = "superfile"

type MinIO struct {
	client *minio.Client
}

func MinIOConnect() (*MinIO, error) {
	endpoint := "172.24.0.2:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin123"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		slog.Error("Failed to connect to MinIO", "error", err)
		return nil, err
	}

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			slog.Info("Bucket already exists", "bucket", bucketName)
		} else {
			slog.Error("Failed to create bucket", "error", err)
			return nil, err
		}
	}

	slog.Info("Connected to MinIO", "bucket", bucketName)
	return &MinIO{client: minioClient}, nil
}

func (m *MinIO) Upload(fileName, filePath string) (string, error) {
	_, err := m.client.FPutObject(context.Background(), bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		slog.Error("Failed to upload file", "error", err, "file", fileName)
		return "", err
	}

	minioURL := fmt.Sprintf("http://localhost:9000/%s/%s", bucketName, fileName)
	slog.Info("File uploaded to MinIO", "file", fileName, "url", minioURL)
	return minioURL, nil
}

func (m *MinIO) Download(fileName, destPath string) error {
	err := m.client.FGetObject(context.Background(), bucketName, fileName, destPath, minio.GetObjectOptions{})
	if err != nil {
		slog.Error("Failed to download file", "error", err, "file", fileName)
		return err
	}

	slog.Info("File downloaded from MinIO", "file", fileName)
	return nil
}

func (m *MinIO) ListFiles() ([]string, error) {
	objectCh := m.client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})

	var files []string
	for object := range objectCh {
		if object.Err != nil {
			slog.Error("Error listing object", "error", object.Err)
			return nil, object.Err
		}
		files = append(files, object.Key)
	}
	slog.Info("Listed files from MinIO", "count", len(files))
	return files, nil
}