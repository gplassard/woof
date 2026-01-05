package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateUserCmd = &cobra.Command{
	Use:     "create-user [payload]",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UserCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.CreateUser(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-user")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(CreateUserCmd)
}
