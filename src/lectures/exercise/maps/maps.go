//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `deathstar` to `Retired`
//  - change server status of `star_destroyer` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(serverStatus map[string]int) {
	//  - number of servers
	fmt.Println("Total servers are:", len(serverStatus))

	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
	stats := make(map[int]int)

	for _, status := range serverStatus {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("Uknown server status")
		}
	}

	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are under maintenance")
	fmt.Println(stats[Retired], "servers are retired")
}

func main() {
	servers := []string{"meliniuim_falcon", "deathstar", "star_destroyer", "AT-AT", "x-wing"}

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["deathstar"] = Offline
	serverStatus["star_destroyer"] = Retired

	fmt.Println()
	printServerStatus(serverStatus)

	for _, server := range servers {
		serverStatus[server] = Maintenance
	}

	fmt.Println()
	printServerStatus(serverStatus)
}
