package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}
	readLogFile(os.Args[1])
	// Your fun starts here.
}

func readLogFile(filePath string) {
	// Open the file.
	f, _ := os.Open(filePath)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.

	programs := map[string]*Program{}
	var installedPackages, removedPackages, upgradedPackages int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if cap(line) > 4 {
			if line[3] != "installed" && line[3] != "upgraded" && line[3] != "removed" {
				continue
			}
			if programs[line[4]] == nil {
				programs[line[4]] = &Program{"-", "-", "-", 0}
			}
			if line[3] == "installed" || line[3] == "reinstalled" {
				programs[line[4]].installed = line[0] + " " + line[1]
				installedPackages++
			} else if line[3] == "upgraded" {
				programs[line[4]].updated = line[0] + " " + line[1]
				programs[line[4]].updates++
				upgradedPackages++
			} else if line[3] == "removed" {
				programs[line[4]].removal = line[0] + " " + line[1]
				removedPackages++
			}
		}

	}

	currentIntalled := installedPackages - removedPackages
	fmt.Println("Packman packages resport")
	fmt.Println("------------------------")
	fmt.Println("- Installed packages : ", installedPackages)
	fmt.Println("- Removed packages : ", removedPackages)
	fmt.Println("- Updated packages : ", upgradedPackages)
	fmt.Println("- Current installed : ", currentIntalled)
	fmt.Println("")
	fmt.Println("List of Packages")
	fmt.Println("------------------------")
	for key, currentProgram := range programs {
		fmt.Println("- Package Name        : ", key)
		fmt.Println("  - Install date      : ", currentProgram.installed)
		fmt.Println("  - Last update date  : ", currentProgram.updated)
		fmt.Println("  - How many updates  : ", currentProgram.updates)
		fmt.Println("  - Removal date      : ", currentProgram.removal)
		fmt.Println("")
	}

}

type Program struct {
	installed, updated, removal string
	updates                     int
}
