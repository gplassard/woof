package org_connections

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOrgConnectionsCmd = &cobra.Command{
	Use:   "list-org-connections",
	Aliases: []string{ "list", },
	Short: "List Org Connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.ListOrgConnections(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-org-connections: %v", err)
		}

		cmdutil.PrintJSON(res, "org_connection")
	},
}

func init() {
	Cmd.AddCommand(ListOrgConnectionsCmd)
}
