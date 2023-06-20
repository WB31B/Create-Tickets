package config

type Article struct {
	Id                            uint16
	Username, Information, Ticket string
}

type Page struct {
	Article      []Article
	CurrentPage  int
	TotalPage    int
	CountTickets int
}

var Inforamations = []Article{}
var ShowTicket = Article{}
