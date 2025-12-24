package util

var (
	ApiKey string
	AppKey string
	Site   string
)

func GetConfig() (string, string, string) {
	return ApiKey, AppKey, Site
}
