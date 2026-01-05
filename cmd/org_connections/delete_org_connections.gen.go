package org_connections

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteOrgConnectionsCmd = &cobra.Command{
	Use:     "delete-org-connections [connection_id]",
	Aliases: []string{"delete"},
	Short:   "Delete Org Connection",
	Long: `Delete Org Connection
Documentation: https://docs.datadoghq.com/api/latest/org-connections/#delete-org-connections`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		_, err = api.DeleteOrgConnections(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-org-connections")

	},
}

func init() {

	Cmd.AddCommand(DeleteOrgConnectionsCmd)
}
