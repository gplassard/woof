package domain_allowlist

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var PatchDomainAllowlistCmd = &cobra.Command{
	Use:     "patch-domain-allowlist [payload]",
	Aliases: []string{"patch"},
	Short:   "Sets Domain Allowlist",
	Long: `Sets Domain Allowlist
Documentation: https://docs.datadoghq.com/api/latest/domain-allowlist/#patch-domain-allowlist`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DomainAllowlistResponse
		var err error

		var body datadogV2.DomainAllowlistRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDomainAllowlistApi(client.NewAPIClient())
		res, _, err = api.PatchDomainAllowlist(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to patch-domain-allowlist")

		cmd.Println(cmdutil.FormatJSON(res, "domain_allowlist"))
	},
}

func init() {
	Cmd.AddCommand(PatchDomainAllowlistCmd)
}
