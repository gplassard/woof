package app_builder

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAppsCmd = &cobra.Command{
	Use: "list-apps",

	Short: "List Apps",
	Long: `List Apps
Documentation: https://docs.datadoghq.com/api/latest/app-builder/#list-apps`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListAppsResponse
		var err error

		optionalParams := datadogV2.NewListAppsOptionalParameters()

		if cmd.Flags().Changed("limit") {
			val, _ := cmd.Flags().GetInt64("limit")
			optionalParams.WithLimit(val)
		}

		if cmd.Flags().Changed("page") {
			val, _ := cmd.Flags().GetInt64("page")
			optionalParams.WithPage(val)
		}

		if cmd.Flags().Changed("filter-user-name") {
			val, _ := cmd.Flags().GetString("filter-user-name")
			optionalParams.WithFilterUserName(val)
		}

		if cmd.Flags().Changed("filter-user-uuid") {
			val, _ := cmd.Flags().GetString("filter-user-uuid")
			optionalParams.WithFilterUserUuid(val)
		}

		if cmd.Flags().Changed("filter-name") {
			val, _ := cmd.Flags().GetString("filter-name")
			optionalParams.WithFilterName(val)
		}

		if cmd.Flags().Changed("filter-query") {
			val, _ := cmd.Flags().GetString("filter-query")
			optionalParams.WithFilterQuery(val)
		}

		if cmd.Flags().Changed("filter-deployed") {
			val, _ := cmd.Flags().GetString("filter-deployed")
			optionalParams.WithFilterDeployed(val)
		}

		if cmd.Flags().Changed("filter-tags") {
			val, _ := cmd.Flags().GetString("filter-tags")
			optionalParams.WithFilterTags(val)
		}

		if cmd.Flags().Changed("filter-favorite") {
			val, _ := cmd.Flags().GetString("filter-favorite")
			optionalParams.WithFilterFavorite(val)
		}

		if cmd.Flags().Changed("filter-self-service") {
			val, _ := cmd.Flags().GetString("filter-self-service")
			optionalParams.WithFilterSelfService(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		api := datadogV2.NewAppBuilderApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListApps(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-apps")

		fmt.Println(cmdutil.FormatJSON(res, "appDefinitions"))
	},
}

func init() {

	ListAppsCmd.Flags().Int64("limit", 0, "The number of apps to return per page.")

	ListAppsCmd.Flags().Int64("page", 0, "The page number to return.")

	ListAppsCmd.Flags().String("filter-user-name", "", "Filter apps by the app creator. Usually the user's email.")

	ListAppsCmd.Flags().String("filter-user-uuid", "", "Filter apps by the app creator's UUID.")

	ListAppsCmd.Flags().String("filter-name", "", "Filter by app name.")

	ListAppsCmd.Flags().String("filter-query", "", "Filter apps by the app name or the app creator.")

	ListAppsCmd.Flags().String("filter-deployed", "", "Filter apps by whether they are published.")

	ListAppsCmd.Flags().String("filter-tags", "", "Filter apps by tags.")

	ListAppsCmd.Flags().String("filter-favorite", "", "Filter apps by whether you have added them to your favorites.")

	ListAppsCmd.Flags().String("filter-self-service", "", "Filter apps by whether they are enabled for self-service.")

	ListAppsCmd.Flags().String("sort", "", "The fields and direction to sort apps by.")

	Cmd.AddCommand(ListAppsCmd)
}
