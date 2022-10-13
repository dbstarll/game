package reactions

import "github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"

type React struct {
	reaction Reaction
	factor   float64
}

func NewReact(reaction Reaction, factor float64) *React {
	return &React{reaction: reaction, factor: factor}
}

func (r *React) Match(classify classifies.Classify) bool {
	return classify == r.reaction.Classify()
}

func (r *React) GetReaction() Reaction {
	return r.reaction
}

func (r *React) GetFactor() float64 {
	return r.factor
}
