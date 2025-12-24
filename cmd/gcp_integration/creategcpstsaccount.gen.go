package gcp_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateGCPSTSAccountCmd = &cobra.Command{
	Use:   "creategcpstsaccount",
	Short: "Create a new entry for your service account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateGCPSTSAccount(client.NewContext(apiKey, appKey, site), datadogV2.GCPSTSServiceAccountCreateRequest{})
		if err != nil {
			log.Fatalf("failed to creategcpstsaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "gcp_integration")
	},
}

func init() {
	Cmd.AddCommand(CreateGCPSTSAccountCmd)
}
