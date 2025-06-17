package config

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func InitMinio() *minio.Client {
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
