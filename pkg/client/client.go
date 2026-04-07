package client

import (
	"context"

	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func NewContext(apiKey, appKey, site string) context.Context {
	apiKeys := map[string]datadog.APIKey{}
	if apiKey != "" {
		apiKeys["apiKeyAuth"] = datadog.APIKey{Key: apiKey}
	}
	if appKey != "" {
		apiKeys["appKeyAuth"] = datadog.APIKey{Key: appKey}
	}

	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		apiKeys,
	)
	if config.AWSAccessKeyID != "" || config.AWSSecretAccessKey != "" || config.AWSSessionToken != "" {
		ctx = context.WithValue(
			ctx,
			datadog.ContextAWSVariables,
			map[string]string{
				datadog.AWSAccessKeyIdName:     config.AWSAccessKeyID,
				datadog.AWSSecretAccessKeyName: config.AWSSecretAccessKey,
				datadog.AWSSessionTokenName:    config.AWSSessionToken,
			},
		)
	}
	ctx = context.WithValue(ctx, datadog.ContextDelegatedToken, &datadog.DelegatedTokenCredentials{})
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
	if config.CloudAuthOrgUUID != "" {
		configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
			OrgUUID:  config.CloudAuthOrgUUID,
			Provider: datadog.ProviderAWS,
			ProviderAuth: &datadog.AWSAuth{
				AwsRegion: config.CloudAuthAWSRegion,
			},
		}
	}
	return datadog.NewAPIClient(configuration)
}
