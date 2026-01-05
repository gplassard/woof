package network_device_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateDeviceUserTagsCmd = &cobra.Command{
	Use: "update-device-user-tags [device_id] [payload]",

	Short: "Update the tags for a device",
	Long: `Update the tags for a device
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#update-device-user-tags`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListTagsResponse
		var err error

		var body datadogV2.ListTagsResponse
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err = api.UpdateDeviceUserTags(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-device-user-tags")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDeviceUserTagsCmd)
}
