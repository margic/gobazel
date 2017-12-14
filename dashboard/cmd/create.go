package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new grafana dashboard",
	Long:  `Creates a new grafana dashboard on the grafana server`,
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}

func create() {
	fmt.Println("create dashboard")

	// read path
	path := viper.GetString("dashboards.path")
	fmt.Printf("Path: %s\n", path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading dashboards: %s\n", err.Error())
	}

	fmt.Printf("Found %d dashboard files\n", len(files))
}
