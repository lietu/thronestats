"""
Generate the graphics assets for ../www/img/ -folders.
"""
import time
import re
import os
import sys
import subprocess
import shutil
from pathlib import Path

FOLDERS = {
    "characters": "../www/img/characters",
    "causesOfDeath": "../www/img/causesOfDeath",
    "crownChoices": "../www/img/crownChoices",
    "deathsByLevel": "../www/img/deathsByLevel",
    "mutationChoices": "../www/img/mutationChoices",
    "weaponChoices": "../www/img/weaponChoices",
}

ENEMIES = {
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
}

CHARACTER_SPRITES = "sprMutant([0-9]+B?)Idle.png"
RESOLUTION_MATCH = re.compile(" ([0-9]+x[0-9]+)\\+0\\+0 ")
TEMPDIR = "temp"


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
        "characters": [],
        "causesOfDeath": {}
    }

    for enemy in ENEMIES:
        files["causesOfDeath"][enemy] = ENEMIES[enemy]

    character_match = re.compile(CHARACTER_SPRITES)

    for filename in os.listdir(src):
        path = Path(src) / filename
        if character_match.match(filename):
            files["characters"].append(str(path.resolve()))
        else:
            for key in ENEMIES:
                name = ENEMIES[key].replace(" ", "")
                if name in filename and "Idle" in filename:
                    print("{}: {}".format(name, filename))
                    files["causesOfDeath"][key] = path

    return files


def generate_resources(sources):
    for src in sources["characters"]:
        basename = os.path.basename(src)
        dst = Path(FOLDERS["characters"]) / basename.replace(".png", ".gif")
        make_gif(src, dst)


def get_dimensions(src):
    output = subprocess.check_output(["identify", src])
    width, height = RESOLUTION_MATCH.search(output).group(1).split("x")
    width, height = int(width), int(height)

    frames = width / height
    if "sprMutant6BIdle" in src:
        frames = 14

    return width, height, frames


def make_gif(src, dst):
    width, height, frames = get_dimensions(src)

    os.system(" ".join([
        "convert",
        src,
        "-crop", "{}x{}".format(width / frames, height),
        "+repage",
        "+adjoin",
        "{}/sprite-%02d.png".format(TEMPDIR)
    ]))

    os.system(" ".join([
        "convert",
        "-delay", "8",
        "-loop", "0",
        "-dispose", "previous",
        "{}/sprite-*.png".format(TEMPDIR),
        "-delete", "30-1",
        dst
    ]))

    clear_temp()

    print("{} -> {}".format(src, dst))


if __name__ == "__main__":
    if len(sys.argv) != 2:
        usage()
        sys.exit(1)

    setup_folders()
    resources = find_resources(sys.argv[1])

    from pprint import pprint

    pprint(resources)
