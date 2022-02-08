package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando - CMDLINE"
	app.Usage = "Busca ip's e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "hashtechsolucoes.tech",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ip's de endereços da internet",
			Flags:  flags,
			Action: buscarIp,
		},
		{
			Name:   "servidor",
			Usage:  "Busca o nome do servidor da internet",
			Flags:  flags,
			Action: buscarServidor,
		},
	}

	return app
}

func buscarIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("IP's encontrados para:", host)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Servidores encontrados para:", host)
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
