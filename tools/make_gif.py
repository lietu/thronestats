from generate_assets import Entry
from argparse import ArgumentParser


def get_options():
    ap = ArgumentParser()
    ap.set_defaults(loop=0)
    ap.add_argument("--src", required=True)
    ap.add_argument("--dst", required=True)
    ap.add_argument("--loop", type=int)

    return ap.parse_args()


if __name__ == "__main__":
    options = get_options()
    entry = Entry(
        source=options.src,
        destination=options.dst
    )

    entry.make_gif(options.loop)