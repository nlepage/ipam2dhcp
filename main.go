package main

import (
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

	// FIXME start a goroutine to iterate on the ticker and call tryRefreshFile

	return ticker
}

func startHTTPServer() error {
	r := mux.NewRouter()

	// FIXME Add a new HandleFunc to the router matching "/refresh" and calling tryRefreshFile

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

	// FIXME write ips to f using json package from standard library

	return nil
}
