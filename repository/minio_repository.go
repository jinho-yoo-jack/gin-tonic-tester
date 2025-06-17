package repository

import (
	"context"
	"fmt"
	"github.com/jinho-yoo-jack/gin-tonic-tester/dto"
	"github.com/minio/minio-go/v7"
)

type MinioRepository struct {
	minioClient *minio.Client
}

func NewMinioRepository(mc *minio.Client) *MinioRepository {
	return &MinioRepository{minioClient: mc}
}

func (mr *MinioRepository) GetMinioClient() *minio.Client {
	return mr.minioClient
}

func (mr *MinioRepository) FindAllByBucket() dto.BucketList {
	var result dto.BucketList
	buckets, err := mr.minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return result
	}
	for _, bucket := range buckets {
		result.Buckets = append(result.Buckets, dto.Bucket{Name: bucket.Name})
	}
	return result
}
