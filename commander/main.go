package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/margic/gobazel/commander/cmd"
	flags "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// cmd.Execute()
	flags.StringP("addr", "a", ":8080", "Http listen address")
	viper.BindPFlag("addr", flags.Lookup("addr"))
	flags.Parse()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract command from querystring
		query := r.URL.Query()
		argsString := query.Get("command")
		args := strings.Split(argsString, " ")

		// prep command args and capture output
		buf := &bytes.Buffer{}
		cmd.RootCmd.SetArgs(args)
		cmd.RootCmd.SetOutput(buf)
		cmd.RootCmd.Execute()

		// send results to caller
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write(buf.Bytes())
		if err != nil {
			fmt.Printf("error writing response: %s\n", err)
		}
	})

	http.ListenAndServe(viper.GetString("addr"), h)
}
