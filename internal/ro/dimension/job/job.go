package job

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"strconv"
)

//职业
type Job int

const (
	Unlimited Job = 0 // 不限
	Novice    Job = 1 // 初心者

	Swordman  Job = 11 // 剑士
	Knight1   Job = 12 // 骑士
	Knight2   Job = 13 // 骑士领主
	Knight3   Job = 14 // 符文骑士
	Knight4   Job = 15 // 魔剑士
	Crusader1 Job = 72 // 十字军
	Crusader2 Job = 73 // 圣殿十字军
	Crusader3 Job = 74 // 皇家卫士
	Crusader4 Job = 75 // 世界圣盾

	Magician Job = 21 // 魔法师
	Wizard1  Job = 22 // 巫师
	Wizard2  Job = 23 // 超魔导士
	Wizard3  Job = 24 // 大法师
	Wizard4  Job = 25 // 无限法师
	Sage1    Job = 82 // 贤者
	Sage2    Job = 83 // 智者
	Sage3    Job = 84 // 元素使
	Sage4    Job = 85 // 时空领主

	Thief     Job = 31 // 盗贼
	Assassin1 Job = 32 // 刺客
	Assassin2 Job = 33 // 十字刺客
	Assassin3 Job = 34 // 十字切割者
	Assassin4 Job = 35 // 灵魂之刃
	Rogue1    Job = 92 // 流氓
	Rogue2    Job = 93 // 神行太保
	Rogue3    Job = 94 // 逐影
	Rogue4    Job = 95 // 影舞

	Archer  Job = 41  // 弓箭手
	Hunter1 Job = 42  // 猎人
	Hunter2 Job = 43  // 神射手
	Hunter3 Job = 44  // 游侠
	Hunter4 Job = 45  // 群星猎手
	Bard1   Job = 102 // 诗人
	Bard2   Job = 103 // 搞笑艺人
	Bard3   Job = 104 // 宫廷乐师
	Bard4   Job = 105 // 日之颂者
	Dancer1 Job = 112 // 舞娘
	Dancer2 Job = 113 // 冷艳舞姬
	Dancer3 Job = 114 // 漫游舞者
	Dancer4 Job = 115 // 月之舞灵

	Acolyte Job = 51  // 服事
	Priest1 Job = 52  // 牧师
	Priest2 Job = 53  // 神官
	Priest3 Job = 54  // 大主教
	Priest4 Job = 55  // 神使
	Monk1   Job = 122 // 武僧
	Monk2   Job = 123 // 武术宗师
	Monk3   Job = 124 // 修罗
	Monk4   Job = 125 // 龙神

	Merchant    Job = 61  // 商人
	Blacksmith1 Job = 62  // 铁匠
	Blacksmith2 Job = 63  // 神工匠
	Blacksmith3 Job = 64  // 机匠
	Blacksmith4 Job = 65  // 光子大师
	Alchemist1  Job = 132 // 炼金术士
	Alchemist2  Job = 133 // 创造者
	Alchemist3  Job = 134 // 基因学者
	Alchemist4  Job = 135 // 生命缔造者

	Novice1 Job = 143 // 进阶初心者
	Novice3 Job = 144 // 超级初心者
	Novice4 Job = 145 // 初心守护者

	NoviceCat Job = 150 // 初心喵
	Warlock   Job = 151 // 术士
	Summoner1 Job = 152 // 灵术师
	Summoner2 Job = 153 // 召唤师
	Summoner3 Job = 154 // 唤灵者
	Summoner4 Job = 155 // 契约灵

	Ninja1 Job = 163 // 忍者
	Ninja3 Job = 164 // 影狼/胧
	Ninja4 Job = 165 // 八岐/天照

	Shooter1 Job = 173 // 神枪手
	Shooter3 Job = 174 // 反叛者
	Shooter4 Job = 175 // 暴君

	Job183 Job = 183 // 悟灵士
	Job184 Job = 184 // 猎灵士
	Job185 Job = 185 // 双生魔灵

	Job204 Job = 204 // 火焰神
	Job205 Job = 205 // 辉煌帝
	Job213 Job = 213 // 小救星
	Job214 Job = 214 // 龙神丸
	Job215 Job = 215 // 龙王丸
	Job223 Job = 223 // 魔语者
	Job224 Job = 224 // 混沌魔导士
	Job225 Job = 225 // 秀逗魔导士
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
	case Job183:
		return "悟灵士"
	case Job184:
		return "猎灵士"
	case Job185:
		return "双生魔灵"
	case Job204:
		return "火焰神"
	case Job205:
		return "辉煌帝"
	case Job213:
		return "小救星"
	case Job214:
		return "龙神丸"
	case Job215:
		return "龙王丸"
	case Job223:
		return "魔语者"
	case Job224:
		return "混沌魔导士"
	case Job225:
		return "秀逗魔导士"
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

func (j *Job) UnmarshalYAML(value *yaml.Node) error {
	for i := Unlimited; i <= Job225; i++ {
		if i.String() == value.Value {
			*j = i
			break
		}
	}
	return nil
}

func (j *Job) UnmarshalJSON(data []byte) error {
	if str, err := strconv.Unquote(string(data)); err != nil {
		return errors.WithStack(err)
	} else if i, err := strconv.Atoi(str); err != nil {
		return errors.WithStack(err)
	} else {
		*j = Job(i)
		return nil
	}
}
