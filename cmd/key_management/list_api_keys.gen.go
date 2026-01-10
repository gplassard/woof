package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAPIKeysCmd = &cobra.Command{
	Use: "list-api-keys",

	Short: "Get all API keys",
	Long: `Get all API keys
Documentation: https://docs.datadoghq.com/api/latest/key-management/#list-api-keys`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.APIKeysResponse
		var err error

		optionalParams := datadogV2.NewListAPIKeysOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		if cmd.Flags().Changed("filter-created-at-start") {
			val, _ := cmd.Flags().GetString("filter-created-at-start")
			optionalParams.WithFilterCreatedAtStart(val)
		}

		if cmd.Flags().Changed("filter-created-at-end") {
			val, _ := cmd.Flags().GetString("filter-created-at-end")
			optionalParams.WithFilterCreatedAtEnd(val)
		}

		if cmd.Flags().Changed("filter-modified-at-start") {
			val, _ := cmd.Flags().GetString("filter-modified-at-start")
			optionalParams.WithFilterModifiedAtStart(val)
		}

		if cmd.Flags().Changed("filter-modified-at-end") {
			val, _ := cmd.Flags().GetString("filter-modified-at-end")
			optionalParams.WithFilterModifiedAtEnd(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("filter-remote-config-read-enabled") {
			val, _ := cmd.Flags().GetString("filter-remote-config-read-enabled")
			optionalParams.WithFilterRemoteConfigReadEnabled(val)
		}

		if cmd.Flags().Changed("filter-category") {
			val, _ := cmd.Flags().GetString("filter-category")
			optionalParams.WithFilterCategory(val)
		}

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAPIKeys(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-api-keys")

		fmt.Println(cmdutil.FormatJSON(res, "api_keys"))
	},
}

func init() {

	ListAPIKeysCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListAPIKeysCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListAPIKeysCmd.Flags().String("sort", "", "API key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign.")

	ListAPIKeysCmd.Flags().String("filter", "", "Filter API keys by the specified string.")

	ListAPIKeysCmd.Flags().String("filter-created-at-start", "", "Only include API keys created on or after the specified date.")

	ListAPIKeysCmd.Flags().String("filter-created-at-end", "", "Only include API keys created on or before the specified date.")

	ListAPIKeysCmd.Flags().String("filter-modified-at-start", "", "Only include API keys modified on or after the specified date.")

	ListAPIKeysCmd.Flags().String("filter-modified-at-end", "", "Only include API keys modified on or before the specified date.")

	ListAPIKeysCmd.Flags().String("include", "", "Comma separated list of resource paths for related resources to include in the response. Supported resource paths are 'created_by' and 'modified_by'.")

	ListAPIKeysCmd.Flags().String("filter-remote-config-read-enabled", "", "Filter API keys by remote config read enabled status.")

	ListAPIKeysCmd.Flags().String("filter-category", "", "Filter API keys by category.")

	Cmd.AddCommand(ListAPIKeysCmd)
}
