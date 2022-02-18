$.extend($.ro, {
    buffs: {
        getConfig: function (buffs) {
            return $.data(buffs, 'config');
        },
        initContainer: function (buffs, items) {
            const container = $(buffs);
            container.children('.buff').remove();
            if (typeof items === 'object' && Array.isArray(items)) {
                const config = $.ro.buffs.getConfig(buffs);
                items.forEach(function (item) {
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
                    container.append(buff
                        .append($.html.div().addClass('buff-name').text(item.effect))
                        .append($.html.div().addClass('buff-value'))
                    );
                });
            }
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
                return buffs;
            },
            decode: function (newValue) {
                return buffs;
            },
            reset: function () {
                return buffs;
            }
        }).reset();
    }
});
