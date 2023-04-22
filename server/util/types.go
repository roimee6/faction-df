package util

type Faction struct {
	Name        string
	Home        *string
	Description *string

	Members FactionMembers

	Claims []string
}

type FactionMembers struct {
	Leader   string
	Officers []string
	Members  []string
}
