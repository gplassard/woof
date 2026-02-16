package restriction_policies

import (
	"fmt"
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

		optionalParams := datadogV2.NewUpdateRestrictionPolicyOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		if cmd.Flags().Changed("allow-self-lockout") {
			val, _ := cmd.Flags().GetString("allow-self-lockout")
			optionalParams.WithAllowSelfLockout(val)
		}

		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to update-restriction-policy")

		fmt.Println(cmdutil.FormatJSON(res, "restriction_policy"))
	},
}

func init() {

	UpdateRestrictionPolicyCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRestrictionPolicyCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	UpdateRestrictionPolicyCmd.Flags().String("allow-self-lockout", "", "Allows admins (users with the 'user_access_manage' permission) to remove their own access from the resource if set to 'true'. By default, this is set to 'false', preventing admins from locking themselves out.")

	Cmd.AddCommand(UpdateRestrictionPolicyCmd)
}
