import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tg",
	Short: "tg is copy template to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tg root")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
