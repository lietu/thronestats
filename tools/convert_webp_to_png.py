"""

"""

import os
import sys
from pathlib import Path


if __name__ == "__main__":
    path = sys.argv[1]

    for file in os.listdir(path):
        if file[-5:] == ".webp":
            src = str((Path(path) / file).resolve())
            dst = src.replace(".webp", ".png")
            os.system("convert {} {}".format(src, dst))
            print("Converted {} to {}".format(src, dst))