package main

import "github.com/spf13/cobra"

var cmd = &cobra.Command{
	Use:   "ipam2dhcp",
	Short: "ipam2dhcp translates IPAM collected data to DHCP rules",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func main() {
	cmd.Execute()
}
