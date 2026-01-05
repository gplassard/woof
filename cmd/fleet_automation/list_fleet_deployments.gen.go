package fleet_automation

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListFleetDeploymentsCmd = &cobra.Command{
	Use: "list-fleet-deployments",

	Short: "List all deployments",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.ListFleetDeployments(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-fleet-deployments")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {
	Cmd.AddCommand(ListFleetDeploymentsCmd)
}
