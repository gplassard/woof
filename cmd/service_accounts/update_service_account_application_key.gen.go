package service_accounts

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:     "update-service-account-application-key [service_account_id] [app_key_id] [payload]",
	Aliases: []string{"update-application-key"},
	Short:   "Edit an application key for this service account",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PartialApplicationKeyResponse
		var err error

		var body datadogV2.ApplicationKeyUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewServiceAccountsApi(client.NewAPIClient())
		res, _, err = api.UpdateServiceAccountApplicationKey(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-service-account-application-key")

		cmd.Println(cmdutil.FormatJSON(res, "application_keys"))
	},
}

func init() {
	Cmd.AddCommand(UpdateServiceAccountApplicationKeyCmd)
}
