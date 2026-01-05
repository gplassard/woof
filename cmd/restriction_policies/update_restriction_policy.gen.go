package restriction_policies

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateRestrictionPolicyCmd = &cobra.Command{
	Use: "update-restriction-policy [resource_id] [payload]",

	Short: "Update a restriction policy",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RestrictionPolicyUpdateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		res, _, err := api.UpdateRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-restriction-policy")

		cmd.Println(cmdutil.FormatJSON(res, "restriction_policy"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRestrictionPolicyCmd)
}
