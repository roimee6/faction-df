package util

type Faction struct {
	Name string
	Home *string

	Money int
	Power int

	Members FactionMembers

	Claims []string
}

type FactionMembers struct {
	Leader   string
	Officers []string
	Members  []string
}
