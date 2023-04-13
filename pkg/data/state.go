package data

import (
	"strings"
)

type State struct {
	Cities []*City
	Aliens []*Alien
}

// not sure there is a kotlin bias here on having the state determine itself

// Terminal state => The game is over
func (s *State) IsTerminal() (bool, string) {
	// no cities or no aliens
	if len(s.Cities) == 0 {
		return true, "No city left"
	}
	if len(s.Aliens) == 0 {
		return true, "No alien left"
	}

	// only isolated cities, no aliens can move => terminal state
	isolated := true
	for _, city := range s.Cities {
		if city.Aliens == nil || len(city.Aliens) == 0 {
			continue
		}
		// there is an alien that can move, NOT a terminal state
		if city.IsConnected() {
			isolated = false
			break
		}
	}
	if isolated {
		return true, "Cities are isolated"
	} else {
		return false, ""
	}
}

func (s *State) KillAlien(alien *Alien) bool {
	for i, a := range s.Aliens {
		if a.Name == alien.Name {
			s.Aliens = append(s.Aliens[:i], s.Aliens[i+1:]...)
			return true
		}
	}
	return false
}

// pardon the kotlin not sure in golang there is a more idiomatic way to do it
func (s *State) ToString() string {
	var sb strings.Builder
	sb.WriteString("=============\n")
	for _, city := range s.Cities {
		sb.WriteString(city.Name + ":")
		if city.Aliens == nil || len(city.Aliens) == 0 {
			sb.WriteString(" empty")
		} else {
			for _, alien := range city.Aliens {
				sb.WriteString(" " + alien.Name)
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("=============\n")
	return sb.String()
}
