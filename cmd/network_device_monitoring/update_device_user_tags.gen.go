package network_device_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDeviceUserTagsCmd = &cobra.Command{
	Use: "update-device-user-tags [device_id]",

	Short: "Update the tags for a device",
	Long: `Update the tags for a device
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#update-device-user-tags`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListTagsResponse
		var err error

		var body datadogV2.ListTagsResponse
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateDeviceUserTags(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-device-user-tags")

		fmt.Println(cmdutil.FormatJSON(res, "device_user_tag"))
	},
}

func init() {

	UpdateDeviceUserTagsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDeviceUserTagsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDeviceUserTagsCmd)
}
