$(function () {

    function log() {
        if (console && console.log) {
            console.log.apply(console, Array.prototype.slice.call(arguments));
        }
    }

    function extend(cls, extension) {
        var object = Object.create(cls);

        // Copy properties
        for (var key in extension) {
            if (extension.hasOwnProperty(key) || object[key] === "undefined") {
                object[key] = extension[key];
            }
        }

        object.super = function _super() {
            return cls;
        };

        return object;
    }

    /**
     * http://stackoverflow.com/a/30810322
     * @param text
     */
    function copyTextToClipboard(text) {
        var textArea = document.createElement("textarea");

        //
        // *** This styling is an extra step which is likely not required. ***
        //
        // Why is it here? To ensure:
        // 1. the element is able to have focus and selection.
        // 2. if element was to flash render it has minimal visual impact.
        // 3. less flakyness with selection and copying which **might** occur if
        //    the textarea element is not visible.
        //
        // The likelihood is the element won't even render, not even a flash,
        // so some of these are just precautions. However in IE the element
        // is visible whilst the popup box asking the user for permission for
        // the web page to copy to the clipboard.
        //

        // Place in top-left corner of screen regardless of scroll position.
        textArea.style.position = 'fixed';
        textArea.style.top = 0;
        textArea.style.left = 0;

        // Ensure it has a small width and height. Setting to 1px / 1em
        // doesn't work as this gives a negative w/h on some browsers.
        textArea.style.width = '2em';
        textArea.style.height = '2em';

        // We don't need padding, reducing the size if it does flash render.
        textArea.style.padding = 0;

        // Clean up any borders.
        textArea.style.border = 'none';
        textArea.style.outline = 'none';
        textArea.style.boxShadow = 'none';

        // Avoid flash of white box if rendered for any reason.
        textArea.style.background = 'transparent';


        textArea.value = text;

        document.body.appendChild(textArea);

        textArea.select();

        try {
            var successful = document.execCommand('copy');
            var msg = successful ? 'successful' : 'unsuccessful';
            console.log('Copying text command was ' + msg);
        } catch (err) {
            console.log('Oops, unable to copy');
        }

        document.body.removeChild(textArea);
    }

    var Popup = extend(Object, {
        create: function (throneStats, header, content, lifetime) {
            var _this = Object.create(this);
            _this.initialize(throneStats, header, content, lifetime);
            return _this;
        },

        initialize: function (throneStats, header, content, lifetime) {
            log("Showing popup: ", header, content);
            this.throneStats = throneStats;

            this.$element = this.throneStats.$popupTemplate.clone();

            this.view = rivets.bind(this.$element, {
                header: header,
                content: content
            });

            this.$element.css("opacity", 0);

            this.$element.prependTo("#popups");

            this.$element.animate({opacity: 1});

            if (!lifetime) {
                if (this.throneStats.settings.popupLifetime) {
                    lifetime = this.throneStats.settings.popupLifetime;
                }
            }

            if (lifetime) {
                setTimeout(this.close.bind(this), lifetime);
            }
        },

        close: function () {
            log("Removing popup");
            this.view.unbind();
            this.$element.fadeOut({
                complete: function () {
                    this.$element.remove();
                }.bind(this)
            });
        }
    });

    var ThroneStats = extend(Object, {
        create: function () {
            var _this = Object.create(this);
            _this.initialize();
            return _this;
        },

        initialize: function () {
            this.$popupTemplate = $($('#message-template').html());
            this.$configuration = $("#configuration");
            this.$overlay = $("#overlay");

            this.defaultSettings = {
                steamId64: null,
                streamKey: null,
                popupLifetime: 15000,
                dataEndpoint: "data",
            };

            this.susbscribed = false;
            this.settings = null;
            this.configurationView = null;
            this.websocket = null;
        },

        parseSettings: function () {
            var settings = JSON.parse(JSON.stringify(this.defaultSettings));

            var args = String(window.location.hash).substr(1).split("&");

            for (var i = 0, count = args.length; i < count; ++i) {
                var arg = args[i];
                var equals = arg.indexOf("=");
                var key = arg.substr(0, equals);
                settings[key] = arg.substr(equals + 1);
            }

            return settings;
        },

        streamKeyOk: function (streamKey) {
            return /^[a-zA-Z0-9]{6}$/.test(streamKey);
        },

        steamId64Ok: function (steamId64) {
            return /^[0-9]{17}$/.test(steamId64);
        },

        checkSettingsOk: function (settings, showHelp) {

            if (!this.streamKeyOk(settings.streamKey)) {
                log("Stream key looks invalid");
                if (showHelp) {
                    $("#streamKeyHelp").modal("show");
                }
            } else if (!this.steamId64Ok(settings.steamId64)) {
                log("Steam ID 64 looks invalid");
                if (showHelp) {
                    $("#steamId64Help").modal("show");
                }
            } else {
                log("Settings ok");
                return true;
            }

            log("Settings NOT ok");
            return false;
        },

        showConfiguration: function () {
            log("Showing configuration");
            $("body").addClass("background");
            this.$overlay.hide();
            this.$configuration.show();
        },

        showOverlay: function () {
            log("Showing overlay");
            $("body").removeClass("background");
            this.$configuration.hide();
            this.$overlay.show();

            this.subscribe();
        },

        showInformation: function () {
            $("#get-overlay, #get-overlay-tab").removeClass("active");
            $("#information, #information-tab").addClass("active");
        },

        showGetOverlay: function () {
            $("#get-overlay, #get-overlay-tab").addClass("active");
            $("#information, #information-tab").removeClass("active");
        },

        popup: function (header, content, lifetime) {
            Popup.create(this, header, content, lifetime);
        },

        getLink: function (settings) {
            var args = [
                "steamId64=" + settings.steamId64,
                "streamKey=" + settings.streamKey
            ];

            var link = window.location + "#" + args.join("&");

            log("Link: " + link);

            return link;
        },

        getOverlayLink: function () {
            log("Get overlay link");
            if (!this.checkSettingsOk(this.settings, true)) {
                return;
            }

            this.configurationView.models.ui.overlayUrl = this.getLink(this.settings);
            $("#overlay-link").removeClass("hidden");
        },


        copyLink: function () {
            copyTextToClipboard(this.getLink(this.settings));
            this.popup("Link copied", "Overlay link copied to clipboard. CTRL+V to paste.", 2500);
        },

        viewOverlay: function () {
            log("View overlay");

            if (!this.checkSettingsOk(this.settings, true)) {
                return;
            }

            window.history.pushState({}, "ThroneStats Overlay", this.getLink(this.settings));
            this.showOverlay();
        },

        playVideo: function () {
            log("Play video");
            $("#video-preview").hide();
            $("#video").removeClass("hidden");
        },

        start: function () {
            log("Starting up...");

            $(".ui.modal").modal();

            this.settings = this.parseSettings();

            log("Got settings: ", this.settings);

            this.configurationView = rivets.bind(this.$configuration, {
                settings: this.settings,
                ui: {
                    getOverlayLink: this.getOverlayLink.bind(this),
                    viewOverlay: this.viewOverlay.bind(this),
                    showInformation: this.showInformation.bind(this),
                    showGetOverlay: this.showGetOverlay.bind(this),
                    copyLink: this.copyLink.bind(this),
                    playVideo: this.playVideo.bind(this),
                    steamId64HelpVisible: false,
                    streamKeyHelpVisible: false
                },
                globalStats: {
                    mostPopularWeapon: "please wait...",
                    mostCommonCauseOfDeath: "please wait...",
                    mostPopularMutation: "please wait...",
                    mostPopularCrown: "please wait...",
                    mostPopularCharacter: "please wait...",
                    mostCommonDeathLevel: "please wait...",
                }
            });

            this.connect();

            if (this.checkSettingsOk(this.settings)) {
                this.showOverlay();
            } else {
                this.showConfiguration();
            }
        },

        updateGlobalStats: function (data) {
            for (var key in data) {
                if (data.hasOwnProperty(key)) {
                    this.configurationView.models.globalStats[key] = data[key];
                }
            }

            log(data);
        },

        connect: function (callback) {
            var proto = (location.protocol === "https:" ? "wss://" : "ws://");
            var server = proto + window.location.hostname + ":" + window.location.port + "/" + this.settings.dataEndpoint;

            log("Connecting to " + server);

            this.websocket = new WebSocket(server);
            this.websocket.onmessage = this._onMessage.bind(this);
            this.websocket.onerror = this._onError.bind(this);
            this.websocket.onclose = this._onClose.bind(this);
            this.websocket.onopen = function (event) {
                this._onOpen(event);

                if (callback) {
                    callback();
                }
            }.bind(this);
        },

        subscribe: function () {

            if (!this.checkSettingsOk(this.settings)) {
                log("Not subscribing, settings are not OK");
                return;
            }

            log("Subscribing to user " + this.settings.steamId64 + " with stream key " + this.settings.streamKey);

            if (this.websocket.readyState === 1) {
                this.websocket.send(JSON.stringify({
                    type: "subscribe",
                    steamId64: this.settings.steamId64,
                    streamKey: this.settings.streamKey
                }));
                this.susbscribed = true;
            } else {
                setTimeout(this.subscribe.bind(this), 250);
            }
        },

        _onOpen: function (event) {
            log("Connected to server");
        },

        _onMessage: function (event) {
            log("Got message", event.data);

            var data = JSON.parse(event.data);

            if (data.type === "message") {
                this.popup(data.header, data.content);
            } else if (data.type === "globalStats") {
                this.updateGlobalStats(JSON.parse(data.content));
            }
        },

        _onError: function () {
            log("Error", arguments);
        },

        _onClose: function () {
            log("Connection closed!");
            if (this.subscribed) {
                this.popup("Connection issues!", "Trying to reconnect to server.", 2500);
            }
            setTimeout(function () {
                this.connect(function () {
                    this.subscribe();
                }.bind(this));
            }.bind(this), 3000);
        }
    });

    var ts = ThroneStats.create();
    ts.start();

    // For access via console
    window.ts = ts;
});