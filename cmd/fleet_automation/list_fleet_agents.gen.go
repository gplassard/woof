package fleet_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetAgentsCmd = &cobra.Command{
	Use: "list-fleet-agents",

	Short: "List all Datadog Agents",
	Long: `List all Datadog Agents
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#list-fleet-agents`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetAgentsResponse
		var err error

		optionalParams := datadogV2.NewListFleetAgentsOptionalParameters()

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("sort-attribute") {
			val, _ := cmd.Flags().GetString("sort-attribute")
			optionalParams.WithSortAttribute(val)
		}

		if cmd.Flags().Changed("sort-descending") {
			val, _ := cmd.Flags().GetString("sort-descending")
			optionalParams.WithSortDescending(val)
		}

		if cmd.Flags().Changed("tags") {
			val, _ := cmd.Flags().GetString("tags")
			optionalParams.WithTags(val)
		}

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListFleetAgents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-fleet-agents")

		fmt.Println(cmdutil.FormatJSON(res, "fleet_automation"))
	},
}

func init() {

	ListFleetAgentsCmd.Flags().Int64("page-number", 0, "Page number for pagination (must be greater than 0).")

	ListFleetAgentsCmd.Flags().Int64("page-size", 0, "Number of results per page (must be greater than 0 and less than or equal to 100).")

	ListFleetAgentsCmd.Flags().String("sort-attribute", "", "Attribute to sort by.")

	ListFleetAgentsCmd.Flags().String("sort-descending", "", "Sort order (true for descending, false for ascending).")

	ListFleetAgentsCmd.Flags().String("tags", "", "Comma-separated list of tags to filter agents.")

	ListFleetAgentsCmd.Flags().String("filter", "", "Filter string for narrowing down agent results.")

	Cmd.AddCommand(ListFleetAgentsCmd)
}
