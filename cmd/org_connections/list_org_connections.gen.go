package org_connections

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListOrgConnectionsCmd = &cobra.Command{
	Use:     "list-org-connections",
	Aliases: []string{"list"},
	Short:   "List Org Connections",
	Long: `List Org Connections
Documentation: https://docs.datadoghq.com/api/latest/org-connections/#list-org-connections`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConnectionListResponse
		var err error

		optionalParams := datadogV2.NewListOrgConnectionsOptionalParameters()

		if cmd.Flags().Changed("sink-org-id") {
			val, _ := cmd.Flags().GetString("sink-org-id")
			optionalParams.WithSinkOrgId(val)
		}

		if cmd.Flags().Changed("source-org-id") {
			val, _ := cmd.Flags().GetString("source-org-id")
			optionalParams.WithSourceOrgId(val)
		}

		if cmd.Flags().Changed("limit") {
			val, _ := cmd.Flags().GetInt64("limit")
			optionalParams.WithLimit(val)
		}

		if cmd.Flags().Changed("offset") {
			val, _ := cmd.Flags().GetInt64("offset")
			optionalParams.WithOffset(val)
		}

		api := datadogV2.NewOrgConnectionsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListOrgConnections(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-org-connections")

		fmt.Println(cmdutil.FormatJSON(res, "org_connection"))
	},
}

func init() {

	ListOrgConnectionsCmd.Flags().String("sink-org-id", "", "The Org ID of the sink org.")

	ListOrgConnectionsCmd.Flags().String("source-org-id", "", "The Org ID of the source org.")

	ListOrgConnectionsCmd.Flags().Int64("limit", 0, "The limit of number of entries you want to return. Default is 1000.")

	ListOrgConnectionsCmd.Flags().Int64("offset", 0, "The pagination offset which you want to query from. Default is 0.")

	Cmd.AddCommand(ListOrgConnectionsCmd)
}
