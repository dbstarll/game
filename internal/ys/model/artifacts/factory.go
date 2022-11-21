package artifacts

import (
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/entry"
	"github.com/dbstarll/game/internal/ys/dimension/artifacts/position"
)

var (
	Factory生之花 = func(star int, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.FlowerOfLife, entry.Hp, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory死之羽 = func(star int, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.PlumeOfDeath, entry.Atk, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory时之沙 = func(star int, primaryEntry entry.Entry, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.SandsOfEon, primaryEntry, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory空之杯 = func(star int, primaryEntry entry.Entry, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.GobletOfEonothem, primaryEntry, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
	Factory理之冠 = func(star int, primaryEntry entry.Entry, secondaryEntries EntriesLooper) (*Artifacts, error) {
		return New(star, position.CircletOfLogos, primaryEntry, base(int(starHpRect[star-1][0])), secondary(secondaryEntries))
	}
)
