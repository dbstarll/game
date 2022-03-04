$.extend($.ro, {
    dialog: {
        init: function (dialog, config) {
            return $(dialog).dialog($.extend({
                modal: true,
                autoOpen: false,
                resizable: false
            }, config));
        },
        initAlert: function (dialog, config) {
            return this.init(dialog, $.extend({
                buttons: {
                    "确定": function () {
                        $(this).dialog("close");
                    }
                }
            }, config));
        },
        initConfirm: function (dialog, config) {
            return this.init(dialog, $.extend({
                buttons: {
                    "确定": function () {
                        $(this).attr('confirm', 'true').dialog("close");
                    },
                    "取消": function () {
                        $(this).dialog("close");
                    }
                }
            }, config));
        },
        initPrompt: function (dialog, config) {
            let form;
            const prompt = this.initConfirm(dialog, $.extend({
                close: function () {
                    prompt.off('dialogbeforeclose');
                    form[0].reset();
                }
            }, config));
            form = prompt.find("form").on("submit", function (event) {
                event.preventDefault();
                prompt.attr('confirm', 'true').dialog("close");
            });
            return prompt;
        }
    }
});

$.fn.extend({
    alert: function (config) {
        if (typeof config === 'object') {
            const finalConfig = $.extend({}, config);
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.ro.dialog.initAlert(this, finalConfig);
            });
        }
        const alert = {that: this};
        return $.extend(alert, {
            alert: function (title, msg, callback) {
                let myPromise = new Promise((resolve) => {
                    const dialog = alert.that
                        .dialog("option", "title", title)
                        .dialog("open")
                        .one("dialogclose", () => {
                            resolve(true);
                            return false;
                        });
                    $('[role=message]', dialog).text(msg);
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
                        .one("dialogclose", () => {
                            resolve(dialog.attr('confirm') === 'true');
                            return false;
                        });
                    $('[role=message]', dialog).text(msg);
                });
                if ('function' === typeof callback) {
                    myPromise.then((confirm) => {
                        confirm && callback();
                    });
                }
                return myPromise;
            }
        });
    },
    prompt: function (config) {
        if (typeof config === 'object') {
            const finalConfig = $.extend({}, config);
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.ro.dialog.initPrompt(this, finalConfig);
            });
        }
        const prompt = {that: this};
        return $.extend(prompt, {
            prompt: function (title, msg, callback, initData) {
                const form = prompt.that.find('form');
                const dialog = prompt.that.removeAttr('confirm')
                    .dialog("option", "title", title)
                    .dialog("open")
                    .on("dialogbeforeclose", () => {
                        if (dialog.attr('confirm') === 'true') {
                            dialog.removeAttr('confirm');
                            return callback(JSON.parse(form.jform()));
                        }
                    });
                $('[role=message]', dialog).text(msg);
                if (initData) {
                    form.jform(initData);
                }
            }
        });
    }
});
