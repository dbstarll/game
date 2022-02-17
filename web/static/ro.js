const manualEffects = {
    'MaxHp': {accuracy: 1, unit: ''},
    'MaxSp': {accuracy: 1, unit: ''},
    'Hp恢复': {accuracy: 1, unit: ''},
    'Hp%恢复': {accuracy: 0.1, unit: '%'},
    'Sp恢复': {accuracy: 1, unit: ''},
    'Sp%恢复': {accuracy: 0.1, unit: '%'},
    '物理攻击': {accuracy: 1, unit: ''},
    '物理攻击%': {accuracy: 0.1, unit: '%'},
    '物理防御': {accuracy: 1, unit: ''},
    '物理防御%': {accuracy: 0.1, unit: '%'},
    '物伤减免': {accuracy: 0.1, unit: '%'},
    '精炼物攻': {accuracy: 1, unit: ''},
    '魔法攻击': {accuracy: 1, unit: ''},
    '魔法攻击%': {accuracy: 0.1, unit: '%'},
    '魔法防御': {accuracy: 1, unit: ''},
    '魔法防御%': {accuracy: 0.1, unit: '%'},
    '魔伤减免': {accuracy: 0.1, unit: '%'},
    '精炼魔攻': {accuracy: 1, unit: ''},
    '力量': {accuracy: 1, unit: ''},
    '敏捷': {accuracy: 1, unit: ''},
    '体质': {accuracy: 1, unit: ''},
    '智力': {accuracy: 1, unit: ''},
    '灵巧': {accuracy: 1, unit: ''},
    '幸运': {accuracy: 1, unit: ''},
    '对无属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对地属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对风属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对水属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对火属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对圣属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对暗属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对念属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对不死属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对毒属性魔物增伤%': {accuracy: 0.1, unit: '%'},
    '对无属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对地属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对风属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对水属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对火属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对圣属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对暗属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对念属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对不死属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '对毒属性伤害减免%': {accuracy: 0.1, unit: '%'},
    '无属性攻击': {accuracy: 0.1, unit: '%'},
    '地属性攻击': {accuracy: 0.1, unit: '%'},
    '风属性攻击': {accuracy: 0.1, unit: '%'},
    '水属性攻击': {accuracy: 0.1, unit: '%'},
    '火属性攻击': {accuracy: 0.1, unit: '%'},
    '圣属性攻击': {accuracy: 0.1, unit: '%'},
    '暗属性攻击': {accuracy: 0.1, unit: '%'},
    '念属性攻击': {accuracy: 0.1, unit: '%'},
    '不死属性攻击': {accuracy: 0.1, unit: '%'},
    '毒属性攻击': {accuracy: 0.1, unit: '%'},
    '对无形减伤%': {accuracy: 0.1, unit: '%'},
    '对人形减伤%': {accuracy: 0.1, unit: '%'},
    '对植物减伤%': {accuracy: 0.1, unit: '%'},
    '对动物减伤%': {accuracy: 0.1, unit: '%'},
    '对昆虫减伤%': {accuracy: 0.1, unit: '%'},
    '对鱼贝减伤%': {accuracy: 0.1, unit: '%'},
    '对天使减伤%': {accuracy: 0.1, unit: '%'},
    '对恶魔减伤%': {accuracy: 0.1, unit: '%'},
    '对不死减伤%': {accuracy: 0.1, unit: '%'},
    '对龙族减伤%': {accuracy: 0.1, unit: '%'},
    '对无形增伤%': {accuracy: 0.1, unit: '%'},
    '对人形增伤%': {accuracy: 0.1, unit: '%'},
    '对植物增伤%': {accuracy: 0.1, unit: '%'},
    '对动物增伤%': {accuracy: 0.1, unit: '%'},
    '对昆虫增伤%': {accuracy: 0.1, unit: '%'},
    '对鱼贝增伤%': {accuracy: 0.1, unit: '%'},
    '对天使增伤%': {accuracy: 0.1, unit: '%'},
    '对恶魔增伤%': {accuracy: 0.1, unit: '%'},
    '对不死增伤%': {accuracy: 0.1, unit: '%'},
    '对龙族增伤%': {accuracy: 0.1, unit: '%'},
    '对小体型增伤%': {accuracy: 0.1, unit: '%'},
    '对中体型增伤%': {accuracy: 0.1, unit: '%'},
    '对大体型增伤%': {accuracy: 0.1, unit: '%'},
    '对小体型减伤%': {accuracy: 0.1, unit: '%'},
    '对中体型减伤%': {accuracy: 0.1, unit: '%'},
    '对大体型减伤%': {accuracy: 0.1, unit: '%'},
    '对玩家增伤%': {accuracy: 0.1, unit: '%'},
    '对MVP/Mini增伤%': {accuracy: 0.1, unit: '%'},
    '中毒抵抗': {accuracy: 1, unit: ''},
    '流血抵抗': {accuracy: 1, unit: ''},
    '灼烧抵抗': {accuracy: 1, unit: ''},
    '眩晕抵抗': {accuracy: 1, unit: ''},
    '冰冻抵抗': {accuracy: 1, unit: ''},
    '石化抵抗': {accuracy: 1, unit: ''},
    '睡眠抵抗': {accuracy: 1, unit: ''},
    '恐惧抵抗': {accuracy: 1, unit: ''},
    '定身抵抗': {accuracy: 1, unit: ''},
    '沉默抵抗': {accuracy: 1, unit: ''},
    '诅咒抵抗': {accuracy: 1, unit: ''},
    '黑暗抵抗': {accuracy: 1, unit: ''},
    '闪避': {accuracy: 1, unit: ''},
    '命中': {accuracy: 1, unit: ''},
    '暴击': {accuracy: 1, unit: ''},
    '暴伤': {accuracy: 0.1, unit: '%'},
    '装备攻速': {accuracy: 0.1, unit: '%'},
    '移动速度%': {accuracy: 0.1, unit: '%'}
};