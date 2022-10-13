package reactions

import "github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"

type Factor struct {
	reaction Reaction
	factor   float64
}

func NewFactor(reaction Reaction, factor float64) *Factor {
	return &Factor{reaction: reaction, factor: factor}
}

func (f *Factor) GetReaction() Reaction {
	return f.reaction
}

func (f *Factor) Match(classify classifies.Classify) bool {
	return classify == f.reaction.Classify()
}

func (f *Factor) GetFactor() float64 {
	return f.factor
}
