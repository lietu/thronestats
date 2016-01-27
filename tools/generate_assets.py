"""
Generate the graphics assets for ../www/img/ -folders.
"""

from argparse import ArgumentParser
import re
import os
import sys
import shutil
import subprocess
from pathlib import Path
from pprint import pprint

TEST_TEMPLATE = """
<html>
<head>
<style>
i {{
  display: block;
  height: 32px;
  width: 32px;
  image-rendering: optimizeSpeed;             /* Legal fallback */
  image-rendering: -moz-crisp-edges;          /* Firefox        */
  image-rendering: -o-crisp-edges;            /* Opera          */
  image-rendering: -webkit-optimize-contrast; /* Safari         */
  image-rendering: optimize-contrast;         /* CSS3 Proposed  */
  image-rendering: crisp-edges;               /* CSS4 Proposed  */
  image-rendering: pixelated;                 /* CSS4 Proposed  */
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}}
</style>
</head>
<body>
<table>
<thead><tr><th>Image</th><th>Name</th></tr></thead>
<tbody>{rows}</tbody>
</table>
</body>
</html>"""

FOLDERS = {
    "characters": "../www/img/characters",
    "causesOfDeath": "../www/img/causesOfDeath",
    "crownChoices": "../www/img/crownChoices",
    "deathsByLevel": "../www/img/deathsByLevel",
    "mutationChoices": "../www/img/mutationChoices",
    "weaponChoices": "../www/img/weaponChoices",
}

DATA = {
    "causesOfDeath": {
        0: "Bandit",
        1: "Maggot",
        2: "Rad Maggot",
        3: "Big Maggot",
        4: "Scorpion",
        5: "Gold Scorpion",
        6: "Bandit Boss",
        7: "Rat",
        8: "Rat King",
        9: "Fast Rat",
        10: "Gator",
        11: "Exploder",
        12: "Super Frog",
        13: "Frog Queen",
        14: "Melee",
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
        37: "Nothing",
        38: "Nothing 2",
        39: "Bone Fish",
        40: "Crab",
        41: "Turtle",
        42: "Molefish",
        43: "Molesarge",
        44: "Fireballer",
        45: "Super Fireballer",
        46: "Jock",
        47: "Inv Spider",
        48: "Inv Laser Crystal",
        49: "Mimic",
        50: "Super Mimic",
        51: "Grunt",
        52: "Inspector",
        53: "Shielder",
        54: "Crown Guardian",
        55: "Explosion",
        56: "Small Explosion",
        57: "Trap Fire",
        58: "Shield",
        59: "Toxic Gas",
        60: "Horror",
        61: "Barrel",
        62: "Toxic Barrel",
        63: "Gold Barrel",
        64: "Car",
        65: "Venus Car",
        66: "Venus Car Fixed",
        67: "Venuz Car 2",
        68: "Frozen Car",
        69: "Frozen Car Thrown",
        70: "Mine",
        71: "Crown2",
        72: "Rogue Strike",
        73: "Blood Nader",
        74: "Blood Cannon",
        75: "Blood Hammer",
        76: "Disc",
        77: "Curse",
        78: "Scrap Boss Missile",
        79: "Spooky Bandit",
        80: "Lil Hunter Dead",
        81: "Nothing Death",
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
        93: "Last",
        94: "Van",
        95: "Buff Gator",
        96: "Generator",
        97: "Lightning Crystal",
        98: "Gold Tank",
        99: "Green Explosion",
        100: "Small Generator",
        101: "Gold Disc",
        102: "Scrap Boss Dead",
        103: "Popo Freak",
        104: "Nothing 2 Death",
        105: "Oasis Boss",
    },
    "weaponChoices": {
        1: "Revolver",
        2: "Triple Machinegun",
        3: "Wrench",
        4: "Machinegun",
        5: "Shotgun",
        6: "Crossbow",
        7: "Nader",
        8: "Super Shotgun",
        9: "Minigun",
        10: "Auto Shotgun",
        11: "Auto Crossbow",
        12: "Super Crossbow",
        13: "Shovel",
        14: "Bazooka",
        15: "Sticky Nader",
        16: "SMG",
        17: "A Rifle",
        18: "Disc Gun",
        19: "Laser Gun",
        20: "Laser Rifle",
        21: "Slugger",
        22: "Gatling Slugger",
        23: "Assault Slugger",
        24: "Energy Sword",
        25: "Super Slugger",
        26: "Hyper Rifle",
        27: "Screwdriver",
        28: "Laser Minigun",
        29: "Blood Nader",
        30: "Splinter Gun",
        31: "Toxic Bow",
        32: "Sentry Gun",
        33: "Wave Gun",
        34: "Plasma Gun",
        35: "Plasma Cannon",
        36: "Energy Hammer",
        37: "Jackhammer",
        38: "Flak Cannon",
        39: "Gold Revolver",
        40: "Gold Wrench",
        41: "Gold Machinegun",
        42: "Gold Shotgun",
        43: "Gold Crossbow",
        44: "Gold Nader",
        45: "Gold Laser Gun",
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
        61: "Sawn Off Shotgun",
        62: "Splinter Pistol",
        63: "Heavy Splinter Gun",
        64: "Lightning SMG",
        65: "Smart Gun",
        66: "Heavy Crossbow",
        67: "Blood Hammer",
        68: "Lightning Cannon",
        69: "Pop Gun",
        70: "Plasma Rifle",
        71: "Pop Rifle",
        72: "Toxic Nader",
        73: "Flame Cannon",
        74: "Lightning Hammer",
        75: "Flame Shotgun",
        76: "Double Flame Shotgun",
        77: "Auto Flame Shotgun",
        78: "Cluster Launcher",
        79: "Nade Shotgun",
        80: "Nade Rifle",
        81: "Rogue Rifle",
        82: "Party Gun",
        83: "Double Minigun",
        84: "Gatling Bazooka",
        85: "Auto Nade Shotgun",
        86: "Ultra Revolver",
        87: "Ultra Laser Gun",
        88: "Hammer",
        89: "Heavy Revolver",
        90: "Heavy Machinegun",
        91: "Heavy Slugger",
        92: "Ultra Shovel",
        93: "Ultra Shotgun",
        94: "Ultra Crossbow",
        95: "Ultra Nader",
        96: "Plasma Minigun",
        97: "Devastator",
        98: "Gold Plasma Gun",
        99: "Gold Slugger",
        100: "Gold Splinter Gun",
        101: "Gold Screwdriver",
        102: "Gold Bazooka",
        103: "Gold A Rifle",
        104: "Super Disc Gun",
        105: "Heavy Auto Crossbow",
        106: "Heavy A Rifle",
        107: "Blood Cannon",
        108: "Dog Spin Attack",
        109: "Scrap Boss Missile Idle",
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
        120: "Frog Blaster",
        121: "Black Sword",
        122: "Gold Nuke Launcher",
        123: "Gold Disc Gun",
        124: "Heavy Nader",
        125: "Gun Gun",
        201: "Gold Frog Blaster",
    },
    "crownChoices": {
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
        13: "Crown of Protection",
    }
}

