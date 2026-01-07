package rum_audience_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetUserFacetInfoCmd = &cobra.Command{
	Use: "get-user-facet-info",

	Short: "Get user facet info",
	Long: `Get user facet info
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#get-user-facet-info`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FacetInfoResponse
		var err error

		var body datadogV2.FacetInfoRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetUserFacetInfo(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to get-user-facet-info")

		fmt.Println(cmdutil.FormatJSON(res, "users_facet_info"))
	},
}

func init() {

	GetUserFacetInfoCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	GetUserFacetInfoCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(GetUserFacetInfoCmd)
}
