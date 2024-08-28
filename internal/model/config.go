package model

type S3Config struct {
	Id       string `yaml:"id"`
	Key      string `yaml:"key"`
	Bucket   string `yaml:"bucket"`
	Endpoint string `yaml:"endpoint"`
	Region   string `yaml:"region"`
}
