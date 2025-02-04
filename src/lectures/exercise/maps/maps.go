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
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
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

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := setAllStatuses(servers, Online)
	fmt.Println("--- Changing status ---")
	showStatus(serverMap)

	serverMap["darkstar"] = Retired
	serverMap["aiur"] = Offline
	fmt.Println("--- Changing status ---")
	showStatus(serverMap)

	serverMap = setAllStatuses(servers, Maintenance)
	fmt.Println("--- Changing status ---")
	showStatus(serverMap)
}

func setAllStatuses(servers []string, status int) map[string]int {
	serverMap := make(map[string]int)
	for i := 0; i < len(servers); i++ {
		current := servers[i]
		serverMap[current] = status
	}
	return serverMap
}

func showStatus(serversMap map[string]int) {
	for server, status := range serversMap {
		fmt.Println("Server:", server, "-- Status:", status)
	}
}
