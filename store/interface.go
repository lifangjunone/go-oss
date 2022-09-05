package store

// Storer define how to upload file to bucket
// Storer abstract, mask difference
type Storer interface {
	Upload(bucketName string, objectKey string, fileName string) error
}
