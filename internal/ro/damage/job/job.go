package job

//职业
type Job int

const (
	Unlimited   Job = iota
	Novice          // 初心者
	Swordman        // 剑士
	Knight1         // 骑士
	Knight2         // 骑士领主
	Knight3         // 符文骑士
	Knight4         // 魔剑士
	Crusader1       // 十字军
	Crusader2       // 圣殿十字军
	Crusader3       // 皇家卫士
	Crusader4       // 世界圣盾
	Magician        // 魔法师
	Wizard1         // 巫师
	Wizard2         // 超魔导士
	Wizard3         // 大法师
	Wizard4         // 无限法师
	Sage1           // 贤者
	Sage2           // 智者
	Sage3           // 元素使
	Sage4           // 时空领主
	Thief           // 盗贼
	Assassin1       // 刺客
	Assassin2       // 十字刺客
	Assassin3       // 十字切割者
	Assassin4       // 灵魂之刃
	Rogue1          // 流氓
	Rogue2          // 神行太保
	Rogue3          // 逐影
	Rogue4          // 影舞
	Archer          // 弓箭手
	Hunter1         // 猎人
	Hunter2         // 神射手
	Hunter3         // 游侠
	Hunter4         // 群星猎手
	Bard1           // 诗人
	Bard2           // 搞笑艺人
	Bard3           // 宫廷乐师
	Bard4           // 日之颂者
	Dancer1         // 舞娘
	Dancer2         // 冷艳舞姬
	Dancer3         // 漫游舞者
	Dancer4         // 月之舞灵
	Acolyte         // 服事
	Priest1         // 牧师
	Priest2         // 神官
	Priest3         // 大主教
	Priest4         // 神使
	Monk1           // 武僧
	Monk2           // 武术宗师
	Monk3           // 修罗
	Monk4           // 龙神
	Merchant        // 商人
	Blacksmith1     // 铁匠
	Blacksmith2     // 神工匠
	Blacksmith3     // 机匠
	Blacksmith4     // 光子大师
	Alchemist1      // 炼金术士
	Alchemist2      // 创造者
	Alchemist3      // 基因学者
	Alchemist4      // 生命缔造者
	Novice1         // 进阶初心者
	Novice3         // 超级初心者
	Novice4         // 初心守护者
	Ninja1          // 忍者
	Ninja3          // 影狼/胧
	Ninja4          // 八岐/天照
	Shooter1        // 神枪手
	Shooter3        // 反叛者
	Shooter4        // 暴君
	NoviceCat       // 初心喵
	Warlock         // 术士
	Summoner1       // 灵术师
	Summoner2       // 召唤师
	Summoner3       // 唤灵者
	Summoner4       // 契约灵
)

func (j Job) String() string {
	switch j {
	case Unlimited:
		return "不限"
	case Novice:
		return "初心者"
	case Swordman:
		return "剑士"
	case Knight1:
		return "骑士"
	case Knight2:
		return "骑士领主"
	case Knight3:
		return "符文骑士"
	case Knight4:
		return "魔剑士"
	case Crusader1:
		return "十字军"
	case Crusader2:
		return "圣殿十字军"
	case Crusader3:
		return "皇家卫士"
	case Crusader4:
		return "世界圣盾"
	case Magician:
		return "魔法师"
	case Wizard1:
		return "巫师"
	case Wizard2:
		return "超魔导士"
	case Wizard3:
		return "大法师"
	case Wizard4:
		return "无限法师"
	case Sage1:
		return "贤者"
	case Sage2:
		return "智者"
	case Sage3:
		return "元素使"
	case Sage4:
		return "时空领主"
	case Thief:
		return "盗贼"
	case Assassin1:
		return "刺客"
	case Assassin2:
		return "十字刺客"
	case Assassin3:
		return "十字切割者"
	case Assassin4:
		return "灵魂之刃"
	case Rogue1:
		return "流氓"
	case Rogue2:
		return "神行太保"
	case Rogue3:
		return "逐影"
	case Rogue4:
		return "影舞"
	case Archer:
		return "弓箭手"
	case Hunter1:
		return "猎人"
	case Hunter2:
		return "神射手"
	case Hunter3:
		return "游侠"
	case Hunter4:
		return "群星猎手"
	case Bard1:
		return "诗人"
	case Bard2:
		return "搞笑艺人"
	case Bard3:
		return "宫廷乐师"
	case Bard4:
		return "日之颂者"
	case Dancer1:
		return "舞娘"
	case Dancer2:
		return "冷艳舞姬"
	case Dancer3:
		return "漫游舞者"
	case Dancer4:
		return "月之舞灵"
	case Acolyte:
		return "服事"
	case Priest1:
		return "牧师"
	case Priest2:
		return "神官"
	case Priest3:
		return "大主教"
	case Priest4:
		return "神使"
	case Monk1:
		return "武僧"
	case Monk2:
		return "武术宗师"
	case Monk3:
		return "修罗"
	case Monk4:
		return "龙神"
	case Merchant:
		return "商人"
	case Blacksmith1:
		return "铁匠"
	case Blacksmith2:
		return "神工匠"
	case Blacksmith3:
		return "机匠"
	case Blacksmith4:
		return "光子大师"
	case Alchemist1:
		return "炼金术士"
	case Alchemist2:
		return "创造者"
	case Alchemist3:
		return "基因学者"
	case Alchemist4:
		return "生命缔造者"
	case Novice1:
		return "进阶初心者"
	case Novice3:
		return "超级初心者"
	case Novice4:
		return "初心守护者"
	case Ninja1:
		return "忍者"
	case Ninja3:
		return "影狼/胧"
	case Ninja4:
		return "八岐/天照"
	case Shooter1:
		return "神枪手"
	case Shooter3:
		return "反叛者"
	case Shooter4:
		return "暴君"
	case NoviceCat:
		return "初心喵"
	case Warlock:
		return "术士"
	case Summoner1:
		return "灵术师"
	case Summoner2:
		return "召唤师"
	case Summoner3:
		return "唤灵者"
	case Summoner4:
		return "契约灵"
	default:
		return "未知"
	}
}

func (j Job) BaseLvAtkRate() int {
	switch j {
	case Thief:
		return 2
	case Assassin1:
		return 3
	case Assassin2, Assassin3, Assassin4:
		return 4
	default:
		return 0
	}
}
