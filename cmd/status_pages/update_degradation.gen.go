package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateDegradationCmd = &cobra.Command{
	Use: "update-degradation [page_id] [degradation_id]",

	Short: "Update degradation",
	Long: `Update degradation
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#update-degradation`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Degradation
		var err error

		var body datadogV2.PatchDegradationRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateDegradation(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]), body)
		cmdutil.HandleError(err, "failed to update-degradation")

		fmt.Println(cmdutil.FormatJSON(res, "degradations"))
	},
}

func init() {

	UpdateDegradationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDegradationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDegradationCmd)
}
