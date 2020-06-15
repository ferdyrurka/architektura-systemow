package action

import (
	"ferdyrurka/architektura-systemow/service"
	"ferdyrurka/architektura-systemow/storage"
	"ferdyrurka/architektura-systemow/util"
	"fmt"
	"github.com/spf13/cobra"
)

func InitMov() *cobra.Command {
	movCmd := &cobra.Command{
		Use:   "mov",
		Short: "Mov command",
		Run: mov(),
	}

	movCmd.Flags().StringP(aFlagName, "a", "", "a records")
	movCmd.Flags().StringP(bFlagName, "b", "", "b records")

	return movCmd
}

func mov() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		a, b := readArgs(cmd)

		fmt.Println(util.CYAN, "############### FROM ###############", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
		fmt.Println("")

		storage.Mov(a, b)

		fmt.Println(util.CYAN, "################ TO ################", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
	}
}

