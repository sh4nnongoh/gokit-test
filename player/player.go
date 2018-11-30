package player

// PlayerID uniquely identifies a player
type PlayerID string

type Player struct {
	PlayerID    PlayerID
	GameHistory GameHistory
}

// New creates a new player with no GameHistory
func New(id PlayerID) *Player {
	return nil
}

type GameHistory struct {
	Games []Game
}
