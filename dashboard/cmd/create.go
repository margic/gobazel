package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

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

	// create flags
	createCmd.Flags().String("addr", "", "Grafana host address e.g. http://host:port")
	createCmd.Flags().String("resource", "/api/dashboards/db", "dashboard resource path")
	createCmd.Flags().String("path", "", "path to local dashboard folder")

	// bind to viper
	viper.BindPFlag("grafana.addr", createCmd.Flags().Lookup("addr"))
	viper.BindPFlag("grafana.resource", createCmd.Flags().Lookup("resource"))
	viper.BindPFlag("dashboards.path", createCmd.Flags().Lookup("path"))
}

func create() {
	fmt.Println("create dashboard")

	// read path
	path := viper.GetString("dashboards.path")
	addr := viper.GetString("grafana.addr")
	resource := viper.GetString("grafana.resource")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading dashboards: %s\n", err.Error())
		return
	}

	client := http.DefaultClient
	for _, file := range files {
		fmt.Printf("Posting %s", file.Name())
		curFile := fmt.Sprintf("%s/%s", path, file.Name())
		dash, err := ioutil.ReadFile(curFile)
		if err != nil {
			fmt.Printf("Error reading dashboard file: %s\n", err.Error())
		}
		req, err := http.NewRequest("POST", addr+resource, bytes.NewReader(dash))
		req.Header.Set("Content-Type", "application/json")
		req = withAuth(req)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error posting dashboard %s:%s\n", file.Name(), err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error status %d returned from grafana", resp.StatusCode)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("File: %s\nResponse: %s\n", file.Name(), body)
	}
}

// withAuth adds the auth header this would normally be a middleware so lets keep it separate for now
func withAuth(req *http.Request) *http.Request {
	user := viper.GetString("grafana.username")
	password := viper.GetString("grafana.password")
	basic := base64.URLEncoding.EncodeToString([]byte(user + ":" + password))
	req.Header.Set("Authorization", "Basic "+basic)
	return req
}
