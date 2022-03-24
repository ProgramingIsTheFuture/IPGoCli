package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func getLocation(ip string) {
	req, err := http.Get(fmt.Sprintf("https://api.ip2country.info/ip?%s", ip))
	if err != nil {
		log.Panicln("Error your location")
		return
	}

	var p = make([]byte, 1024)
	n, _ := req.Body.Read(p)
	p = p[:n]
	var location = map[string]string{}
	json.Unmarshal(p, &location)
	fmt.Printf("%s is from: %s  %s\n", ip, location["countryEmoji"], location["countryName"])

}

func main() {

	var (
		getLocationOpt bool
	)

	rootCmd := &cobra.Command{
		Use: "GoCli",
	}

	meCmd := &cobra.Command{
		Use:   "me",
		Short: "Gets your IP address",
		Long:  "Gets your IP address from the out side world",
		Run: func(cmd *cobra.Command, args []string) {
			req, err := http.Get("https://ip-fast.com/api/ip/?format=json")
			if err != nil {
				log.Panicln("Error receiving your IP")
				return
			}

			var p = make([]byte, 1024)
			n, _ := req.Body.Read(p)
			p = p[:n]

			var data = map[string]string{}
			json.Unmarshal(p, &data)

			fmt.Printf("Your IP address is: %s\n", data["ip"])

			if getLocationOpt {
				getLocation(data["ip"])
			}
		},
	}

	locationCmd := &cobra.Command{
		Use:   "location [IP ADDRESS]",
		Short: "Gets the IP address location",
		Long:  "Gets the IP address location",
		Args:  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			getLocation(args[0])
		},
	}

	meCmd.Flags().BoolVarP(&getLocationOpt, "get-location", "l", false, "Decide if you want to receive your location or not!")

	rootCmd.AddCommand(meCmd, locationCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
