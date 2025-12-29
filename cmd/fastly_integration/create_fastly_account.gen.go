package fastly_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFastlyAccountCmd = &cobra.Command{
	Use:   "create-fastly-account",
	
	Short: "Add Fastly account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateFastlyAccount(client.NewContext(apiKey, appKey, site), datadogV2.FastlyAccountCreateRequest{})
		cmdutil.HandleError(err, "failed to create-fastly-account")

		cmd.Println(cmdutil.FormatJSON(res, "fastly-accounts"))
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyAccountCmd)
}
