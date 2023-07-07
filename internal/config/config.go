package config

type Tickets struct {
	Id                            uint16
	Username, Information, Ticket string
}

type Notes struct {
	Id                 uint16
	Title, Description string
}

type PageTickets struct {
	Tickets      []Tickets
	CurrentPage  int
	TotalPage    int
	CountTickets int
}

type PageNotes struct {
	Notes        []Notes
	CurrentPage  int
	TotalPage    int
	CountTickets int
}

// var Inforamations = []Tickets{}
// var ShowTicket = Tickets{}
