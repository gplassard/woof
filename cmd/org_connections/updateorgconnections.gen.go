package org_connections

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOrgConnectionsCmd = &cobra.Command{
	Use:   "updateorgconnections [connection_id]",
	Short: "Update Org Connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.UpdateOrgConnections(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), datadogV2.OrgConnectionUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateorgconnections: %v", err)
		}

		cmdutil.PrintJSON(res, "org_connections")
	},
}

func init() {
	Cmd.AddCommand(UpdateOrgConnectionsCmd)
}
