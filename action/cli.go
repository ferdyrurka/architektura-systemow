package action

import (
	"ferdyrurka/architektura-systemow/service"
	"ferdyrurka/architektura-systemow/storage"
	"ferdyrurka/architektura-systemow/util"
	"fmt"
	"github.com/spf13/cobra"
)

func InitCli() *cobra.Command {
	cliCmd := &cobra.Command{
		Use:   "cli",
		Short: "cli command",
		Run: cli(),
	}

	return cliCmd
}

func cli() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(util.CYAN, "############### FROM ###############", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
		fmt.Println("")

		storage.Cli()

		fmt.Println(util.CYAN, "################ TO ################", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
	}
}

