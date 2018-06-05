package main

import (
	"encoding/json"
	"fmt"
	"formation-go/ipam2dhcp/ipam"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "ipam2dhcp",
	Short: "ipam2dhcp translates IPAM collected data to DHCP rules",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Initial refresh...")
		tryRefreshFile()

		ticker := startPulling()
		defer ticker.Stop()

		return startHTTPServer()
	},
}

func main() {
	cmd.Execute()
}

func startPulling() *time.Ticker {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for range ticker.C {
			fmt.Println("Pulling refresh...")
			tryRefreshFile()
		}
	}()

	return ticker
}

func startHTTPServer() error {
	r := mux.NewRouter()

	r.HandleFunc("/refresh", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Manual refresh...")
		tryRefreshFile()
	})

	http.Handle("/", r)
	return http.ListenAndServe(":8080", nil)
}

func tryRefreshFile() {
	if err := refreshFile(); err != nil {
		fmt.Println("Error:", err)
	}
}

func refreshFile() error {
	ips, err := ipam.ListIpAddresses()
	if err != nil {
		return err
	}

	f, err := os.Create("ips.json")
	if err != nil {
		return err
	}
	defer f.Close()

	if err = json.NewEncoder(f).Encode(ips); err != nil {
		return err
	}

	return nil
}
