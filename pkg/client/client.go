package client

import (
	"context"

	"github.com/gplassard/woof/pkg/config"

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
	ctx = context.WithValue(ctx, datadog.ContextServerIndex, 1)
	ctx = context.WithValue(ctx,
		datadog.ContextServerVariables,
		map[string]string{
			"name": site,
			"site": site,
		})
	return ctx
}

func NewAPIClient() *datadog.APIClient {
	configuration := datadog.NewConfiguration()
	configuration.Debug = config.Debug
	if config.EnableUnstableOperations {
		for _, operation := range configuration.GetUnstableOperations() {
			if operation == "" {
				continue
			}
			configuration.SetUnstableOperationEnabled(operation, true)
		}
	}
	return datadog.NewAPIClient(configuration)
}