CROWN_SPRITES = "sprCrown([0-9]+)Walk.png"
CHARACTER_SPRITES = "sprMutant([0-9]+B?)Idle.png"
RESOLUTION_MATCH = re.compile(" ([0-9]+x[0-9]+)\\+0\\+0 ")
TEMPDIR = "temp"


class Entry(object):
    def __init__(self, key=None, type=None, name=None, source=None,
                 fixed=False, destination=None):
        self.type = type
        self.key = key
        self.source = source
        self.destination = destination
        self.name = name
        self.fixed = fixed

    def set_name(self, name):
        self.name = name

    def get_destination(self):
        if not self.source:
            return ""

        if self.destination:
            return self.destination

        if self.type == "characters":
            filename = "sprMutant{}Idle.gif".format(self.name)
        else:
            filename = "{}.gif".format(self.key)

        return "{}/{}".format(FOLDERS[self.type], filename)

    def get_filename_guess(self):
        return self.name.replace(" ", "").lower()

    def found(self, source):
        source = str(source)

        if self.fixed:
            return

        if not self.source:
            self.source = source
        elif len(source) <= len(self.source):
            if "Idle" in self.source or "Idle" not in self.source:
                self.source = source

    def get_dimensions(self):
        output = subprocess.check_output(["identify", self.source])
        width, height = RESOLUTION_MATCH.search(output).group(1).split("x")
        width, height = int(width), int(height)

        frames = max(width / height, 1)
        if "sprMutant6BIdle" in self.source:
            frames = 14
        elif "sprTechnoMancer" in self.source:
            frames = 10
        elif "sprFrozenCarThrown" in self.source:
            frames = 6
        elif "sprBloodCannon" in self.source:
            frames = 7
        elif "sprBloodHammer" in self.source:
            frames = 7
        elif "sprBloodNader" in self.source:
            frames = 7
        elif "sprRatWalk" in self.source:
            frames = 6
        elif "sprFastRatWalk" in self.source:
            frames = 6

        if self.type == "weaponChoices":
            if self.key in (25, 46):
                frames = 6
            elif self.key == 109:
                frames = 4
            else:
                frames = 7

        return width, height, frames

    def make_gif(self, only_missing=False, loop=0):
        dst = self.get_destination()

        if only_missing and os.path.isfile(dst):
            return

        width, height, frames = self.get_dimensions()

        os.system(" ".join([
            "convert",
            self.source,
            "-crop", "{}x{}".format(width / frames, height),
            "+repage",
            "+adjoin",
            "{}/sprite-%02d.png".format(TEMPDIR)
        ]))

        # Copy first frame as last frame if we're not supposed to loop
        if loop > 0:
            last = 0
            for file in os.listdir(TEMPDIR):
                last = max(last, int(file[-7:][2]))

            new_last_file = "sprite-{:0>2d}.png".format(last + 1)

            tmp = Path(TEMPDIR)

            shutil.copy(str(tmp / "sprite-00.png"), str(tmp / new_last_file))

        os.system(" ".join([
            "convert",
            "-delay", "8",
            "-loop", str(loop),
            "-dispose", "previous",
            "{}/sprite-*.png".format(TEMPDIR),
            "-delete", "30-1",
            dst
        ]))

        clear_temp()

        print("{:>38} -> {}".format(os.path.basename(self.source), dst))

    def __str__(self):
        return '<Entry type={type}, source={source}, name={name}>'.format(
            type=self.type,
            source=self.source,
            name=self.name
        )

    def __repr__(self):
        return str(self)


