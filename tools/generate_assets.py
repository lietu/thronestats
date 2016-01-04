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
        "characters": []
    }

    character_match = re.compile(CHARACTER_SPRITES)

    for filename in os.listdir(src):
        if character_match.match(filename):
            path = Path(src) / filename
            dst = str((Path(FOLDERS["characters"]) / filename)).replace(".png",
                                                                        ".gif")
            fullpath = str(path.resolve())

            files["characters"].append(fullpath)
            make_gif(fullpath, dst)

    return files


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
