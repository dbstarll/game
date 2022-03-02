$.extend($.ro, {
    dialog: {
        init: function (dialog, config) {
            $(dialog).dialog($.extend({
                modal: true,
                autoOpen: false,
                resizable: false
            }, config)).append($.html.p({'role': 'message'}));
        },
        initMessage: function (dialog, config) {
            this.init(dialog, $.extend({
                buttons: {
                    "确定": function () {
                        $(this).dialog("close");
                    }
                }
            }, config));
        },
        initConfirm: function (dialog, config) {
            this.init(dialog, $.extend({
                buttons: {
                    "确定": function () {
                        $(this).attr('confirm', 'true').dialog("close");
                    },
                    "取消": function () {
                        $(this).dialog("close");
                    }
                }
            }, config));
        }
    }
});

$.fn.extend({
    message: function (config) {
        if (typeof config === 'object') {
            const finalConfig = $.extend({}, config);
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.ro.dialog.initMessage(this, finalConfig);
            });
        }
        const message = {that: this};
        return $.extend(message, {
            message: function (title, msg, callback) {
                let myPromise = new Promise((resolve) => {
                    const dialog = message.that
                        .dialog("option", "title", title)
                        .dialog("open")
                        .on("dialogclose", () => {
                            dialog.off("dialogclose");
                            resolve(true);
                            return false;
                        });
                    dialog.children('p[role=message]').text(msg);
                });
                if ('function' === typeof callback) {
                    myPromise.then(callback);
                }
                return myPromise;
            }
        });
    },
    confirm: function (config) {
        if (typeof config === 'object') {
            const finalConfig = $.extend({}, config);
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.ro.dialog.initConfirm(this, finalConfig);
            });
        }
        const confirm = {that: this};
        return $.extend(confirm, {
            confirm: function (title, msg, callback) {
                let myPromise = new Promise((resolve) => {
                    const dialog = confirm.that.removeAttr('confirm')
                        .dialog("option", "title", title)
                        .dialog("open")
                        .on("dialogclose", () => {
                            dialog.off("dialogclose");
                            resolve(dialog.attr('confirm') === 'true');
                            return false;
                        });
                    dialog.children('p[role=message]').text(msg);
                });
                if ('function' === typeof callback) {
                    myPromise.then((confirm) => {
                        confirm && callback();
                    });
                }
                return myPromise;
            }
        });
    }
});
