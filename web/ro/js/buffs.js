$.extend($.ro, {
    buffs: {
        getConfig: function (buffs) {
            return $.data(buffs, 'config');
        },
        init: function (buffs, items) {
            $(buffs).children('.buff').remove();
            if (typeof items === 'object' && Array.isArray(items)) {
                items.forEach(function (item) {
                    $.ro.buffs.addBuff(buffs, item);
                });
            }
        },
        getBuff: function (buffs, item) {
            const buffName = item.alias || item.effect;
            let buff = undefined;
            $(buffs).children('.buff').each(function () {
                if ($.ro.buff.getName(this) === buffName) {
                    buff = $(this);
                    return false;
                }
            })
            return buff;
        },
        addBuff: function (buffs, item) {
            const config = $.ro.buffs.getConfig(buffs);
            const buffConfig = $.extend({}, config && {
                click: config.click,
                change: config.change
            }, config && config.buff, item);
            const buff = $.html.div().addClass('buff').buff(buffConfig).that;
            if (item.name) {
                buff.append($.html.div().addClass('buff-title').text(item.name));
            }
            $(buffs).append(buff
                .append($.html.div().addClass('buff-name').text(item.effect))
                .append($.html.div().addClass('buff-value'))
            );
            return buff;
        }
    }
});

$.fn.extend({
    buffs: function (config, items) {
        if (typeof config === 'object') {
            let finalConfig = $.extend({}, config);
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.ro.buffs.init(this, items);
            })
        }
        const buffs = {that: this};
        return $.extend(buffs, {
            encode: function () {
                let effects = [];
                buffs.that.children('.buff').each(function () {
                    const effect = $(this).buff().encode();
                    if ('string' === typeof effect) {
                        effects.push(effect);
                    }
                });
                return effects.length == 0 ? undefined : effects;
            },
            decode: function (items) {
                buffs.reset();
                if (Array.isArray(items)) {
                    items.forEach(function (key) {
                        const idx = key.indexOf('+');
                        if (idx > 0) {
                            const name = key.substr(0, idx);
                            const val = key.substring(idx + 1);
                            buffs.addBuff({effect: name}, val);
                        } else {
                            console.warn('invalid buff', key);
                        }
                    });
                }
                return buffs;
            },
            reset: function () {
                buffs.that.each(function () {
                    const config = $.ro.buffs.getConfig(this);
                    if (config && config.dynamic) {
                        $(this).children('.buff').remove();
                    } else {
                        $(this).children('.buff').buff().reset();
                    }
                });
                return buffs;
            },
            addBuff: function (item, value, errorCallback) {
                const caller = this;
                buffs.that.each(function () {
                    let create = false;
                    let buff = $.ro.buffs.getBuff(this, item);
                    if ('undefined' === typeof buff) {
                        create = true;
                        buff = $.ro.buffs.addBuff(this, item);
                    }
                    buff.buff().value(value, function (error) {
                        if (create) {
                            this.that.remove();
                        }
                        (errorCallback || console.warn).call(caller, error);
                    });
                })
                return buffs;
            }
        });
    }
});
