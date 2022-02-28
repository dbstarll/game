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
            fail: 'loader-fail'
        },
        init: function (loader, config) {
            const file = $.html.input({type: 'file', accept: config.accept}).hide();
            this.initUploadFile(file);
            const upload = this.icon(loader, config, 'upload', {
                before: this.beforeUpload,
                fail: this.uploadFail,
                done: this.reload
            });
            const refresh = this.icon(loader, config, 'refresh', {
                before: this.beforeRefresh,
                fail: this.refreshFail,
                done: this.reload
            });
            const save = this.icon(loader, config, 'save', {fail: this.saveFail}).hide();
            const download = this.icon(loader, config, 'download', {before: this.beforeDownload}).hide();
            const character = $.html.span({role: 'character'}, config.name).addClass(config.classes.character);
            this.done($(loader).append(file, save, download, refresh, upload, character)
                .on(this.events.refresh.before, this.refresh)
                .on(this.events.download.before, this.download)
                .on(this.events.save.before, this.save)
                .on(this.events.upload.before, this.upload)
                .on(this.events.refresh.fail, this.fail)
                .on(this.events.download.fail, this.fail)
                .on(this.events.save.fail, this.fail)
                .on(this.events.upload.fail, this.fail), config);
        },
        icon: function (loader, config, role, events) {
            const span = $.html.span({role: role}, config.icons[role]).addClass(config.classes.icons);

            const roleEvents = this.events[role];
            span.on('click', function (_, data) {
                span.trigger(roleEvents.before, $.extend({}, data, {
                    loader: loader,
                    span: this,
                    dirty: $.ro.loader.dirty(loader)
                }));
                return false;
            });

            const handlers = config.handlers;
            if ('object' == typeof handlers) {
                const roleHandlers = handlers[role];
                if ('object' == typeof roleHandlers) {
                    ['before', 'fail'].forEach(function (key) {
                        const handler = roleHandlers[key];
                        const event = roleEvents[key];
                        if ('function' === typeof handler && 'string' === typeof event) {
                            span.on(event, handler);
                        }
                    })
                }
            }

            if ('object' === typeof events) {
                Object.keys(events).forEach(function (key) {
                    const handler = events[key];
                    const event = roleEvents[key];
                    if ('function' === typeof handler && 'string' === typeof event) {
                        span.on(event, handler);
                    }
                })
            }
            return span;
        },
        done: function (loader, config) {
            const handlers = config.handlers;
            if ('object' == typeof handlers) {
                ['upload', 'refresh', 'save', 'download'].forEach(function (role) {
                    const roleHandlers = handlers[role];
                    if ('object' == typeof roleHandlers) {
                        const handler = roleHandlers.done;
                        const event = $.ro.loader.events[role].done;
                        if ('function' === typeof handler && 'string' === typeof event) {
                            loader.on(event, handler);
                        }
                    }
                });
                const handler = handlers.fail;
                if ('function' === typeof handler) {
                    loader.on(this.events.fail, handler);
                }
            }
        },
        wrap: function (ui) {
            return {
                done: function (data, textStatus, jqXHR) {
                    const role = $(ui.span).attr('role');
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
                        $(ui.span).trigger(event, $.extend(ext, ui));
                    } else {
                        console.warn('unknown eventKey: ', role, eventKey);
                    }
                },
                fail: function (jqXHR, textStatus, errorThrown) {
                    const role = $(ui.span).attr('role');
                    const event = $.ro.loader.events[role].fail;
                    if ('string' === typeof event) {
                        $(ui.span).trigger(event, $.extend({
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
        initUploadFile: function (file) {
            file.on("change", function () {
                const files = $(this).prop('files');
                if (!files || files.length !== 1) {
                    message("上传配置", "请选择文件");
                    return false;
                }
                const file = files[0];
                if (!file) {
                    message("上传配置", "请选择文件");
                    return false;
                } else if (file.size > 10240) {
                    message("上传配置", "文件太大(限制10K)，请重新选择文件");
                    return false;
                } else {
                    const fd = new FormData();
                    fd.append('player', file);
                    $.post({
                        url: '/player/upload',
                        data: fd,
                        cache: false,
                        contentType: false,
                        processData: false,
                        dataType: "json"
                    }).done(function (data) {
                        if (data && data.ok && data.data) {
                            refresh(data.data);
                            saved();
                            downloaded();
                        }
                    }).fail(function (xhr, status, errorThrown) {
                        message("上传配置", "上传失败: [" + status + "]" + errorThrown);
                    })
                }
            });
        },
        dirty: function (loader) {
            return {
                saved: $('span[role=save]', loader).css('display') === 'none',
                downloaded: $('span[role=download]', loader).css('display') === 'none'
            }
        },
        beforeUpload: function (_, ui) {
            const config = $.ro.loader.getConfig(ui.loader);
            const confirm = jsonPath(config, '$.dialog.confirm.confirm');
            const checkDownload = function () {
                if (true !== ui.dirty.downloaded && true !== ui.force && 'function' === typeof confirm[0]) {
                    confirm[0]("上传配置", "有修改的内容未下载，继续上传会忽略新修改的内容，请确认是否需要继续上传？", function () {
                        $(ui.span).trigger('click', {force: true});
                    });
                    return false;
                }
            }
            if (true !== ui.dirty.saved && true !== ui.force && 'function' === typeof confirm[0]) {
                confirm[0]("上传配置", "有修改的内容未保存，继续上传会忽略新修改的内容，请确认是否需要继续上传？", checkDownload);
                return false;
            } else {
                return checkDownload();
            }
        },
        beforeRefresh: function (_, ui) {
            const config = $.ro.loader.getConfig(ui.loader);
            const confirm = jsonPath(config, '$.dialog.confirm.confirm');
            if (true !== ui.dirty.saved && true !== ui.force && 'function' === typeof confirm[0]) {
                confirm[0]("载入配置", "有修改的内容未保存，继续载入会使新修改的内容丢失，请确认是否需要继续载入？", function () {
                    $(ui.span).trigger('click', {force: true});
                });
                return false;
            }
        },
        beforeDownload: function (_, ui) {
            const config = $.ro.loader.getConfig(ui.loader);
            const confirm = jsonPath(config, '$.dialog.confirm.confirm');
            if (true !== ui.dirty.saved && true !== ui.force && 'function' === typeof confirm[0]) {
                confirm[0]("下载配置", "有修改的内容未保存，继续下载会忽略新修改的内容，请确认是否需要继续下载？", function () {
                    $(ui.span).trigger('click', {force: true});
                });
                return false;
            }
        },
        uploadFail: function () {
        },
        refreshFail: function (_, ui) {
            const config = $.ro.loader.getConfig(ui.loader);
            const message = jsonPath(config, '$.dialog.message.message');
            if ('function' === typeof message[0]) {
                message[0]("载入配置", "载入失败: [" + ui.textStatus + "]" + ui.errorThrown);
            }
        },
        saveFail: function (_, ui) {
            const config = $.ro.loader.getConfig(ui.loader);
            const message = jsonPath(config, '$.dialog.message.message');
            if ('function' === typeof message[0]) {
                message[0]("保存配置", "保存失败: [" + ui.textStatus + "]" + ui.errorThrown);
            }
        },
        reload: function (_, ui) {
            const name = jsonPath(ui.data, '$.player.character-name');
            name && $('span[role=character]', ui.loader).text(name);
        },
        fail: function (_, ui) {
            $(ui.loader).trigger($.ro.loader.events.fail, $.extend({
                role: $(ui.span).attr('role')
            }, ui));
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
        upload: function (_, ui) {
            $(ui.loader).children('input[type=file]').trigger('click');
        },
        download: function () {
            window.open('/player/download', '_blank');
            downloaded();
            return false;
        },
        save: function (event, ui) {
            const wrap = $.ro.loader.wrap(ui);
            const character = $('span[role=character]', ui.loader);
            if (!character.text()) {
                message("保存配置", "请输入角色名称!");
                return false;
            }

            const union = $('#tabs-union');
            let data = {
                'character-name': character.text(),
                manual: $('#manual').buffs().encode(),
                union: {
                    pray: $('.pray', union).buffs().encode(),
                    attack: $('.attack', union).buffs().encode(),
                    defence: $('.defence', union).buffs().encode(),
                    element: $('.element', union).buffs().encode()
                },
                rune: $('#rune').buffs().encode()
            };

            $.post({
                url: '/player/save',
                data: JSON.stringify(data),
                cache: false,
                dataType: "json"
            }).done(wrap.done).fail(wrap.fail);
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
