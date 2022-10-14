package reactions

import (
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/reactions/classifies"
	"github.com/dbstarll/game/internal/ys/dimension/elementalism/states"
)

type React struct {
	reaction Reaction
	factor   float64
	state    states.State
}

func NewReact(reaction Reaction, factor float64) *React {
	return NewReactWithState(reaction, factor, -1)
}

func NewReactWithState(reaction Reaction, factor float64, state states.State) *React {
	return &React{reaction: reaction, factor: factor, state: state}
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

func (r *React) GetState() states.State {
	return r.state
}
