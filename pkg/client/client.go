package client

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func NewContext(apiKey, appKey, site string) context.Context {
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: apiKey,
			},
			"appKeyAuth": {
				Key: appKey,
			},
		},
	)
	ctx = context.WithValue(ctx,
		datadog.ContextServerVariables,
		map[string]string{
			"site": site,
		})
	return ctx
}

func NewAPIClient() *datadog.APIClient {
	configuration := datadog.NewConfiguration()
	return datadog.NewAPIClient(configuration)
}
