package restriction_policies

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRestrictionPolicyCmd = &cobra.Command{
	Use: "get-restriction-policy [resource_id]",

	Short: "Get a restriction policy",
	Long: `Get a restriction policy
Documentation: https://docs.datadoghq.com/api/latest/restriction-policies/#get-restriction-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionPolicyResponse
		var err error

		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		res, _, err = api.GetRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-restriction-policy")

		cmd.Println(cmdutil.FormatJSON(res, "restriction_policy"))
	},
}

func init() {
	Cmd.AddCommand(GetRestrictionPolicyCmd)
}
