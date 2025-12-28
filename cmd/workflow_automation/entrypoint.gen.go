package workflow_automation

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "workflow-automation",
	Short: "workflow_automation endpoints",
	Aliases: []string{
		"workflow",
		"wf",
	},
}
