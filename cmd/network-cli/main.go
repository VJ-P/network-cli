package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network CLI"
	app.Usage = "Check IPs, CNAMES, MX records, Name Servers, and scan your systems ports"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "vjpatel.ca",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Queries the name servers for a given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
