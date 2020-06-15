package action

import (
	"ferdyrurka/architektura-systemow/util"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type CobraFn func(cmd *cobra.Command, args []string)

const aFlagName = "a-record"
const bFlagName = "b-record"

func readArgs(cmd *cobra.Command) (a string, b string) {
	a, _ = cmd.Flags().GetString(aFlagName)
	b, _ = cmd.Flags().GetString(bFlagName)

	if b == "" || a == "" {
		log.Fatal(util.WARNING_TEXT, "a and b is required parameters", util.NORMAL_TEXT)
		os.Exit(1)
	}
	if a == b  {
		log.Fatal(util.WARNING_TEXT, "a and b parameters must is not equal", util.NORMAL_TEXT)
		os.Exit(2)
	}

	return a, b
}
