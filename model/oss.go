package model

type Oss struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	BucketName      string `json:"bucketName"`
	Url             string `json:"url"`
}
