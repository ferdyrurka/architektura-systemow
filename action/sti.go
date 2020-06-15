package action

import (
	"ferdyrurka/architektura-systemow/service"
	"ferdyrurka/architektura-systemow/storage"
	"ferdyrurka/architektura-systemow/util"
	"fmt"
	"github.com/spf13/cobra"
)

func InitSti() *cobra.Command {
	stiCmd := &cobra.Command{
		Use:   "sti",
		Short: "sti command",
		Run: sti(),
	}

	return stiCmd
}

func sti() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println(util.CYAN, "############### FROM ###############", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
		fmt.Println("")

		storage.Sti()

		fmt.Println(util.CYAN, "################ TO ################", util.NORMAL_TEXT)
		service.Print(storage.FindAll())
	}
}

