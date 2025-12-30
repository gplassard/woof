package org_connections

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOrgConnectionsCmd = &cobra.Command{
	Use:     "create-org-connections",
	Aliases: []string{"create"},
	Short:   "Create Org Connection",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.CreateOrgConnections(client.NewContext(apiKey, appKey, site), datadogV2.OrgConnectionCreateRequest{})
		cmdutil.HandleError(err, "failed to create-org-connections")

		cmd.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {
	Cmd.AddCommand(CreateOrgConnectionsCmd)
}
