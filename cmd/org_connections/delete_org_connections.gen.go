package org_connections

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var DeleteOrgConnectionsCmd = &cobra.Command{
	Use:   "delete-org-connections [connection_id]",
	Aliases: []string{ "delete", },
	Short: "Delete Org Connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		_, err := api.DeleteOrgConnections(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to delete-org-connections: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteOrgConnectionsCmd)
}
