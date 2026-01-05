package restriction_policies

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteRestrictionPolicyCmd = &cobra.Command{
	Use: "delete-restriction-policy [resource_id]",

	Short: "Delete a restriction policy",
	Long: `Delete a restriction policy
Documentation: https://docs.datadoghq.com/api/latest/restriction-policies/#delete-restriction-policy`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRestrictionPoliciesApi(client.NewAPIClient())
		_, err = api.DeleteRestrictionPolicy(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-restriction-policy")

	},
}

func init() {
	Cmd.AddCommand(DeleteRestrictionPolicyCmd)
}
