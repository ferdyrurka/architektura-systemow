package action

import (
	"ferdyrurka/architektura-systemow/service"
	"ferdyrurka/architektura-systemow/storage"
	"ferdyrurka/architektura-systemow/util"
	"fmt"
	"github.com/spf13/cobra"
)

func InitXchg() *cobra.Command {
	movXchg := &cobra.Command{
		Use:   "xchg",
		Short: "xchg command",
		Run: xchg(),
	}

	movXchg.Flags().StringP(aFlagName, "a", "", "a records")
	movXchg.Flags().StringP(bFlagName, "b", "", "b records")

	return movXchg
}

func xchg() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		a, b := readArgs(cmd)

		fmt.Println(util.CYAN, "############### FROM ###############", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
		fmt.Println("")

		storage.Xchg(a, b)

		fmt.Println(util.CYAN, "################ TO ################", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
	}
}
