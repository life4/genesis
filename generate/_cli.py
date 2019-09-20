import sys
from argparse import ArgumentParser
from pathlib import Path
from typing import List

from ._file import File


def get_parser() -> ArgumentParser:
    parser = ArgumentParser()
    parser.add_argument('--input', default='implementation')
    parser.add_argument('--output', default='generated.go')
    parser.add_argument('--tests', default='generated_test.go')
    parser.add_argument('--package', default='genesis')
    return parser


def get_paths(input_path: Path):
    if input_path.is_file():
        yield input_path
    else:
        yield from input_path.iterdir()


def render(paths, package: str) -> str:
    files = []
    for path in paths:
        files.append(File.from_path(path))
    file = File.merge(*files)
    file.package = package
    return file.render()


def entrypoint(argv: List[str] = None):
    if argv is None:
        argv = sys.argv[1:]
    parser = get_parser()
    args = parser.parse_args(argv)

    paths = list(get_paths(Path(args.input)))

    # generate main content
    content = render(
        paths=[p for p in paths if not p.stem.endswith('_test')],
        package=args.package,
    )
    Path(args.output).write_text(content)

    # generate tests
    content = render(
        paths=[p for p in paths if p.stem.endswith('_test')],
        package=args.package,
    )
    Path(args.tests).write_text(content)
