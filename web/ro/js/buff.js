$.extend($.ro, {
    buff: {
        init: function (buff, config) {
            const container = $(buff).on('click', function (event) {
                event.preventDefault();
                $(buff).trigger('buff-click', {buff: buff});
                return false;
            });
            if ('function' === typeof config.click) {
                container.on('buff-click', config.click);
            }
            if ('function' === typeof config.change) {
                container.on('buff-change', config.change);
            }
        },
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
                    if (config && config.invalid) {
                        errorCallback.call(this, '未知的属性: ' + $.ro.buff.getName(buff));
                        return false;
                    } else if (isNaN(value)) {
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
                        const oldValue = $.ro.buff.getValue(buff);
                        if (oldValue !== value) {
                            $.data(buff, 'buff-value', value);

                            const v1 = value + '';
                            const v2 = value.toFixed(config && config.accuracy === 0.1 ? 1 : 0);
                            let displayValue = '+' + (v1.length > v2.length ? v1 : v2);
                            if (config && config.unit) {
                                displayValue += config.unit;
                            }
                            const oldDisplayValue = $.ro.buff.getDisplayValue(buff);
                            $(buff).children('.buff-value').text(displayValue);

                            $(buff).trigger('buff-change', {
                                buff: buff,
                                oldValue: oldValue, value: value,
                                oldDisplayValue: oldDisplayValue, displayValue: displayValue
                            });
                            return true;
                        } else {
                            return false;
                        }
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
                $.extend(finalConfig, {invalid: true});
                console.warn("unknown effect", config);
            }
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.data(this, 'buff-name', buffName);
                $.ro.buff.init(this, finalConfig);
            })
        }
        const buff = {that: this};
        return $.extend(buff, {
            export: function () {
                const caller = this;
                let result = undefined;
                buff.that.each(function () {
                    result = $.ro.buff.getEffect.call(caller, this);
                    return 'undefined' === typeof result;
                })
                return result;
            },
            import: function (newValue) {
                return buff.value(newValue);
            },
            reset: function () {
                buff.that.each(function () {
                    const config = $.ro.buff.getConfig(this);
                    if (config && config.zero) {
                        $.ro.buff.setValue(this, 0);
                    }
                })
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
