package cmd

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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
	createCmd.Flags().String("health", "/api/health", "health api endpoint used to test for grafana availability")

	// bind to viper
	viper.BindPFlag("grafana.addr", createCmd.Flags().Lookup("addr"))
	viper.BindPFlag("grafana.dashboardPath", createCmd.Flags().Lookup("resource"))
	viper.BindPFlag("dashboards.path", createCmd.Flags().Lookup("path"))
	viper.BindPFlag("grafana.healthPath", createCmd.Flags().Lookup("health"))
}

func create() {
	fmt.Println("create dashboard")
	// read values
	addr := viper.GetString("grafana.addr")
	dashboardPath := viper.GetString("grafana.dashboardPath")
	healthPath := viper.GetString("grafana.healthPath")
	filePath := viper.GetString("dashboards.filePath")

	err := waitForGrafana(addr, healthPath)
	if err != nil {
		fmt.Printf("grafana not responding. is it running at %s, errror: %s", addr, err.Error())
	} else {
		postDashboard(addr, dashboardPath, filePath)
	}

}

func waitForGrafana(addr string, healthPath string) error {
	fmt.Println("waiting for grafana")

	ticker := time.NewTicker(time.Second * 5)
	ticks := 0
	for _ = range ticker.C {
		ticks++
		fmt.Printf("testing connectivity, addr: %s, path: %s\n", addr, healthPath)
		resp, err := http.Get(addr + healthPath)
		if err != nil {
			fmt.Printf("error waiting for grafana: %s\n", err.Error())
		} else {
			fmt.Printf("got a response")
			if resp.StatusCode == http.StatusOK {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					// got a response
					fmt.Printf("response\n status: %s\n body: %s\n", resp.Status, body)
					// check content type
					if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
						ticker.Stop()
						return nil
					}
				}
			}
		}
		if ticks > 20 {
			ticker.Stop()
			return errors.New("grafana did not respond")
		}
	}
	return nil
}

func postDashboard(addr string, dashboardPath string, filePath string) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Printf("Error reading dashboards: %s\n", err.Error())
		return
	}

	client := http.DefaultClient
	for _, file := range files {
		fmt.Printf("Posting %s", file.Name())
		curFile := fmt.Sprintf("%s/%s", filePath, file.Name())
		dash, err := ioutil.ReadFile(curFile)
		if err != nil {
			fmt.Printf("Error reading dashboard file: %s\n", err.Error())
		}
		buf := bytes.NewBufferString(`{
			"dashboard": `)
		buf.Write(dash)
		buf.WriteString(`,
			"overwrite": true
			}`)
		req, err := http.NewRequest("POST", addr+dashboardPath, buf)
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
