package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCurrentUserApplicationKeysCmd = &cobra.Command{
	Use: "list-current-user-application-keys",

	Short: "Get all application keys owned by current user",
	Long: `Get all application keys owned by current user
Documentation: https://docs.datadoghq.com/api/latest/key-management/#list-current-user-application-keys`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListApplicationKeysResponse
		var err error

		optionalParams := datadogV2.NewListCurrentUserApplicationKeysOptionalParameters()

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

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCurrentUserApplicationKeys(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-current-user-application-keys")

		fmt.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {

	ListCurrentUserApplicationKeysCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListCurrentUserApplicationKeysCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListCurrentUserApplicationKeysCmd.Flags().String("sort", "", "Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign.")

	ListCurrentUserApplicationKeysCmd.Flags().String("filter", "", "Filter application keys by the specified string.")

	ListCurrentUserApplicationKeysCmd.Flags().String("filter-created-at-start", "", "Only include application keys created on or after the specified date.")

	ListCurrentUserApplicationKeysCmd.Flags().String("filter-created-at-end", "", "Only include application keys created on or before the specified date.")

	ListCurrentUserApplicationKeysCmd.Flags().String("include", "", "Resource path for related resources to include in the response. Only 'owned_by' is supported.")

	Cmd.AddCommand(ListCurrentUserApplicationKeysCmd)
}
