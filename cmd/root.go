package cmd

import (
	"fmt"
	"os"

	"github.com/kylie-a/waypoint/waypoint"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	conf    *waypoint.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "waypoint",
	Short: "A command line interface for deploying SRE services",
	Long:  `A command line interface for deploying SRE services`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".waypoint.yaml", "config file (default is .waypoint.yaml)")
	conf = waypoint.GetConf(cfgFile)
	InitDB(conf)
}
