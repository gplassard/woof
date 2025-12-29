package gcp_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateGCPSTSAccountCmd = &cobra.Command{
	Use:   "create-gcp-sts-account",
	
	Short: "Create a new entry for your service account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateGCPSTSAccount(client.NewContext(apiKey, appKey, site), datadogV2.GCPSTSServiceAccountCreateRequest{})
		cmdutil.HandleError(err, "failed to create-gcp-sts-account")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_service_account"))
	},
}

func init() {
	Cmd.AddCommand(CreateGCPSTSAccountCmd)
}