def usage():
    print("Usage:")
    print("  {} /path/to/asset/source".format(sys.argv[0]))


def setup_folders():
    paths = []
    for key in FOLDERS:
        paths.append(FOLDERS[key])
    paths.append(TEMPDIR)

    for path in paths:
        p = Path(path)
        if not p.exists():
            p.mkdir(0o750, parents=True)
            print("Created {}".format(p.resolve()))


def clear_temp():
    for filename in os.listdir(TEMPDIR):
        (Path(TEMPDIR) / filename).unlink()


def find_resources(src):
    files = {
        "characters": {},
        "causesOfDeath": {
            0: Entry(
                key=0,
                type="causesOfDeath",
                source=str(Path(src) / "sprBanditIdle.png"),
                fixed=True),
        },
        "weaponChoices": {},
        "crownChoices": {},
    }

    # Initialize Character entries
    for char in range(1, 17):
        bparts = ("", "B") if char < 13 else ("",)
        for b in bparts:
            key = "{}{}".format(char, b)
            files["characters"][key] = Entry(
                key,
                "characters",
                key
            )

    # Initialize other entries
    for type in DATA:
        for key in DATA[type]:
            name = DATA[type][key]
            if key in files[type]:
                files[type][key].set_name(name)
            else:
                files[type][key] = Entry(key, type, name)

    # Regex to match some sprites
    character_match = re.compile(CHARACTER_SPRITES)
    crown_match = re.compile(CROWN_SPRITES)

    skip = [
        "sprCarpet.png",
        "sprPopoFreakGun.png"
    ]

    for filename in os.listdir(src):
        if filename[:3] != "spr" or filename in skip:
            continue

        path = Path(src) / filename
        resolved = path.resolve()

        match = character_match.match(filename)
        if match:
            key = match.group(1)
            files["characters"][key].found(resolved)
            continue

        match = crown_match.match(filename)
        if match:
            key = int(match.group(1))
            files["crownChoices"][key].found(resolved)
            continue

        # Not a crown or a character
        lower = filename.lower()
        for type in DATA:
            if type == "crownChoices":
                continue
            if type != "weaponChoices" and "gun" in lower:
                continue

            for key in DATA[type]:
                entry = files[type][key]
                if entry.fixed:
                    continue

                name = entry.get_filename_guess()
                if name in lower:
                    entry.found(resolved)

    return files


def generate_resources(sources, only_missing):
    for type in sources:
        for key in sources[type]:
            entry = sources[type][key]
            if not entry.source:
                continue

            if type == "weaponChoices":
                entry.make_gif(only_missing, loop=1)
            else:
                entry.make_gif(only_missing)


def generate_tests(src):
    tmpl = "<tr><td><i style='background-image:url({filename});'></td><td>{name}</td></tr>"

    for type in FOLDERS:
        folder = FOLDERS[type]

        if type not in src:
            continue

        rows = []

        for key in src[type]:
            entry = src[type][key]
            name = entry.name
            filename = os.path.basename(entry.get_destination())

            if entry.source is None:
                name += " (MISSING)"

            rows.append(tmpl.format(filename=filename, name=name))

        with open("{}/test.html".format(folder), "w") as f:
            f.write(TEST_TEMPLATE.format(rows="\n".join(rows)))


if __name__ == "__main__":
    ap = ArgumentParser()
    ap.set_defaults(only_missing=False)
    ap.add_argument("source")
    ap.add_argument("--only_missing", action="store_true")

    args = ap.parse_args()

    setup_folders()
    resources = find_resources(args.source)

    generate_resources(resources, args.only_missing)

    generate_tests(resources)

    print("")
    print("Missing enemies:")
    missing = []
    for id in resources["causesOfDeath"]:
        entry = resources["causesOfDeath"][id]
        if not entry.source:
            print(entry.name)
            missing.append(entry.key)

    pprint(missing)

    print("")
    print("Missing weapons:")
    missing = []
    for id in resources["weaponChoices"]:
        entry = resources["weaponChoices"][id]
        if not entry.source:
            print(entry.name)
            missing.append(entry.key)

    pprint(missing)
