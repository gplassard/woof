package org_connections

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListOrgConnectionsCmd = &cobra.Command{
	Use:     "list-org-connections",
	Aliases: []string{"list"},
	Short:   "List Org Connections",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.ListOrgConnections(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-org-connections")

		cmd.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {
	Cmd.AddCommand(ListOrgConnectionsCmd)
}
