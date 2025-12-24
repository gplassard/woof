package org_connections

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOrgConnectionsCmd = &cobra.Command{
	Use:   "listorgconnections",
	Short: "List Org Connections",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.ListOrgConnections(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listorgconnections: %v", err)
		}

		cmdutil.PrintJSON(res, "org_connections")
	},
}

func init() {
	Cmd.AddCommand(ListOrgConnectionsCmd)
}
