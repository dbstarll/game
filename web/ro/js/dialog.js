$.extend($.ro, {
    dialog: {
        init: function (dialog, config) {
            $(dialog).dialog($.extend({
                modal: true,
                autoOpen: false,
                resizable: false,
                close: function () {
                    $(this).off("dialogclose");
                }
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
        return $.extend(confirm, {
            message: function (title, msg, callback) {
                const dialog = message.that;
                dialog.children('p[role=message]').text(msg);
                if ('function' === typeof callback) {
                    dialog.on("dialogclose", callback);
                }
                dialog.dialog("option", "title", title).dialog("open");
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
                const dialog = confirm.that.removeAttr('confirm');
                dialog.children('p[role=message]').text(msg);
                if ('function' === typeof callback) {
                    dialog.on("dialogclose", function () {
                        if (dialog.attr('confirm') === 'true') {
                            setTimeout(callback, 200);
                        }
                    });
                }
                dialog.dialog("option", "title", title).dialog("open");
            }
        });
    }
});
