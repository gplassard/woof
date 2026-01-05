package containers

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListContainersCmd = &cobra.Command{
	Use:     "list-containers",
	Aliases: []string{"list"},
	Short:   "Get All Containers",
	Long: `Get All Containers
Documentation: https://docs.datadoghq.com/api/latest/containers/#list-containers`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ContainersResponse
		var err error

		api := datadogV2.NewContainersApi(client.NewAPIClient())
		res, _, err = api.ListContainers(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-containers")

		cmd.Println(cmdutil.FormatJSON(res, "containers"))
	},
}

func init() {
	Cmd.AddCommand(ListContainersCmd)
}
