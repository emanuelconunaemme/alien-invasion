package data

type State struct {
	Cities []*City
	Aliens []*Alien
}

// not sure there is a kotlin bias here on having the state determine itself

// Terminal state => The game is over
func (s *State) IsTerminal() bool {
	// no cities or no aliens
	if len(s.Cities) == 0 || len(s.Aliens) == 0 {
		return true
	}

	// only isolated cities, no aliens can move => terminal state
	isolated := true
	for _, city := range s.Cities {
		if city.Aliens == nil || len(city.Aliens) == 0 {
			continue
		}
		// there is an alien that can move, NOT a terminal state
		if city.Neighbors != nil && len(city.Neighbors) > 0 {
			isolated = false
			break
		}
	}
	return isolated
}
