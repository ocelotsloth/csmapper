package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/ocelotsloth/csmapper"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "csmapper"
	app.Usage = "Generate map data from callsign list"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "geojson",
			Aliases: []string{"g"},
			Usage:   "Generate a geojson file from callsign list",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "QRZ.com Username",
				},
				cli.StringFlag{
					Name:  "pass, p",
					Usage: "QRZ.com Password",
				},
				cli.StringFlag{
					Name:  "clubs, c",
					Usage: "Clubs CSV File",
				},
			},
			Action: func(c *cli.Context) error {
				clubCsv, err := os.Open(c.String("clubs"))
				if err != nil {
					return err
				}
				defer clubCsv.Close()

				log.Println("parsing file")
				clubs := []csmapper.Club{}
				if err := gocsv.UnmarshalFile(clubCsv, &clubs); err != nil {
					return err
				}

				log.Println("calling generate geojson")
				features := csmapper.GenerateGeoJSON(c.String("user"), c.String("pass"), clubs)
				log.Println("marshalling geojson")
				jsonBytes, err := features.MarshalJSON()
				if err != nil {
					return err
				}
				fmt.Println(string(jsonBytes))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
