package rum_audience_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListConnectionsCmd = &cobra.Command{
	Use: "list-connections [entity]",

	Short: "List connections",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.ListConnections(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-connections")

		cmd.Println(cmdutil.FormatJSON(res, "list_connections_response"))
	},
}

func init() {
	Cmd.AddCommand(ListConnectionsCmd)
}
