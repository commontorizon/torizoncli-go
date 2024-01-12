package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	openapi "github.com/commontorizon/torizon-openapi-go"
	"github.com/urfave/cli/v2"
)

func main() {

	API_CLIENT_SECRET := os.Getenv("TORIZON_API_CLIENT_SECRET")
	API_CLIENT_ID := os.Getenv("TORIZON_API_CLIENT_ID")

	if API_CLIENT_SECRET == "" || API_CLIENT_ID == "" {
		log.Fatalln("TORIZON_API_CLIENT_SECRET or TORIZON_API_CLIENT_ID is not set in the environment variables")
	}

	// boolean flag destinations
	var jsonFlag bool

	app := &cli.App{
		Name:  "torizoncli",
		Usage: "Previsible and composable CLI for the Torizon Cloud",
		Commands: []*cli.Command{
			{
				Name:    "device",
				Aliases: []string{"d"},
				Usage:   "options for task devices",
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Usage:   "list provisioned devices",
						Aliases: []string{"l"},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "json",
								Aliases:     []string{"j"},
								Usage:       "Output in plain json ",
								Destination: &jsonFlag,
							},
						},
						Action: func(cCtx *cli.Context) error {
							c := CreateNewAPIClient(API_CLIENT_ID, API_CLIENT_SECRET)
							resp, httpRes, err := c.DevicesAPI.GetDevices(context.Background()).Execute()

							if err != nil {
								fmt.Errorf("Error", err)
							}

							if httpRes.StatusCode != http.StatusOK {
								fmt.Errorf("received status")
							}

							j, err := resp.MarshalJSON()
							if err != nil {
								fmt.Errorf("Error", err)
							}

							if jsonFlag {
								fmt.Println(string(j))
								return nil
							}

							var result openapi.PaginationResultDeviceInfoBasic
							if err := json.Unmarshal(j, &result); err != nil {
								fmt.Println("Error decoding JSON:", err)
								return nil
							}

							PrintDeviceListInGrid(result.Values)
							return nil
						},
					},
					{
						Name:    "network",
						Usage:   "list network information (local IP address, MAC, hostname) for provisioned devices.",
						Aliases: []string{"n"},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "json",
								Aliases:     []string{"j"},
								Usage:       "Output in plain json ",
								Destination: &jsonFlag,
							},
						},
						Action: func(cCtx *cli.Context) error {
							c := CreateNewAPIClient(API_CLIENT_ID, API_CLIENT_SECRET)
							resp, httpRes, err := c.DevicesAPI.GetDevicesNetwork(context.Background()).Execute()

							if err != nil {
								fmt.Errorf("Error", err)
							}

							if httpRes.StatusCode != http.StatusOK {
								fmt.Errorf("received status")
							}

							j, err := resp.MarshalJSON()
							if err != nil {
								fmt.Errorf("Error", err)
							}

							if jsonFlag {
								fmt.Println(string(j))
								return nil
							}

							var result openapi.PaginationResultNetworkInfo
							if err := json.Unmarshal(j, &result); err != nil {
								fmt.Println("Error decoding JSON:", err)
								return nil
							}

							PrintDeviceNetworkInfoInGrid(result.Values)
							return nil
						},
					},
					{
						Name:    "packages",
						Usage:   "return a list of devices and the packages those devices have installed.",
						Aliases: []string{"pkg"},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "json",
								Aliases:     []string{"j"},
								Usage:       "Output in plain json ",
								Destination: &jsonFlag,
							},
						},
						Action: func(cCtx *cli.Context) error {
							c := CreateNewAPIClient(API_CLIENT_ID, API_CLIENT_SECRET)
							resp, httpRes, err := c.DevicesAPI.GetDevicesPackages(context.Background()).Execute()

							if err != nil {
								fmt.Errorf("Error", err)
							}

							if httpRes.StatusCode != http.StatusOK {
								fmt.Errorf("received status")
							}

							j, err := resp.MarshalJSON()
							if err != nil {
								fmt.Errorf("Error", err)
							}

							if jsonFlag {
								fmt.Println(string(j))
								return nil
							}

							var result openapi.PaginationResultDevicePackages
							if err := json.Unmarshal(j, &result); err != nil {
								fmt.Println("Error decoding JSON:", err)
								return nil
							}

							PrintDevicePackagesInGrid(result.Values)
							return nil
						},
					},
					{
						Name:    "provision",
						Usage:   "provision a device",
						Aliases: []string{"pro"},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task device: ", cCtx.Args().First())
							return nil
						},
						Subcommands: []*cli.Command{
							{
								Name:    "token",
								Usage:   "prints a provisioning token",
								Aliases: []string{"t"},
								Action: func(cCtx *cli.Context) error {
									c := CreateNewAPIClient(API_CLIENT_ID, API_CLIENT_SECRET)
									resp, httpRes, err := c.DevicesAPI.GetDevicesToken(context.Background()).Execute()

									if err != nil {
										fmt.Errorf("Error", err)
									}

									if httpRes.StatusCode != http.StatusOK {
										fmt.Errorf("received status")
									}

									fmt.Println(resp.Token)
									return nil
								},
							},
						},
					},
					{
						Name:    "uptane",
						Usage:   "provide uptane information about a device",
						Aliases: []string{"u"},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:        "json",
								Aliases:     []string{"j"},
								Usage:       "Output in plain json ",
								Destination: &jsonFlag,
							},
						},
						Subcommands: []*cli.Command{
							{
								Name:    "assignment",
								Usage:   "show detailed information about the currently-assigned update for a single device",
								Aliases: []string{"a"},
								Action: func(cCtx *cli.Context) error {
									c := CreateNewAPIClient(API_CLIENT_ID, API_CLIENT_SECRET)
									deviceUuid := cCtx.Args().Get(0)
									resp, httpRes, err := c.DevicesAPI.GetDevicesUptaneDeviceuuidAssignment(context.Background(), deviceUuid).Execute()

									if err != nil {
										fmt.Errorf("Error", err)
									}

									if httpRes.StatusCode != http.StatusOK {
										fmt.Errorf("received status")
									}

									fmt.Println(resp)
									return nil
								},
							},
							{
								Name:    "components",
								Usage:   "show a list of the device uptane components (ECUs).",
								Aliases: []string{"c"},
								Action: func(cCtx *cli.Context) error {
									c := CreateNewAPIClient(API_CLIENT_ID, API_CLIENT_SECRET)
									deviceUuid := cCtx.Args().Get(0)
									_, httpRes, err := c.DevicesAPI.GetDevicesUptaneDeviceuuidComponents(context.Background(), deviceUuid).Execute()

									if err != nil {
										fmt.Errorf("Error", err)
									}

									if httpRes.StatusCode != http.StatusOK {
										fmt.Errorf("received status")
									}

									//fmt.Println(resp.)
									return nil
								},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
