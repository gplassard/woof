package case_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateAttributesCmd = &cobra.Command{
	Use:   "update-attributes [case_id]",
	
	Short: "Update case attributes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateAttributes(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdateAttributesRequest{})
		cmdutil.HandleError(err, "failed to update-attributes")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAttributesCmd)
}
