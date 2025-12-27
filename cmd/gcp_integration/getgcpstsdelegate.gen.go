package gcp_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetGCPSTSDelegateCmd = &cobra.Command{
	Use:   "getgcpstsdelegate",
	Short: "List delegate account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetGCPSTSDelegate(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getgcpstsdelegate: %v", err)
		}

		cmdutil.PrintJSON(res, "gcp_integration")
	},
}

func init() {
	Cmd.AddCommand(GetGCPSTSDelegateCmd)
}
