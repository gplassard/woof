package org_connections

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateOrgConnectionsCmd = &cobra.Command{
	Use:     "update-org-connections [connection_id]",
	Aliases: []string{"update"},
	Short:   "Update Org Connection",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.UpdateOrgConnections(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), datadogV2.OrgConnectionUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-org-connections")

		cmd.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {
	Cmd.AddCommand(UpdateOrgConnectionsCmd)
}
