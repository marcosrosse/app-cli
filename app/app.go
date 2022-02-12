package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return an app CLI to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application Command Line"
	app.Usage = "Search Ips and Names in the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP address of your site",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "dns",
			Usage:  "NSLookup of your site",
			Flags:  flags,
			Action: searchDns,
		},
		{
			Name:   "email",
			Usage:  "Lookup for an email config",
			Flags:  flags,
			Action: lookupEmail,
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

func searchDns(c *cli.Context) {
	host := c.String("host")

	dns, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, dns := range dns {
		fmt.Println(dns)
	}
}

func lookupEmail(c *cli.Context) {
	host := c.String("host")

	txts, err := net.LookupTXT(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, email := range txts {
		fmt.Println(email)
	}
}
