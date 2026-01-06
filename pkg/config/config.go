package config

var (
	ApiKey  string
	AppKey  string
	Site    string
	Debug   bool
	Version = "dev"
)

func GetConfig() (string, string, string) {
	return ApiKey, AppKey, Site
}
