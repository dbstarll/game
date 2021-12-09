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

//力量料理A
func StrA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Str: 5}),
		model.AddGains(false, &model.Gains{NearDamage: 10}),
	)
}

//力量料理B
func StrB() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Str: 10}),
		model.AddGains(false, &model.Gains{NearDamage: 20}),
	)
}

//敏捷料理A
func AgiA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Agi: 5}),
		model.AddGeneral(&model.General{AttackSpeed: 10}),
	)
}

//敏捷料理B
func AgiB() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Agi: 10}),
		model.AddGeneral(&model.General{AttackSpeed: 20}),
	)
}

//体质料理A
func VitA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Vit: 5}),
	)
}

//体质料理B
func VitB() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Vit: 10}),
	)
}

//智力料理A
func IntA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Int: 5}),
		model.AddGains(true, &model.Gains{Damage: 10}),
	)
}

//智力料理B
func IntB() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Int: 10}),
		model.AddGains(true, &model.Gains{Damage: 20}),
	)
}

//灵巧料理A
func DexA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Dex: 5}),
		model.AddGains(false, &model.Gains{RemoteDamage: 10}),
	)
}

//灵巧料理B
func DexB() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Dex: 10}),
		model.AddGains(false, &model.Gains{RemoteDamage: 20}),
	)
}

//幸运料理A
func LukA() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Luk: 5}),
		model.AddGeneral(&model.General{Critical: 10}),
	)
}

//幸运料理B
func LukB() model.CharacterModifier {
	return model.Merge(
		model.AddQuality(&model.Quality{Luk: 10}),
		model.AddGeneral(&model.General{Critical: 20}),
	)
}

//所有料理B
func AllB() model.CharacterModifier {
	return model.Merge(StrB(), AgiB(), VitB(), IntB(), DexB(), LukB())
}
