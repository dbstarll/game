package buff

import "github.com/dbstarll/game/internal/ro/model"

//庄园
func Manor() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Str: 10, Agi: 10, Vit: 10, Int: 10, Dex: 10, Luk: 10}),
		model.AddGeneral(&model.General{Critical: 30}),
		model.AddGains(false, &model.Gains{Damage: 60, Ignore: 30, Resist: 40}),
		model.AddGains(true, &model.Gains{Damage: 60, Ignore: 30, Resist: 40}),
	)
}

//灵巧料理A
func DexA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Dex: 5}),
		model.AddGains(false, &model.Gains{RemoteDamage: 10}),
	)
}
