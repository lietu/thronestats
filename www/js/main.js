$(function () {

    /**
     * Names of all the items we want to show names for
     */
    var NAMES = {
        characters: {
            0: "Random",
            1: "Fish",
            2: "Crystal",
            3: "Eyes",
            4: "Melting",
            5: "Plant",
            6: "Yung Venuz",
            7: "Steroids",
            8: "Robot",
            9: "Chicken",
            10: "Rebel",
            11: "Horror",
            12: "Rogue",
            13: "Big Dog",
            14: "Skeleton",
            15: "Frog",
            16: "Cuz"
        },
        causesOfDeath: {
            0: "Bandit",
            1: "Maggot",
            2: "Rad Maggot",
            3: "Big Maggot",
            4: "Scorpion",
            5: "Gold Scorpion",
            6: "Big Bandit",
            7: "Rat",
            8: "Rat King",
            9: "Green Rat",
            10: "Gator",
            11: "Ballguy",
            12: "Toxic Ballguy",
            13: "Ballguy Mama",
            14: "Assassin",
            15: "Raven",
            16: "Salamander",
            17: "Sniper",
            18: "Big Dog",
            19: "Spider",
            20: "New Cave Thing",
            21: "Laser Crystal",
            22: "Hyper Crystal",
            23: "Snow Bandit",
            24: "Snowbot",
            25: "Wolf",
            26: "Snowtank",
            27: "Lil Hunter",
            28: "Freak",
            29: "Explo Freak",
            30: "Rhino Freak",
            31: "Necromancer",
            32: "Turret",
            33: "Technomancer",
            34: "Guardian",
            35: "Explo Guardian",
            36: "Dog Guardian",
            37: "Throne",
            38: "Throne II",
            39: "Bone Fish",
            40: "Crab",
            41: "Turtle",
            42: "Venus Grunt",
            43: "Venus Sarge",
            44: "Fireballer",
            45: "Super Fireballer",
            46: "Jock",
            47: "Cursed Spider",
            48: "Cursed Crystal",
            49: "Mimic",
            50: "Health Mimic",
            51: "Grunt",
            52: "Inspector",
            53: "Shielder",
            54: "Crown Guardian",
            55: "Explosion",
            56: "Small Explosion",
            57: "Fire Trap",
            58: "Shield",
            59: "Toxin",
            60: "Horror",
            61: "Barrel",
            62: "Toxic Barrel",
            63: "Golden Barrel",
            64: "Car",
            65: "Venus Car",
            66: "Venus Car Fixed",
            67: "Venuz Car 2",
            68: "Icy Car",
            69: "Thrown Car",
            70: "Mine",
            71: "Crown of Death",
            72: "Rogue Strike",
            73: "Blood Launcher",
            74: "Blood Cannon",
            75: "Blood Hammer",
            76: "Disc",
            77: "Curse Eat",
            78: "Big Dog Missile",
            79: "Halloween Bandit",
            80: "Lil Hunter Death",
            81: "Throne Death",
            82: "Jungle Bandit",
            83: "Jungle Assassin",
            84: "Jungle Fly",
            85: "Crown of Hatred",
            86: "Ice Flower",
            87: "Cursed Ammo Pickup",
            88: "Underwater Lightning",
            89: "Elite Grunt",
            90: "Blood Gamble",
            91: "Elite Shielder",
            92: "Elite Inspector",
            93: "Captain",
            94: "Van",
            95: "Buff Gator",
            96: "Generator",
            97: "Lightning Crystal",
            98: "Golden Snowtank",
            99: "Green Explosion",
            100: "Small Generator",
            101: "Golden Disc",
            102: "Big Dog Explosion",
            103: "IDPD Freak",
            104: "Throne II Death",
            105: "Oasis Boss",
            "-1": "Nothing"
        },
        crownChoices: {
            1: "Bare Head",
            2: "Crown of Death",
            3: "Crown of Life",
            4: "Crown of Haste",
            5: "Crown of Guns",
            6: "Crown of Hatred",
            7: "Crown of Blood",
            8: "Crown of Destiny",
            9: "Crown of Love",
            10: "Crown of Risk",
            11: "Crown of Curses",
            12: "Crown of Luck",
            13: "Crown of Protection"
        },
        mutationChoices: {
            0: "Heavy Heart",
            1: "Rhino Skin",
            2: "Extra Feet",
            3: "Plutonium Hunger",
            4: "Rabbit Paw",
            5: "Throne Butt",
            6: "Lucky Shot",
            7: "Bloodlust",
            8: "Gamma Guts",
            9: "Second Stomach",
            10: "Back Muscle",
            11: "Scarier Face",
            12: "Euphoria",
            13: "Long Arms",
            14: "Boiling Veins",
            15: "Shotgun Shoulders",
            16: "Recycle Gland",
            17: "Laser Brain",
            18: "Last Wish",
            19: "Eagle Eyes",
            20: "Impact Wrists",
            21: "Bolt Marrow",
            22: "Stress",
            23: "Trigger Fingers",
            24: "Sharp Teeth",
            25: "Patience",
            26: "Hammer Head",
            27: "Strong Spirit",
            28: "Open Mind"
        },
        weaponChoices: {
            0: "Nothing",
            1: "Revolver",
            2: "Triple Machinegun",
            3: "Wrench",
            4: "Machinegun",
            5: "Shotgun",
            6: "Crossbow",
            7: "Grenade Launcher",
            8: "Double Shotgun",
            9: "Minigun",
            10: "Auto Shotgun",
            11: "Auto Crossbow",
            12: "Super Crossbow",
            13: "Shovel",
            14: "Bazooka",
            15: "Sticky Launcher",
            16: "SMG",
            17: "Assault Rifle",
            18: "Disc Gun",
            19: "Laser Pistol",
            20: "Laser Rifle",
            21: "Slugger",
            22: "Gatling Slugger",
            23: "Assault Slugger",
            24: "Energy Sword",
            25: "Super Slugger",
            26: "Hyper Rifle",
            27: "Screwdriver",
            28: "Laser Minigun",
            29: "Blood Launcher",
            30: "Splinter Gun",
            31: "Toxic Bow",
            32: "Sentry Gun",
            33: "Wave Gun",
            34: "Plasma Gun",
            35: "Plasma Cannon",
            36: "Energy Hammer",
            37: "Jackhammer",
            38: "Flak Cannon",
            39: "Golden Revolver",
            40: "Golden Wrench",
            41: "Golden Machinegun",
            42: "Golden Shotgun",
            43: "Golden Crossbow",
            44: "Golden Grenade Launcer",
            45: "Golden Laser Pistol",
            46: "Chicken Sword",
            47: "Nuke Launcher",
            48: "Ion Cannon",
            49: "Quadruple Machinegun",
            50: "Flamethrower",
            51: "Dragon",
            52: "Flare Gun",
            53: "Energy Screwdriver",
            54: "Hyper Launcher",
            55: "Laser Cannon",
            56: "Rusty Revolver",
            57: "Lightning Pistol",
            58: "Lightning Rifle",
            59: "Lightning Shotgun",
            60: "Super Flak Cannon",
            61: "Sawed-off Shotgun",
            62: "Splinter Pistol",
            63: "Super Splinter Gun",
            64: "Lighting SMG",
            65: "Smart Gun",
            66: "Heavy Crossbow",
            67: "Blood Hammer",
            68: "Lightning Cannon",
            69: "Pop Gun",
            70: "Plasma Rifle",
            71: "Pop Rifle",
            72: "Toxic Launcher",
            73: "Flame Cannon",
            74: "Lightning Hammer",
            75: "Flame Shotgun",
            76: "Double Flame Shotgun",
            77: "Auto Flame Shotgun",
            78: "Cluster Launcher",
            79: "Grenade Shotgun",
            80: "Grenade Rifle",
            81: "Rogue Rifle",
            82: "Party Gun",
            83: "Double Minigun",
            84: "Gatling Bazooka",
            85: "Auto Grenade Shotgun",
            86: "Ultra Revolver",
            87: "Ultra Laser Pistol",
            88: "Sledgehammer",
            89: "Heavy Revolver",
            90: "Heavy Machinegun",
            91: "Heavy Slugger",
            92: "Ultra Shovel",
            93: "Ultra Shotgun",
            94: "Ultra Crossbow",
            95: "Ultra Grenade Launcher",
            96: "Plasma Minigun",
            97: "Devastator",
            98: "Golden Plasma Gun",
            99: "Golden Slugger",
            100: "Golden Splinter Gun",
            101: "Golden Screwdriver",
            102: "Golden Bazooka",
            103: "Golden Assault Rifle",
            104: "Super Disc Gun",
            105: "Heavy Auto Crossbow",
            106: "Heavy Assault Rifle",
            107: "Blood Cannon",
            108: "Dog Spin Attack",
            109: "Dog Missile",
            110: "Incinerator",
            111: "Super Plasma Cannon",
            112: "Seeker Pistol",
            113: "Seeker Shotgun",
            114: "Eraser",
            115: "Guitar",
            116: "Bouncer SMG",
            117: "Bouncer Shotgun",
            118: "Hyper Slugger",
            119: "Super Bazooka",
            120: "Frog Pistol",
            121: "Black Sword",
            122: "Golden Nuke Launcher",
            123: "Golden Disc Gun",
            124: "Heavy Grenade Launcher",
            125: "Gun Gun",
            201: "Golden Frog Pistol"
        }
    };

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
                statsSteamId64: null,
                streamKey: null,
                popupLifetime: 15000,
                dataEndpoint: "data",
                view: "information"
            };

            this.stats = {};
            this.globalStats = {};

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
                stats: this.stats,
                ui: {
                    getOverlayLink: this.showOverlayLink.bind(this),
                    viewOverlay: this.viewOverlay.bind(this),
                    showInformation: function () {
                        this.showView("information");
                    }.bind(this),
                    showGetOverlay: function () {
                        this.showView("get-overlay");
                    }.bind(this),
                    showStats: function () {
                        this.showView("stats");
                    }.bind(this),
                    showPlayerStats: this.showPlayerStats.bind(this),
                    showGlobalStats: function() {
                        this._updateStats("Global run data", this.globalStats);
                    }.bind(this),
                    copyLinkToClipboard: this.copyLinkToClipboard.bind(this),
                    playVideo: this.playVideo.bind(this),
                    goToView: this.goToView.bind(this),
                    steamId64HelpVisible: false,
                    streamKeyHelpVisible: false
                },
                globalStats: this.globalStats
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
        updateSettings: function (settings) {
            this.settings = settings;

            if (this.settings.view === null) {
                this.settings.view = this.defaultSettings.view;
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
         * Merge settings so no existing data is lost
         *
         * @param old
         * @param settings
         */
        mergeSettings: function (old, settings) {
            var newSettings = {};
            var key;

            for (key in settings) {
                if (settings.hasOwnProperty(key)) {
                    newSettings[key] = settings[key];
                }
            }

            for (key in old) {
                if (old.hasOwnProperty(key) && !newSettings[key]) {
                    newSettings[key] = old[key];
                }
            }

            return newSettings;
        },

        /**
         * Get current location, hash-free
         *
         * @return {string}
         */
        getLocation: function () {
            var location = String(window.location);
            var pos = location.indexOf("#");
            if (pos !== -1) {
                location = location.substr(0, pos);
            }

            return location;
        },

        /**
         * Convert settings object to an URL
         *
         * @param settings
         */
        settingsToUrl: function (settings) {
            settings = settings || this.settings;
            var args = [];

            for (var key in settings) {
                if (settings.hasOwnProperty(key) && key != "" && key.substr(0, 1) !== "_") {
                    if (settings[key] !== this.defaultSettings[key]) {
                        args.push(key + "=" + settings[key]);
                    }
                }
            }

            var url = this.getLocation();

            if (args.length > 0) {
                url = url + "#" + args.join("&");
            }

            log("URL: " + url);

            return url
        },

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

            var link = this.getLocation() + "#" + args.join("&");

            log("Overlay link: " + link);

            return link;
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
        showView: function (view) {
            if (view) {
                this.settings.view = view;
            }

            log("Was asked to show view " + this.settings.view);

            if (this.settings.view === "overlay") {
                if (this.checkOverlaySettings(this.settings, true)) {
                    this.showOverlay();
                    return;
                } else {
                    log("Invalid settings prevent showing overlay, showing get-overlay instead");
                    this.settings.view = "get-overlay";
                }
            }

            this._showContentView(this.settings.view);
            var url = this.settingsToUrl();

            if (url !== String(window.location)) {
                log("Pushing to history: " + url);
                this.addToHistory(url);
            }
        },

        /**
         * Click on a link that's supposed to go to a view
         */
        goToView: function (event) {
            var $target = $(event.target);

            var targetSettings = this.parseSettings($target.attr("href"));

            this.updateSettings(this.mergeSettings(this.settings, targetSettings));
            this.showView();

            // Never do what the link normally does -> navigate uncontrollably
            event.preventDefault();
            return false;
        },

        /**
         * Add something to browser navigation history
         *
         * @param url
         */
        addToHistory: function (url) {
            if (this.historyTimeout !== null) {
                clearTimeout(this.historyTimeout);
            }

            setTimeout(function () {
                this.historyTimeout = null;
                window.history.pushState({}, "ThroneStats - " + this.settings.view, url);
            }.bind(this), 50);
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

            this.addToHistory(this.getOverlayLink(this.settings));
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
         * "Show player stats" -button
         */
        showPlayerStats: function() {
            var steamId64 = this.settings.statsSteamId64;

            if (steamId64 === null) {
                this.popup("Error", "Enter a SteamID64");
            } else if (this.steamId64Ok(steamId64)) {
                this._requestPlayerStats(steamId64);
                this._updateStats("Global run data", this.globalStats);
            } else {
                this.popup("Error", "SteamID64 " + steamId64 + " looks invalid.");
            }
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

        /**
         * Calculate percentages
         *
         * @param totalRuns
         * @param runs
         */
        getPercentage: function (totalRuns, runs) {
            if (totalRuns === 0 || runs === 0) {
                return "0.00%";
            }

            return String(((runs / totalRuns) * 100).toFixed(2)) + "%"
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
                $(".tabular.menu .item").removeClass("active");
                $("#information, #get-overlay, #stats").removeClass("active");

                var tab = "#" + view + "-tab";
                var selector = "#" + view + ", " + tab;

                $(selector).addClass("active");

                $("html, body").animate({
                    scrollTop: $(tab).offset().top
                }, 150);

                this.activeView = view;
            }
        },

        /**
         * Process a global stats update
         *
         * @param data
         */
        _updateGlobalStats: function (data) {
            if (this.settings.statsSteamId64 === null) {
                this._updateStats("Global run data", data);
            }

            for (var key in data) {
                if (data.hasOwnProperty(key)) {
                    this.globalStats[key] = data[key];
                }
            }
        },

        /**
         * Update statistical data from given data, calculates percentages
         *
         * @param steamIdText
         * @param data
         * @private
         */
        _updateStats: function (steamIdText, data) {
            this.stats.steamIdText = steamIdText;

            var skip = {
                "characters": ["", "0", "13", "14", "16"],
                "causesOfDeath": [""],
                "crownChoices": ["", "1"],
                "deathsByLevel": [""],
                "mutationChoices": [""],
                "weaponChoices": ["", "0"]
            };

            for (var key in data) {
                if (data.hasOwnProperty(key)) {
                    switch (key) {
                        case "characters":
                        case "causesOfDeath":
                        case "crownChoices":
                        case "deathsByLevel":
                        case "mutationChoices":
                        case "weaponChoices":
                            this.stats[key] = [];
                            for (var item in data[key]) {
                                if (data[key].hasOwnProperty(item)) {
                                    if (skip[key].indexOf(item) !== -1) {
                                        continue;
                                    }

                                    var name = item;
                                    if (key != "deathsByLevel") {
                                        name = NAMES[key][item];
                                    }

                                    this.stats[key].push({
                                        name: name,
                                        runs: data[key][item],
                                        percentage: this.getPercentage(data.runs, data[key][item])
                                    })
                                }
                            }
                            break;

                        default:
                            this.stats[key] = data[key];
                    }
                }
            }
        },

        /**
         * Request player stats from the server.
         * @param steamId64
         * @private
         */
        _requestPlayerStats: function(steamId64) {
            log("Requesting stats for " + steamId64);
            this.websocket.send(JSON.stringify({
                type: "requestStats",
                steamId64: steamId64
            }));
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
        _onHashChange: function (event) {
            var newURL = event.newURL;

            log("Hash change detected, new URL: " + newURL);

            var pos = newURL.indexOf("#");
            if (pos !== -1) {
                var hash = newURL.substr(pos);
                this.updateSettings(this.mergeSettings(this.settings, this.parseSettings(hash)));
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
            } else if (data.type === "stats") {
                var content = JSON.parse(data.content);
                log("Got stats update for " + data.header, content);
                this._updateStats("Player " + data.header, content);
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