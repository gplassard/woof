package metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTagConfigurationCmd = &cobra.Command{
	Use: "create-tag-configuration [metric_name]",

	Short: "Create a tag configuration",
	Long: `Create a tag configuration
Documentation: https://docs.datadoghq.com/api/latest/metrics/#create-tag-configuration`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricTagConfigurationResponse
		var err error

		var body datadogV2.MetricTagConfigurationCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-tag-configuration")

		fmt.Println(cmdutil.FormatJSON(res, "tag_configuration"))
	},
}

func init() {

	CreateTagConfigurationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTagConfigurationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTagConfigurationCmd)
}
