package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli"
)

func scanPort(protocol, hostname string, port string) bool {
	address := hostname + ":" + port
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()

	return true

}

func main() {
	app := cli.NewApp()
	app.Name = "Network CLI"
	app.Usage = "Check IPs, CNAMES, MX records, Name Servers, and scan your systems ports"

	hostFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Usage: "Hostname to search",
		},
	}

	portFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "port",
			Usage: "Port to check",
		},
		&cli.StringFlag{
			Name:  "protocol",
			Usage: "Protocol for the port, TCP or UDP",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Queries the name servers for a given Host",
			Flags: hostFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a given Host",
			Flags: hostFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a given Host",
			Flags: hostFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a given Host",
			Flags: hostFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
		{
			Name:  "pscan",
			Usage: "Scan a given port to see if its open",
			Flags: portFlags,
			Action: func(c *cli.Context) error {
				open := scanPort(c.String("protocol"), "localhost", c.String("port"))
				fmt.Println("Port "+c.String("port")+":", open)
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
