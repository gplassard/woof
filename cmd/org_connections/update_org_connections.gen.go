package org_connections

import (
	"fmt"
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
	Long: `Update Org Connection
Documentation: https://docs.datadoghq.com/api/latest/org-connections/#update-org-connections`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConnectionResponse
		var err error

		var body datadogV2.OrgConnectionUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateOrgConnections(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-org-connections")

		fmt.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {

	UpdateOrgConnectionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOrgConnectionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateOrgConnectionsCmd)
}
