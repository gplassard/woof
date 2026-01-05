package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var GetUserFacetInfoCmd = &cobra.Command{
	Use: "get-user-facet-info [payload]",

	Short: "Get user facet info",
	Long: `Get user facet info
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#get-user-facet-info`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FacetInfoResponse
		var err error

		var body datadogV2.FacetInfoRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err = api.GetUserFacetInfo(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to get-user-facet-info")

		cmd.Println(cmdutil.FormatJSON(res, "users_facet_info"))
	},
}

func init() {
	Cmd.AddCommand(GetUserFacetInfoCmd)
}
