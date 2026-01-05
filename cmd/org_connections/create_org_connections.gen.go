package org_connections

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateOrgConnectionsCmd = &cobra.Command{
	Use:     "create-org-connections [payload]",
	Aliases: []string{"create"},
	Short:   "Create Org Connection",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.OrgConnectionCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err := api.CreateOrgConnections(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-org-connections")

		cmd.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {
	Cmd.AddCommand(CreateOrgConnectionsCmd)
}
