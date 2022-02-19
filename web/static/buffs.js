$.extend($.ro, {
    buffs: {
        getConfig: function (buffs) {
            return $.data(buffs, 'config');
        },
        initContainer: function (buffs, items) {
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
            const buffConfig = $.extend({}, config && config.buff, item);
            const buff = $.html.div().addClass('buff').buff(buffConfig).that;
            if (config && 'function' === typeof config.click) {
                buff.on('click', config.click);
            }
            if (config && 'function' === typeof config.change) {
                buff.on('buff-change', config.change);
            }
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
                $.ro.buffs.initContainer(this, items);
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
            decode: function (newValue) {
                buffs.reset();
                if (Array.isArray(newValue)) {
                    newValue.forEach(function (key) {
                        const idx = key.indexOf('+');
                        const name = key.substr(0, idx);
                        const val = key.substring(idx);
                        buffs.addBuff({effect: name}, val);
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
