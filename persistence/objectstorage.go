package persistence

type ObjectStorage[T any] interface {
	Save(data T, bucketName, objectName string) (bool, error)
	FindById(id string) T
	FindAll() []T
	IsValidPath(path string) bool
}
