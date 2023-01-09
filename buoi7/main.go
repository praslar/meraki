package main

import (
	"buoi7/team"
	"fmt"
)

func main() {
	getTeam := team.ReturnAllMemberInTeam("meraki")
	fmt.Print("hello team: ", getTeam)
}
