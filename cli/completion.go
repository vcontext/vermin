package cli

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates completion scripts (bash and zsh)",
	Long:  `Generates completion scripts (bash and zsh)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("bash or zsh expected")
	},
}

var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generates bash completion scripts",
	Long: `To load completion run:
. <(vermin completion bash)

You might consider adding it to your shell config file .bashrc
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return rootCmd.GenBashCompletion(os.Stdout)
	},
}

var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generates zsh completion scripts",
	Long: `To load completion run:
vermin completion zsh > ~/.zsh/completion/_vermin
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(bashCmd)
	completionCmd.AddCommand(zshCmd)
}