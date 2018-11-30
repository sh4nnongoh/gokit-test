package player

type GameID string

type Game struct {
	GameID  GameID
	Ranking Ranking
}

// NewGame creates a new, unranked game.
