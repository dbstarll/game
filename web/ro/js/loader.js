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
            beforeRefresh: 'loader-before-refresh',
            beforeDownload: 'loader-before-download',
            beforeSave: 'loader-before-save',
            refreshFail: 'loader-refresh-fail',
            saveFail: 'loader-save-fail',
            refresh: 'loader-refresh',
            save: 'loader-save',
            fail: 'loader-fail'
        },
        wrap: function (ui) {
            return {
                done: function (data, textStatus, jqXHR) {
                    let eventKey = $(ui.span).attr('role');
                    const ext = {jqXHR: jqXHR, textStatus: textStatus};
                    if (data && data.ok) {
                        ext.data = data.data;
                    } else {
                        eventKey += 'Fail';
                        ext.data = data;
                    }
                    const event = $.ro.loader.events[eventKey];
                    if ('string' === typeof event) {
                        $(ui.span).trigger(event, $.extend(ext, ui));
                    } else {
                        console.warn('unknown eventKey: ', eventKey);
                    }
                },
                fail: function (jqXHR, textStatus, errorThrown) {
                    const eventKey = $(ui.span).attr('role') + 'Fail';
                    const event = $.ro.loader.events[eventKey];
                    if ('string' === typeof event) {
                        $(ui.span).trigger(event, $.extend({
                            jqXHR: jqXHR,
                            textStatus: textStatus,
                            errorThrown: errorThrown
                        }, ui));
                    } else {
                        console.warn('unknown eventKey: ', eventKey);
                    }
                }
            };
        },
        before: function (loader, event) {
            return function (_, force) {
                $(this).trigger(event, {
                    loader: loader,
                    span: this,
                    dirty: $.extend({force: force}, $.ro.loader.dirty(loader))
                });
                return false;
            }
        },
        init: function (loader, config) {
            const that = this;
            const file = $.html.input({type: 'file', accept: config.accept}).hide();
            this.initUploadFile(file);
            const upload = $.html.span({role: 'upload'}, config.icons.upload).addClass(config.classes.icons).on('click', this.upload);
            const refresh = $.html.span({role: 'refresh'}, config.icons.refresh).addClass(config.classes.icons)
                .on('click', $.ro.loader.before(loader, that.events.beforeRefresh))
                .on(that.events.refresh, that.reload);
            if ('function' === typeof config.beforeRefresh) {
                refresh.on(that.events.beforeRefresh, config.beforeRefresh);
            }
            if ('function' === typeof config.refreshFail) {
                refresh.on(that.events.refreshFail, config.refreshFail);
            }
            const save = $.html.span({role: 'save'}, config.icons.save).addClass(config.classes.icons)
                .on('click', $.ro.loader.before(loader, that.events.beforeSave)).hide();
            if ('function' === typeof config.beforeSave) {
                save.on(that.events.beforeSave, config.beforeSave);
            }
            if ('function' === typeof config.saveFail) {
                refresh.on(that.events.saveFail, config.saveFail);
            }
            const download = $.html.span({role: 'download'}, config.icons.download).addClass(config.classes.icons)
                .on('click', $.ro.loader.before(loader, that.events.beforeDownload)).hide();
            if ('function' === typeof config.beforeDownload) {
                download.on(that.events.beforeDownload, config.beforeDownload);
            }
            const character = $.html.span({role: 'character'}, config.name).addClass(config.classes.character);
            $(loader).append(file, save, download, refresh, upload, character)
                .on(that.events.beforeRefresh, that.refresh)
                .on(that.events.beforeDownload, that.download)
                .on(that.events.beforeSave, that.save)
                .on(that.events.refreshFail, that.fail)
                .on(that.events.saveFail, that.fail);
            if ('function' === typeof config.fail) {
                $(loader).on(that.events.fail, config.fail);
            }
            if ('function' === typeof config.refresh) {
                $(loader).on(that.events.refresh, config.refresh);
            }
            if ('function' === typeof config.save) {
                $(loader).on(that.events.save, config.save);
            }
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
        },
        download: function () {
            window.open('/player/download', '_blank');
            downloaded();
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
            save: function (force) {
                this.that.children('span[role=save]').trigger('click', force);
            },
            change: function () {
                this.that.children('span[role=save]').show();
            }
        });
    }
});
