package org_connections

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateOrgConnectionsCmd = &cobra.Command{
	Use:     "update-org-connections [connection_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update Org Connection",
	Long: `Update Org Connection
Documentation: https://docs.datadoghq.com/api/latest/org-connections/#update-org-connections`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConnectionResponse
		var err error

		var body datadogV2.OrgConnectionUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err = api.UpdateOrgConnections(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-org-connections")

		cmd.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOrgConnectionsCmd)
}
