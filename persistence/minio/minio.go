package minio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Minio[T any] struct {
	client *minio.Client
}

func NewMinio() *minio.Client {
	endPoint := "localhost:9000"
	accessKeyID := "dataplatform-developer"
	secretAccessKey := "dataplatform"

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully connected to minio")
	return minioClient
}

func (m *Minio[T]) Save(data T, bucketName, objectName string) (minio.UploadInfo, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return minio.UploadInfo{}, err
	}
	reader := bytes.NewReader(jsonData)
	dataSize := int64(len(jsonData))

	ctx := context.Background()
	uploadInfo, err := m.client.PutObject(ctx, bucketName, objectName, reader, dataSize, minio.PutObjectOptions{})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return uploadInfo, nil
}

func (m *Minio[T]) MakeNewBucket(bucket string) bool {
	err := m.client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		return false
	}
	return true
}

func (m *Minio[T]) IsValidPath(path string) bool {
	if _, err := m.client.BucketExists(context.Background(), path); err != nil {
		fmt.Println(err)
	}
	return true
}
