package rum_audience_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetUserFacetInfoCmd = &cobra.Command{
	Use:   "get-user-facet-info",
	
	Short: "Get user facet info",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.GetUserFacetInfo(client.NewContext(apiKey, appKey, site), datadogV2.FacetInfoRequest{})
		cmdutil.HandleError(err, "failed to get-user-facet-info")

		cmd.Println(cmdutil.FormatJSON(res, "users_facet_info"))
	},
}

func init() {
	Cmd.AddCommand(GetUserFacetInfoCmd)
}
