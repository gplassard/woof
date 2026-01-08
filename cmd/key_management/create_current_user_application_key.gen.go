package key_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use: "create-current-user-application-key",

	Short: "Create an application key for current user",
	Long: `Create an application key for current user
Documentation: https://docs.datadoghq.com/api/latest/key-management/#create-current-user-application-key`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationKeyResponse
		var err error

		var body datadogV2.ApplicationKeyCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-current-user-application-key")

		fmt.Println(cmdutil.FormatJSON(res, "current_user_application_key"))
	},
}

func init() {

	CreateCurrentUserApplicationKeyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCurrentUserApplicationKeyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCurrentUserApplicationKeyCmd)
}
