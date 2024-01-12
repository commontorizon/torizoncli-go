package main

import (
	"fmt"
	"strings"
	"time"

	openapi "github.com/commontorizon/torizon-openapi-go"
)

func PrintDeviceListInGrid(values []openapi.DeviceInfoBasic) {
	fmt.Printf("%-40s%-30s%-30s%-30s%-30s%-20s%-20s%-15s\n", "DeviceUuid", "DeviceName", "DeviceId", "LastSeen", "CreatedAt", "ActivatedAt", "DeviceStatus", "Notes")
	for _, v := range values {
		fmt.Printf("%-40s%-30s%-30s%-30s%-30s%-20s%-20v%-15v\n",
			v.DeviceUuid, v.DeviceName, v.DeviceId,
			formatTime(v.LastSeen),
			v.CreatedAt.Format("2006-01-02 15:04:05"),
			formatTime(v.ActivatedAt),
			v.DeviceStatus, v.Notes)
	}
}

func PrintDeviceNetworkInfoInGrid(values []openapi.NetworkInfo) {
	fmt.Printf("%-40s%-30s%-30s%-30s\n", "DeviceUuid", "Hostname", "LocalIpV4", "MacAddress")
	for _, v := range values {
		// hostname, localipv4 and macaddress may be nil
		hostname := ""
		localipv4 := ""
		macaddress := ""
		if v.Hostname != nil {
			hostname = *v.Hostname
		}

		if v.LocalIpV4 != nil {
			localipv4 = *v.LocalIpV4
		}

		if v.MacAddress != nil {
			macaddress = *v.MacAddress
		}

		fmt.Printf("%-40s%-30s%-30s%-30s\n",
			v.DeviceUuid, hostname, localipv4, macaddress)
	}
}

type TorizonPackageName struct {
	YoctoCodename string
	Machine       string
	Distribution  string
	Flavour       string
	Build         string
}

func parseTorizonPackageName(input string) *TorizonPackageName {
	words := strings.Split(input, "/")

	// TorizonPackageName always has five fields
	// FIXME: this is most likely a very wrong assumption and we should at least try to show the fields if we can.
	if len(words) < 5 {
		words = nil
		words = append(words, "", "", "", "", "")
	}

	result := &TorizonPackageName{
		YoctoCodename: words[0],
		Machine:       words[1],
		Distribution:  words[2],
		Flavour:       words[3],
		Build:         words[4],
	}

	return result
}

func PrintDevicePackagesInGrid(values []openapi.DevicePackages) {
	fmt.Printf("%-40s%-30s%-40s%-30s%-30s%-30s\n", "DeviceUuid", "Flavour", "Build", "Component", "Version", "Checksum")
	for _, v := range values {
		for _, pkg := range v.InstalledPackages {
			pn := parseTorizonPackageName(pkg.Installed.PackageName)
			fmt.Printf("%-40s%-30s%-40s%-30s%-30s%-30s\n",
				v.DeviceUuid, pn.Flavour, pn.Build, pkg.Component, pkg.Installed.PackageVersion, pkg.Installed.Checksum)
		}
	}
}

func formatTime(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}
