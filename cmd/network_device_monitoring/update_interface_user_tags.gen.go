package network_device_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateInterfaceUserTagsCmd = &cobra.Command{
	Use: "update-interface-user-tags [interface_id]",

	Short: "Update the tags for an interface",
	Long: `Update the tags for an interface
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#update-interface-user-tags`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListInterfaceTagsResponse
		var err error

		var body datadogV2.ListInterfaceTagsResponse
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateInterfaceUserTags(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-interface-user-tags")

		fmt.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {

	UpdateInterfaceUserTagsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateInterfaceUserTagsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateInterfaceUserTagsCmd)
}
