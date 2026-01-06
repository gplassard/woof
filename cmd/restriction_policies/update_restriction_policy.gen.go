package restriction_policies

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRestrictionPolicyCmd = &cobra.Command{
	Use: "update-restriction-policy [resource_id]",

	Short: "Update a restriction policy",
	Long: `Update a restriction policy
Documentation: https://docs.datadoghq.com/api/latest/restriction-policies/#update-restriction-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RestrictionPolicyResponse
		var err error

		var body datadogV2.RestrictionPolicyUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-restriction-policy")

		cmd.Println(cmdutil.FormatJSON(res, "restriction_policy"))
	},
}

func init() {

	UpdateRestrictionPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRestrictionPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateRestrictionPolicyCmd)
}
