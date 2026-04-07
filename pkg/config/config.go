package config

var (
	ApiKey             string
	AppKey             string
	Site               string
	Debug              bool
	CloudAuthOrgUUID   string
	CloudAuthAWSRegion string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSSessionToken    string
	Version            = "dev"
)

func GetConfig() (string, string, string) {
	return ApiKey, AppKey, Site
}
