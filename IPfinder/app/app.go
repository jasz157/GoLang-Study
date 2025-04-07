package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func FindIP(c *cli.Context) {
	host := c.String("host")

	ips, failure := net.LookupIP(host)

	if failure != nil {
		log.Fatal(failure)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func ServerFinder(c *cli.Context) {
	host := c.String("host")

	servers, failure := net.LookupNS(host)

	if failure != nil {
		log.Fatal(failure)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

// Init retorna a cli app pronta para ser executada
func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "IPfinder"
	app.Usage = "Busca IP's e nomes de servidores na web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "IPfinder",
			Usage: "Busca endereços de IP na web",
			Flags: flags,

			Action: FindIP,
		},

		{
			Name:   "ServerFinder",
			Usage:  "Encontra os servers de cada host via web",
			Flags:  flags,
			Action: ServerFinder,
		},
	}

	return app
}
