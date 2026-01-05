package organizations

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateOrgConfigCmd = &cobra.Command{
	Use: "update-org-config [org_config_name]",

	Short: "Update a specific Org Config",
	Long: `Update a specific Org Config
Documentation: https://docs.datadoghq.com/api/latest/organizations/#update-org-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OrgConfigGetResponse
		var err error

		var body datadogV2.OrgConfigWriteRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		res, _, err = api.UpdateOrgConfig(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-org-config")

		cmd.Println(cmdutil.FormatJSON(res, "org_configs"))
	},
}

func init() {

	UpdateOrgConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOrgConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateOrgConfigCmd)
}
