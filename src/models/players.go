package models

const (
	StartingPoints = 1200
)

type Player struct {
	FirstName   string
	LastName    string
	Nationality string
	DoB    Date
	Points int
}

func (p *Player) Score(points int) {
	p.Points += points
}