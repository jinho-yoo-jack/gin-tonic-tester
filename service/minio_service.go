package service

import (
	"fmt"
	"github.com/jinho-yoo-jack/gin-tonic-tester/dto"
	"github.com/jinho-yoo-jack/gin-tonic-tester/repository"
)

type MinioService struct {
	minioRepository *repository.MinioRepository
}

func NewMinioService(mr *repository.MinioRepository) *MinioService {
	return &MinioService{minioRepository: mr}
}

func (ms *MinioService) GetInfo() {
	fmt.Print(ms.minioRepository.GetMinioClient())
}

func (ms *MinioService) GetBucketList() dto.BucketList {
	return ms.minioRepository.FindAllByBucket()
}
