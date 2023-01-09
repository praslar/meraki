package team

func ReturnAllMemberInTeam(teamName string) []string {
	if teamName == "meraki" {
		return []string{"Vu", "Lo", "Thinh"}
	}
	if teamName == "thay" {
		return []string{"Thang"}
	}
	return []string{}
}
