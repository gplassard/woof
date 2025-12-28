package rum_audience_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetMappingCmd = &cobra.Command{
	Use:   "get-mapping [entity]",
	
	Short: "Get mapping",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		res, _, err := api.GetMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-mapping")

		cmdutil.PrintJSON(res, "get_mappings_response")
	},
}

func init() {
	Cmd.AddCommand(GetMappingCmd)
}
