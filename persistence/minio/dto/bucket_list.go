package dto

type Bucket struct {
	Name string `json:"name"`
}

type BucketList struct {
	Buckets []Bucket `json:"buckets"`
}
