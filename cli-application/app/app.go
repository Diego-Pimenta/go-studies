package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns a CLI Application ready for usage
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search for IPs and domains in the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",           // host is our specified flag
			Value: "devbook.com.br", // default value if none input
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs in ter internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "dns",
			Usage:  "Search for domains in ter internet",
			Flags:  flags,
			Action: searchDomains,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchDomains(c *cli.Context) {
	host := c.String("host")
	domains, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, domain := range domains {
		fmt.Println(domain.Host)
	}
}
