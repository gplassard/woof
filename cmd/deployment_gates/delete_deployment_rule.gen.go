package deployment_gates

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDeploymentRuleCmd = &cobra.Command{
	Use: "delete-deployment-rule [gate_id] [id]",

	Short: "Delete deployment rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		_, err := api.DeleteDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-deployment-rule")

	},
}

func init() {
	Cmd.AddCommand(DeleteDeploymentRuleCmd)
}
