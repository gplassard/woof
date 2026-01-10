package service_accounts

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListServiceAccountApplicationKeysCmd = &cobra.Command{
	Use:     "list-service-account-application-keys [service_account_id]",
	Aliases: []string{"list-application-keys"},
	Short:   "List application keys for this service account",
	Long: `List application keys for this service account
Documentation: https://docs.datadoghq.com/api/latest/service-accounts/#list-service-account-application-keys`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListApplicationKeysResponse
		var err error

		optionalParams := datadogV2.NewListServiceAccountApplicationKeysOptionalParameters()

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

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceAccountApplicationKeys(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to list-service-account-application-keys")

		fmt.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	ListServiceAccountApplicationKeysCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListServiceAccountApplicationKeysCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListServiceAccountApplicationKeysCmd.Flags().String("sort", "", "Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign.")

	ListServiceAccountApplicationKeysCmd.Flags().String("filter", "", "Filter application keys by the specified string.")

	ListServiceAccountApplicationKeysCmd.Flags().String("filter-created-at-start", "", "Only include application keys created on or after the specified date.")

	ListServiceAccountApplicationKeysCmd.Flags().String("filter-created-at-end", "", "Only include application keys created on or before the specified date.")

	Cmd.AddCommand(ListServiceAccountApplicationKeysCmd)
}
