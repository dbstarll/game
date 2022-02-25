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
        span: function (loader, config, role, events) {
            const roleEvents = this.events[role];
            const span = $.html.span({role: role}, config.icons[role]).addClass(config.classes.icons)
                .on('click', function (event, force) {
                    $(this).trigger(roleEvents.before, {
                        loader: loader,
                        span: this,
                        dirty: $.extend({force: force}, $.ro.loader.dirty(loader))
                    });
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
                            $(loader).on(event, handler);
                        }
                    }
                });
                const handler = handlers.fail;
                if ('function' === typeof handler) {
                    $(loader).on(this.events.fail, handler);
                }
            }
        },
        init: function (loader, config) {
            const file = $.html.input({type: 'file', accept: config.accept}).hide();
            this.initUploadFile(file);
            const upload = this.span(loader, config, 'upload', {done: this.reload});
            const refresh = this.span(loader, config, 'refresh', {done: this.reload});
            const save = this.span(loader, config, 'save').hide();
            const download = this.span(loader, config, 'download').hide();
            const character = $.html.span({role: 'character'}, config.name).addClass(config.classes.character);
            $(loader).append(file, save, download, refresh, upload, character)
                .on(this.events.refresh.before, this.refresh)
                .on(this.events.download.before, this.download)
                .on(this.events.save.before, this.save)
                .on(this.events.upload.before, this.upload)
                .on(this.events.refresh.fail, this.fail)
                .on(this.events.download.fail, this.fail)
                .on(this.events.save.fail, this.fail)
                .on(this.events.upload.fail, this.fail);
            this.done(loader, config);
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
        reload: function (event, ui) {
            const name = jsonPath(ui.data, '$.player.character-name');
            name && $('span[role=character]', ui.loader).text(name);
        },
        fail: function (event, ui) {
            $(ui.loader).trigger($.ro.loader.events.fail, $.extend({
                role: $(ui.span).attr('role')
            }, ui));
            return false;
        },
        refresh: function (event, ui) {
            const wrap = $.ro.loader.wrap(ui);
            $.get({
                url: '/player/load',
                cache: false,
                dataType: "json"
            }).done(wrap.done).fail(wrap.fail);
            return false;
        },
        upload: function () {
            const fn = function () {
                $('#tool-bar > input[type=file]').click();
            }
            const fn2 = function () {
                if (needDownload()) {
                    confirm("上传配置", "有修改的内容未下载，继续上传会忽略新修改的内容，请确认是否需要继续上传？", fn);
                } else {
                    return fn();
                }
            }
            if (needSave()) {
                confirm("上传配置", "有修改的内容未保存，继续上传会忽略新修改的内容，请确认是否需要继续上传？", fn2);
            } else {
                return fn2();
            }
            return false;
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
        getConfig: function (buff) {
            return $.data(buff, 'config');
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
            })
        }
        const loader = {that: this};
        return $.extend(loader, {
            refresh: function (force) {
                this.that.children('span[role=refresh]').trigger('click', force);
            },
            download: function (force) {
                this.that.children('span[role=download]').trigger('click', force);
            },
            change: function () {
                this.that.children('span[role=save]').show();
            }
        });
    }
});
