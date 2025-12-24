package rum_audience_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAccountFacetInfoCmd = &cobra.Command{
	Use:   "getaccountfacetinfo",
	Short: "Get account facet info",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.GetAccountFacetInfo(client.NewContext(apiKey, appKey, site), datadogV2.FacetInfoRequest{})
		if err != nil {
			log.Fatalf("failed to getaccountfacetinfo: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_audience_management")
	},
}

func init() {
	Cmd.AddCommand(GetAccountFacetInfoCmd)
}
