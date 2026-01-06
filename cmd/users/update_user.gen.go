package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateUserCmd = &cobra.Command{
	Use:     "update-user [user_id]",
	Aliases: []string{"update"},
	Short:   "Update a user",
	Long: `Update a user
Documentation: https://docs.datadoghq.com/api/latest/users/#update-user`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserResponse
		var err error

		var body datadogV2.UserUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateUser(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-user")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {

	UpdateUserCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateUserCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateUserCmd)
}
