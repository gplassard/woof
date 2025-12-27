package config

var (
	ApiKey string
	AppKey string
	Site   string
	Debug  bool
)

func GetConfig() (string, string, string) {
	return ApiKey, AppKey, Site
}
