package main

import "fmt"

type NSPEvent struct {
	Event_date, Event_name, PrimaryContact string
	OrgTeam                                OrganizingTeam
}

type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string
}

func main() {

	Team1 := OrganizingTeam{TeamMembers: []string{"yeswanth", "soundarya"}, PrimaryContact: "yeswanth"}
	Team2 := OrganizingTeam{TeamMembers: []string{"Divya", "Harshita", "chaitanya"}, PrimaryContact: "Divya"}

	//Events iniatilization
	Event1 := NSPEvent{"Jan-23", "NSP-sankranti", "foo1", Team1}
	Event2 := NSPEvent{"Apr-16", "NSP-Eid party", "foo2", Team2}
	Event3 := NSPEvent{"Jun-27", "NSP-Team outing", " ", OrganizingTeam{[]string{"shruti", "giri", "sahana"}, "sahana"}}

	printEventDetails(Event1)
	printEventDetails(Event2)
	printEventDetails(Event3)

}

func printEventDetails(e NSPEvent) {

	fmt.Println("Event Details:")
	fmt.Println("Date:", e.Event_date)
	fmt.Println("Name:", e.Event_name)
	fmt.Println("Primary Contact:", e.PrimaryContact)
	fmt.Println("Organizing Team Members:", e.OrgTeam.TeamMembers)
	fmt.Println("Organizing Team Primary Contact:", e.OrgTeam.PrimaryContact)
	fmt.Println()
}
