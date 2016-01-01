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
     * Rate limit a function call, but execute all calls
     *
     * @param {Function} fn Function to wrap
     * @param {Number} wait Milliseconds between executions
     * @returns {Function} Rate-limited wrapper to fn
     */
    function throttle(fn, wait) {
        var _queue = [];
        var _running = false;

        function _call(scope, args) {
            _running = true;
            fn.apply(scope, args);

            setTimeout(function () {
                if (_queue.length) {
                    var next = _queue.shift();
                    _call(next.scope, next.args);
                } else {
                    _running = false;
                }
            }, wait);
        }

        return function () {
            var args = Array.prototype.slice.call(arguments);
            var _this = this;

            if (_running) {
                _queue.push({scope: _this, args: args});
            } else {
                _call(_this, args);
            }
        };
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

            this.$element.css("display", "none");
            this.$element.prependTo("#popups");
            this.$element.slideDown();

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
            this.$content = $("#content");
            this.$overlay = $("#overlay");

            this.defaultSettings = {
                steamId64: null,
                streamKey: null,
                popupLifetime: 15000,
                dataEndpoint: "data",
                view: "information"
            };

            this.view = null;
            this.subscribed = false;
            this.settings = null;
            this.contentView = null;
            this.websocket = null;
        },

        /**
         * Start the webapp
         */
        start: function () {
            log("Starting up...");

            // Setup Semantic UI modals
            $(".ui.modal").modal();

            this.updateSettings(this.parseSettings());

            log("Got settings: ", this.settings);

            this.contentView = rivets.bind(this.$content, {
                settings: this.settings,
                ui: {
                    getOverlayLink: this.showOverlayLink.bind(this),
                    viewOverlay: this.viewOverlay.bind(this),
                    showInformation: function() { this.showView("information"); }.bind(this),
                    showGetOverlay: function() { this.showView("get-overlay"); }.bind(this),
                    copyLinkToClipboard: this.copyLinkToClipboard.bind(this),
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

            this.showView();
            window.addEventListener("hashchange", this._onHashChange.bind(this), false);
        },

        /**
         * Update data from new settings
         *
         * @param settings
         */
        updateSettings: function(settings) {
            this.settings = settings;

            if (this.view === null && !settings.view) {
                this.view = this.defaultSettings.view;
            } else {
                this.view = settings.view;
            }
        },

        /**
         * Parse settings from URL hash
         *
         * @param {String} hash
         */
        parseSettings: function (hash) {
            hash = hash || window.location.hash;

            var settings = JSON.parse(JSON.stringify(this.defaultSettings));

            var args = String(hash).substr(1).split("&");

            for (var i = 0, count = args.length; i < count; ++i) {
                var arg = args[i];
                var equals = arg.indexOf("=");
                var key = arg.substr(0, equals);
                settings[key] = arg.substr(equals + 1);
            }

            return settings;
        },

        /**
         * Check if overlay settings are good to go, optionally displays help.
         *
         * @param settings
         * @param showHelp
         * @returns {boolean}
         */
        checkOverlaySettings: function (settings, showHelp) {

            if (!this.streamKeyOk(settings.streamKey)) {
                log("Stream key " + settings.streamKey + " looks invalid");
                if (showHelp) {
                    $("#streamKeyHelp").modal("show");
                }
            } else if (!this.steamId64Ok(settings.steamId64)) {
                log("Steam ID 64 " + settings.steamId64 + " looks invalid");
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

        /**
         * Display a view
         *
         * @param {String} view "overlay" or one of the content area tab names
         */
        showView: function(view) {
            if (view) {
                this.view = view;
            }

            log("Was asked to show view " + this.view);

            if (this.view === "overlay") {
                if (this.checkOverlaySettings(this.settings, true)) {
                    this.showOverlay();
                    return;
                } else {
                    log("Invalid settings prevent showing overlay, showing get-overlay instead");
                    this.view = "get-overlay";
                }
            }

            this._showContentView(this.view);
        },

        /**
         * Show the stats overlay
         */
        showOverlay: function () {
            log("Showing overlay");

            $("body").removeClass("background");

            this.$content.hide();
            this.$overlay.show();

            this.subscribe();
        },

        /**
         * Show a popup message
         *
         * @param {String} header Title text
         * @param {String} content Content body
         * @param {Number} lifetime Time in milliseconds to display popup
         */
        popup: throttle(function (header, content, lifetime) {
            Popup.create(this, header, content, lifetime);
        }, 1500),

        /**
         * Get URL to overlay regardless of current settings
         *
         * @param settings
         * @returns {string}
         */
        getOverlayLink: function (settings) {
            var args = [
                "view=overlay",
                "steamId64=" + settings.steamId64,
                "streamKey=" + settings.streamKey
            ];

            var location = String(window.location);
            var pos = location.indexOf("#");
            if (pos !== -1) {
                location = location.substr(0, pos);
            }

            var link = location + "#" + args.join("&");

            log("Link: " + link);

            return link;
        },

        /**
         * Get overlay link -button.
         *
         * Display the overlay link in the UI for user to copy.
         */
        showOverlayLink: function () {
            log("Show overlay link");

            if (!this.checkOverlaySettings(this.settings, true)) {
                return;
            }

            this.contentView.models.ui.overlayUrl = this.getOverlayLink(this.settings);
            $("#overlay-link").removeClass("hidden");
        },

        /**
         * View the overlay -button.
         *
         * Checks the settings and if they're OK navigates to the overlay view.
         */
        viewOverlay: function () {
            log("View overlay");

            if (!this.checkOverlaySettings(this.settings, true)) {
                return;
            }

            window.history.pushState({}, "ThroneStats Overlay", this.getOverlayLink(this.settings));
            this.showView("overlay");
        },

        /**
         * Copy link to clipboard
         */
        copyLinkToClipboard: function () {
            copyTextToClipboard(this.getOverlayLink(this.settings));
            this.popup("Link copied", "Overlay link copied to clipboard. CTRL+V to paste.", 2500);
        },

        /**
         * Play the video (click on screenshot)
         */
        playVideo: function () {
            log("Play video");
            $("#video-preview").hide();
            $("#video").removeClass("hidden");
        },

        /**
         * Connect to backend
         *
         * @param callback
         */
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

        /**
         * Subscribe to receive live stats events from backend
         */
        subscribe: function () {

            if (!this.checkOverlaySettings(this.settings)) {
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
                this.subscribed = true;
            } else {
                setTimeout(this.subscribe.bind(this), 250);
            }
        },

        /**
         * Check if stream key looks good
         *
         * @param streamKey
         * @returns {boolean}
         */
        streamKeyOk: function (streamKey) {
            return /^[a-zA-Z0-9]{6}$/.test(streamKey);
        },

        /**
         * Check if Steam ID 64 looks good
         *
         * @param steamId64
         * @returns {boolean}
         */
        steamId64Ok: function (steamId64) {
            return /^[0-9]{17}$/.test(steamId64);
        },



        /*
         * Private methods, shouldn't be called unless you know what you're doing
         */


        /**
         * Show a content view, should NOT be called outside of showView()
         *
         * @param {String} view
         * @private
         */
        _showContentView: function (view) {
            log("Showing content view " + view);

            $("body").addClass("background");

            this.$overlay.hide();
            this.$content.show();

            this._switchTab(view);
        },


        /**
         * Switch active tab in content view
         * @param {String} view
         */
        _switchTab: function (view) {
            if (view != this.activeView) {
                $("#get-overlay, #get-overlay-tab").removeClass("active");
                $("#information, #information-tab").removeClass("active");

                var selector = "#" + view + ", #" + view + "-tab";
                $(selector).addClass("active");

                this.activeView = view;
            }
        },

        /**
         * Process a global stats update
         *
         * @param data
         */
        _updateGlobalStats: function (data) {
            for (var key in data) {
                if (data.hasOwnProperty(key)) {
                    this.contentView.models.globalStats[key] = data[key];
                }
            }
        },

        /*
         * Event handlers, should never be called manually
         */


        /**
         * Triggered when location hash changes
         *
         * @param event
         * @private
         */
        _onHashChange: function(event) {
            var newURL = event.newURL;

            log("Hash change detected, new URL: " + newURL);

            var pos = newURL.indexOf("#");
            if (pos !== -1) {
                var hash = newURL.substr(pos);
                this.updateSettings(this.parseSettings());
                this.showView();
            }
        },

        /**
         * Triggered when connection to backend is established
         *
         * @param event
         * @private
         */
        _onOpen: function (event) {
            log("Connected to server");
        },

        /**
         * Triggered when backend sends messages to us
         *
         * @param event
         * @private
         */
        _onMessage: function (event) {
            var data = JSON.parse(event.data);

            if (data.type === "message") {
                log("Got message", data);
                this.popup(data.header, data.content);
            } else if (data.type === "globalStats") {
                var content = JSON.parse(data.content);
                log("Got global stats update", content);
                this._updateGlobalStats(content);
            } else {
                log("Got unsupported message?", data);
            }
        },

        /**
         * Triggered when there is an error with the backend connection
         *
         * @private
         */
        _onError: function () {
            log("Error", arguments);
        },

        /**
         * Triggered when connection to backend was closed
         *
         * @private
         */
        _onClose: function () {
            log("Connection closed!");
            if (this.subscribed) {
                this.popup("Connection issues!", "Trying to reconnect to server.", 2500);
            }
            setTimeout(function () {
                this.connect(function () {
                    if (this.subscribed) {
                        this.subscribe();
                    }
                }.bind(this));
            }.bind(this), 3000);
        }
    });

    var ts = ThroneStats.create();
    ts.start();

    // For access via console
    window.ts = ts;
});