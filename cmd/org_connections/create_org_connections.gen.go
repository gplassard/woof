package org_connections

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOrgConnectionsCmd = &cobra.Command{
	Use:   "create_org_connections",
	Short: "Create Org Connection",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.CreateOrgConnections(client.NewContext(apiKey, appKey, site), datadogV2.OrgConnectionCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_org_connections: %v", err)
		}

		cmdutil.PrintJSON(res, "org_connection")
	},
}

func init() {
	Cmd.AddCommand(CreateOrgConnectionsCmd)
}
