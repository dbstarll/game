$.extend($.ro, {
    loader: {
        defaultConfig: {
            name: '匿名用户',
            accept: 'application/json',
            classes: {
                icons: 'material-icons',
                character: 'character-name'
            },
            icons: {
                upload: 'upload',
                refresh: 'refresh',
                save: 'save',
                download: 'download'
            }
        },
        events: {
            refresh: {
                before: 'loader-before-refresh',
                fail: 'loader-refresh-fail',
                done: 'loader-refresh-done'
            },
            save: {
                before: 'loader-before-save',
                fail: 'loader-save-fail',
                done: 'loader-save-done'
            },
            upload: {
                before: 'loader-before-upload',
                fail: 'loader-upload-fail',
                done: 'loader-upload-done'
            },
            download: {
                before: 'loader-before-download',
                fail: 'loader-download-fail',
                done: 'loader-download-done'
            },
            file: {
                before: 'loader-before-file'
            },
            fail: 'loader-fail'
        },
        init: function (loader, config) {
            const upload = this.icon(loader, config, 'upload', {before: this.checkUpload, done: this.decode});
            const file = this.file(loader, config, 'file', {before: this.checkFile}).hide();
            const refresh = this.icon(loader, config, 'refresh', {before: this.checkRefresh, done: this.decode});
            const save = this.icon(loader, config, 'save', {before: this.encode}).hide();
            const download = this.icon(loader, config, 'download', {before: this.checkDownload}).hide();
            const character = $.html.span({role: 'character'}, config.name).addClass(config.classes.character);
            return this.loader($(loader).append(file, save, download, refresh, upload, character), config.handlers, {
                [this.events.file.before]: this.upload, [this.events.refresh.before]: this.refresh,
                [this.events.save.before]: this.save, [this.events.download.before]: this.download,
                [this.events.upload.done]: this.uploaded, [this.events.refresh.done]: this.refreshed,
                [this.events.save.done]: this.saved, [this.events.download.done]: this.downloaded,
                [this.events.refresh.fail]: this.rise, [this.events.download.fail]: this.rise,
                [this.events.save.fail]: this.rise, [this.events.upload.fail]: this.rise,
                [this.events.fail]: this.fail, [this.events.upload.before]: () => {
                    file.trigger('click');
                }
            });
        },
        file: function (loader, config, role, events) {
            return this.bind($.html.input({role: role, type: 'file', accept: config.accept})
                .on('change', function (_, data) {
                    $(this).trigger($.ro.loader.events[role].before, $.extend({}, data, {
                        loader: loader,
                        span: this,
                        files: $(this).prop('files')
                    }));
                    return false;
                }), role, config.handlers, events);
        },
        icon: function (loader, config, role, events) {
            return this.bind($.html.span({role: role}, config.icons[role]).addClass(config.classes.icons)
                .on('click', function (_, data) {
                    $(this).trigger($.ro.loader.events[role].before, $.extend({}, data, {
                        loader: loader,
                        span: this,
                        dirty: {
                            saved: $('span[role=save]', loader).css('display') === 'none',
                            downloaded: $('span[role=download]', loader).css('display') === 'none'
                        }
                    }));
                    return false;
                }), role, config.handlers, events);
        },
        bind: function (span, role, handlers, events) {
            const roleEvents = this.events[role];
            if ('object' === typeof events) {
                Object.keys(events).forEach(function (key) {
                    const handler = events[key];
                    const event = roleEvents[key];
                    if ('function' === typeof handler && 'string' === typeof event) {
                        span.on(event, handler);
                    }
                })
            }
            if ('object' == typeof handlers) {
                const roleHandlers = handlers[role];
                if ('object' == typeof roleHandlers) {
                    Object.keys(roleHandlers).forEach(function (key) {
                        const handler = roleHandlers[key];
                        const event = roleEvents[key];
                        if ('function' === typeof handler && 'string' === typeof event) {
                            span.on(event, handler);
                        }
                    })
                }
            }
            return span;
        },
        loader: function (loader, handlers, events) {
            const that = $.ro.loader;
            if ('object' === typeof events) {
                Object.keys(events).forEach(function (event) {
                    const handler = events[event];
                    if ('function' === typeof handler) {
                        loader.on(event, handler);
                    }
                })
            }
            if ('object' == typeof handlers) {
                Object.keys(handlers).forEach(function (key) {
                    const handler = handlers[key];
                    const event = that.events[key];
                    if ('function' === typeof handler && 'string' === typeof event) {
                        loader.on(event, handler);
                    }
                })
            }
            return loader;
        },
        wrap: function (ui, target) {
            const span = target || $(ui.span);
            const role = span.attr('role');
            return {
                done: function (data, textStatus, jqXHR) {
                    const ext = {jqXHR: jqXHR, textStatus: textStatus};
                    let eventKey;
                    if (data && data.ok) {
                        eventKey = 'done';
                        ext.data = data.data;
                    } else {
                        eventKey = 'fail';
                        ext.data = data;
                    }
                    const event = $.ro.loader.events[role][eventKey];
                    if ('string' === typeof event) {
                        span.trigger(event, $.extend(ext, ui));
                    } else {
                        console.warn('unknown eventKey: ', role, eventKey);
                    }
                },
                fail: function (jqXHR, textStatus, errorThrown) {
                    const event = $.ro.loader.events[role].fail;
                    if ('string' === typeof event) {
                        span.trigger(event, $.extend({
                            jqXHR: jqXHR,
                            textStatus: textStatus,
                            errorThrown: errorThrown
                        }, ui));
                    } else {
                        console.warn('unknown eventKey: ', role, 'fail');
                    }
                }
            };
        },
        confirm: function (ui) {
            const confirm = jsonPath($.ro.loader.getConfig(ui.loader), '$.dialog.confirm.confirm');
            return Array.isArray(confirm) && 'function' === typeof confirm[0] ? confirm[0] : console.warn;
        },
        alert: function (ui) {
            const alert = jsonPath($.ro.loader.getConfig(ui.loader), '$.dialog.alert.alert');
            return Array.isArray(alert) && 'function' === typeof alert[0] ? alert[0] : console.info;
        },
        prompt: function (ui) {
            const prompt = jsonPath($.ro.loader.getConfig(ui.loader), '$.dialog.prompt.prompt');
            return Array.isArray(prompt) && 'function' === typeof prompt[0] ? prompt[0] : console.warn;
        },
        checkUpload: function (_, ui) {
            const confirm = $.ro.loader.confirm(ui);
            const checkDownload = function () {
                if (true !== ui.dirty.downloaded && true !== ui.force) {
                    confirm("上传配置", "有修改的内容未下载，继续上传会忽略新修改的内容，请确认是否需要继续上传？")
                        .then((confirm) => {
                            confirm && $(ui.span).trigger('click', {force: true});
                        });
                    return false;
                }
            };
            if (true !== ui.dirty.saved && true !== ui.force) {
                confirm("上传配置", "有修改的内容未保存，继续上传会忽略新修改的内容，请确认是否需要继续上传？")
                    .then((confirm) => {
                        confirm && false !== checkDownload() && $(ui.span).trigger('click', {force: true});
                    });
                return false;
            } else {
                return checkDownload();
            }
        },
        checkFile: function (_, ui) {
            const file = jsonPath(ui, '$.files.*');
            const alert = $.ro.loader.alert(ui);
            if (!Array.isArray(file) || file.length === 0) {
                alert("上传配置", "请选择文件");
                return false;
            } else if (file[0].size > 10240) {
                alert("上传配置", "文件太大(限制10K)，请重新选择文件");
                return false;
            } else {
                ui.file = file[0];
            }
        },
        checkRefresh: function (_, ui) {
            if (true !== ui.dirty.saved && true !== ui.force) {
                $.ro.loader.confirm(ui)("载入配置", "有修改的内容未保存，继续载入会使新修改的内容丢失，请确认是否需要继续载入？")
                    .then((confirm) => {
                        confirm && $(ui.span).trigger('click', {force: true});
                    });
                return false;
            }
        },
        checkDownload: function (_, ui) {
            if (true !== ui.dirty.saved && true !== ui.force) {
                $.ro.loader.confirm(ui)("下载配置", "有修改的内容未保存，继续下载会忽略新修改的内容，请确认是否需要继续下载？")
                    .then((confirm) => {
                        confirm && $(ui.span).trigger('click', {force: true});
                    });
                return false;
            }
        },
        upload: function (_, ui) {
            const wrap = $.ro.loader.wrap(ui, $(ui.loader).children('span[role=upload]'));
            const fd = new FormData();
            fd.append('player', ui.file);
            $.post({
                url: '/player/upload',
                data: fd,
                cache: false,
                contentType: false,
                processData: false,
                dataType: "json"
            }).done(wrap.done).fail(wrap.fail);
            return false;
        },
        refresh: function (_, ui) {
            const wrap = $.ro.loader.wrap(ui);
            $.get({
                url: '/player/load',
                cache: false,
                dataType: "json"
            }).done(wrap.done).fail(wrap.fail);
            return false;
        },
        save: function (event, ui) {
            const wrap = $.ro.loader.wrap(ui);
            $.post({
                url: '/player/save',
                data: JSON.stringify(ui.player),
                cache: false,
                dataType: "json"
            }).done(wrap.done).fail(wrap.fail);
            return false;
        },
        download: function (_, ui) {
            window.open('/player/download', '_blank');
            $.ro.loader.wrap(ui).done({ok: true});
            return false;
        },
        rise: function (_, ui) {
            $(ui.loader).trigger($.ro.loader.events.fail, $.extend({
                role: $(ui.span).attr('role')
            }, ui));
            return false;
        },
        fail: function (_, ui) {
            const alert = $.ro.loader.alert(ui);
            let action = ui.role;
            switch (ui.role) {
                case 'upload':
                    action = '上传';
                    break;
                case 'refresh':
                    action = '载入';
                    break;
                case 'save':
                    action = '保存';
                    break;
            }
            if (ui.data) {
                alert(action + "配置", action + "失败: [" + ui.data.code + "]" + ui.data.msg);
            } else {
                alert(action + "配置", action + "失败: [" + ui.textStatus + "]" + ui.errorThrown);
            }
        },
        encode: function (_, ui) {
            ui.player = {'character-name': $('span[role=character]', ui.loader).text()};
        },
        decode: function (_, ui) {
            const name = jsonPath(ui.data, '$.player.character-name');
            name && $('span[role=character]', ui.loader).text(name);
        },
        uploaded: function (_, ui) {
            $(ui.loader).children('span[role=save]').hide();
            $(ui.loader).children('span[role=download]').hide();
            $.ro.loader.alert(ui)("上传配置", "上传完成");
            return false;
        },
        refreshed: function (_, ui) {
            $(ui.loader).children('span[role=save]').hide();
            $.ro.loader.alert(ui)("载入配置", "载入完成");
            return false;
        },
        saved: function (_, ui) {
            $(ui.loader).children('span[role=save]').hide();
            $(ui.loader).children('span[role=download]').show();
            $.ro.loader.alert(ui)("保存配置", "保存成功");
            return false;
        },
        downloaded: function (_, ui) {
            $(ui.loader).children('span[role=download]').hide();
            return false;
        },
        getConfig: function (loader) {
            return $.data(loader, 'config');
        }
    }
});

$.fn.extend({
    loader: function (config) {
        if (typeof config === 'object') {
            const finalConfig = $.extend({}, $.ro.loader.defaultConfig, config);
            this.each(function () {
                $.data(this, 'config', finalConfig);
                $.ro.loader.init(this, finalConfig);
            });
        }
        const loader = {that: this};
        return $.extend(loader, {
            refresh: function (force) {
                if (true === force) {
                    this.that.children('span[role=refresh]').trigger('click', {force: true});
                } else {
                    this.that.children('span[role=refresh]').trigger('click');
                }
            },
            change: function () {
                loader.that.children('span[role=save]').show();
            }
        });
    }
});
