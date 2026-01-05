package org_connections

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOrgConnectionsCmd = &cobra.Command{
	Use:     "create-org-connections",
	Aliases: []string{"create"},
	Short:   "Create Org Connection",
	Long: `Create Org Connection
Documentation: https://docs.datadoghq.com/api/latest/org-connections/#create-org-connections`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConnectionResponse
		var err error

		var body datadogV2.OrgConnectionCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		res, _, err = api.CreateOrgConnections(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-org-connections")

		cmd.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {

	CreateOrgConnectionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOrgConnectionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOrgConnectionsCmd)
}
