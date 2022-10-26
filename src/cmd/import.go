package cmd

import (
	"github.com/spf13/cobra"
	"showcase/app"
)

func importSL(app *app.Server) *cobra.Command {
	return &cobra.Command{
		Use:   "import-sl",
		Short: "Import data from sima-land.ru",
		Run: func(cmd *cobra.Command, args []string) {
			app.Init()
			defer app.Close()

			app.Import()
		},
	}
}
