$.extend($.ro, {
    buff: {
        getConfig: function (buff) {
            return $.data(buff, 'config');
        },
        setValue: function (buff, value, errorCallback) {
            switch (typeof value) {
                case "string":
                    const num = parseFloat(value);
                    if (isNaN(num)) {
                        errorCallback.call(this, '不是有效数值: ' + value);
                        return false;
                    } else {
                        value = num;
                    }
                case "number":
                    const config = $.ro.buff.getConfig(buff);
                    if (isNaN(value)) {
                        errorCallback.call(this, '不是有效数值: ' + value);
                        return false;
                    } else if (value < 0) {
                        errorCallback.call(this, '数值不能小于零: ' + value);
                        return false;
                    } else if (config && !config.zero && value == 0) {
                        errorCallback.call(this, '数值不能等于零: ' + value);
                        return false;
                    } else if (config && config.max && value > config.max) {
                        errorCallback.call(this, '数值不能超过最大值[' + config.max + ']: ' + value);
                        return false;
                    } else {
                        const v1 = value + '';
                        const v2 = value.toFixed(config && config.accuracy === 0.1 ? 1 : 0);
                        let displayValue = '+' + (v1.length > v2.length ? v1 : v2);
                        if (config && config.unit) {
                            displayValue += config.unit;
                        }

                        const oldValue = $.ro.buff.getValue(buff);
                        const oldDisplayValue = $.ro.buff.getDisplayValue(buff);
                        $.data(buff, 'buff-value', value);
                        $(buff).children('.buff-value').text(displayValue);

                        $(buff).trigger('buff-change', {
                            oldValue: oldValue, value: value,
                            oldDisplayValue: oldDisplayValue, displayValue: displayValue
                        });
                        return true;
                    }
                default:
                    errorCallback.call(this, '不是有效数值[' + typeof value + ']: ' + value);
                    return false;
            }
        },
        getName: function (buff) {
            return $.data(buff, 'buff-name');
        },
        getValue: function (buff) {
            return $.data(buff, 'buff-value');
        },
        getDisplayValue: function (buff) {
            return $(buff).children('.buff-value').text();
        },
        getEffect: function (buff) {
            const value = $.ro.buff.getValue(buff);
            if ('number' === typeof value && value > 0) {
                return $.ro.buff.getName(buff) + $.ro.buff.getDisplayValue(buff);
            }
        }
    }
});

$.fn.extend({
    buff: function (config) {
        if (typeof config === 'object') {
            let finalConfig = $.extend({}, config);
            const buffName = config.alias || config.effect;
            const item = manualEffects[buffName];
            if (item) {
                $.extend(finalConfig, item);
            } else {
                console.warn("unknown effect", config);
            }
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.data(this, 'buff-name', buffName);
            })
        }
        const buff = {that: this};
        return $.extend(buff, {
            encode: function () {
                const caller = this;
                let result = undefined;
                buff.that.each(function () {
                    result = $.ro.buff.getEffect.call(caller, this);
                    return 'undefined' === typeof result;
                })
                return result;
            },
            decode: function (newValue) {
                return buff.value(newValue);
            },
            name: function () {
                return buff.that.children('.buff-name').text();
            },
            value: function (newValue, errorCallback) {
                const caller = this;
                if ('undefined' === typeof newValue) {
                    let result = undefined;
                    buff.that.each(function () {
                        result = $.ro.buff.getValue.call(caller, this);
                        return 'undefined' === typeof result;
                    })
                    return result;
                } else {
                    buff.that.each(function () {
                        $.ro.buff.setValue.call(caller, this, newValue, errorCallback || console.warn);
                    })
                    return buff;
                }
            }
        });
    }
});
