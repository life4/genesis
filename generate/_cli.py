import sys
from argparse import ArgumentParser
from pathlib import Path
from typing import List

from ._file import File


def get_parser() -> ArgumentParser:
    parser = ArgumentParser()
    parser.add_argument('--input', default='implementation')
    parser.add_argument('--output', default='generated.go')
    # parser.add_argument('--tests', default='generated_test.go')
    parser.add_argument('--package', default='genesis')
    return parser


def get_paths(input_path: Path):
    if input_path.is_file():
        yield input_path
    else:
        yield from input_path.iterdir()


def entrypoint(argv: List[str] = None):
    if argv is None:
        argv = sys.argv[1:]
    parser = get_parser()
    args = parser.parse_args(argv)

    files = []
    for path in get_paths(Path(args.input)):
        files.append(File.from_path(path))
    file = File.merge(*files)
    file.package = args.package
    content = file.render()
    Path(args.output).write_text(content)
