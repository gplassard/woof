package service_accounts

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateServiceAccountCmd = &cobra.Command{
	Use:   "create-service-account",
	Aliases: []string{ "create", },
	Short: "Create a service account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err := api.CreateServiceAccount(client.NewContext(apiKey, appKey, site), datadogV2.ServiceAccountCreateRequest{})
		cmdutil.HandleError(err, "failed to create-service-account")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(CreateServiceAccountCmd)
}
