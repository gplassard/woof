package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use: "update-current-user-application-key [app_key_id]",

	Short: "Edit an application key owned by current user",
	Long: `Edit an application key owned by current user
Documentation: https://docs.datadoghq.com/api/latest/key-management/#update-current-user-application-key`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationKeyResponse
		var err error

		var body datadogV2.ApplicationKeyUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-current-user-application-key")

		fmt.Println(cmdutil.FormatJSON(res, "current_user_application_key"))
	},
}

func init() {

	UpdateCurrentUserApplicationKeyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCurrentUserApplicationKeyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCurrentUserApplicationKeyCmd)
}
