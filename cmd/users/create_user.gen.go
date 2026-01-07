package users

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateUserCmd = &cobra.Command{
	Use:     "create-user",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Long: `Create a user
Documentation: https://docs.datadoghq.com/api/latest/users/#create-user`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserResponse
		var err error

		var body datadogV2.UserCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateUser(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-user")

		fmt.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {

	CreateUserCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateUserCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateUserCmd)
}
