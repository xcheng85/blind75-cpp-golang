package main

// implement sort.Interface so we can sort Player

type Player struct {
	First string
	Last string
	Number int
}

// mimic payload received from web server
var PlayerPayload = `
{
	"First": "Giannis"
	"Last": "Antetokounmpo"
}

`


// c++ : using PlayerSlice = std::vector<Player>
type PlayerSlice []Player

// three sort.Interface
func (ps PlayerSlice) Len() int {
	return len(ps)
}

// c++: bool operator<=(const Player& other)
// asending order
func (ps PlayerSlice) Less(i,j int) bool {
	return ps[i].Number < ps[j].Number
}

func (ps PlayerSlice) Swap(i,j int) {
	// c++ swap
	// std::swap()
	ps[i], ps[j] = ps[j], ps[i]
}

var players = PlayerSlice{}

