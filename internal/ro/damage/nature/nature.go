package nature

//属性
type Nature int

const (
	Unlimited Nature = iota
	Neutral          // 无
	Earth            // 地
	Wind             // 风
	Water            // 水
	Fire             // 火
	Holy             // 圣
	Dark             // 暗
	Undead           // 不死
	Ghost            // 念
	Poison           // 毒
)

var restraints = [][]float32{
	//{无,地,风,水,火,圣,暗,不死,念,毒}
	{1, 1, 1, 1, 1, 1, 1, 1, 0.25, 1},                    //无
	{1, 0.25, 2, 1, 0.5, 0.75, 1, 1, 1, 0.75},            //地
	{1, 0.5, 0.25, 2, 1, 0.75, 1, 1, 1, 0.75},            //风
	{1, 1, 0.5, 0.25, 2, 0.75, 1, 1.5, 1, 0.75},          //水
	{1, 2, 1, 0.5, 0.25, 0.75, 1, 2, 1, 0.75},            //火
	{1, 1, 1, 1, 1, 0.25, 2, 2, 1, 1.25},                 //圣
	{1, 1, 1, 1, 1, 2, 0.25, 0.25, 1, 0.25},              //暗
	{1, 0.5, 0.5, 0.5, 0.5, 1.75, 0.25, 0.25, 1, 0.25},   //不死
	{0.25, 1, 1, 1, 1, 0.75, 0.75, 1.75, 2, 1},           //念
	{1, 1.25, 1.25, 1, 1.25, 0.5, 0.25, 0.25, 0.5, 0.25}, //毒
}

func (n Nature) String() string {
	switch n {
	case Unlimited:
		return "不限"
	case Neutral:
		return "无"
	case Earth:
		return "地"
	case Wind:
		return "风"
	case Water:
		return "水"
	case Fire:
		return "火"
	case Holy:
		return "圣"
	case Dark:
		return "暗"
	case Undead:
		return "不死"
	case Ghost:
		return "念"
	case Poison:
		return "毒"
	default:
		return "未知"
	}
}

func (n Nature) Restraint(defence Nature) float32 {
	if n > Unlimited && n <= Poison {
		if defence > Unlimited && defence <= Poison {
			return restraints[n-1][defence-1]
		}
	}
	return 1
}
