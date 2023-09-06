from __future__ import annotations
from pathlib import Path
from typing import Iterator
import pytest


def get_funcs(pkg: str) -> Iterator[str]:
    for fpath in Path(pkg).iterdir():
        if fpath.suffix != '.go':
            continue
        if fpath.stem.endswith('_test'):
            continue
        content = fpath.read_text()
        for line in content.splitlines():
            if not line.startswith('func '):
                continue
            line = line.removeprefix('func ')
            fname = line.split('[')[0]
            if not fname[0].isupper():
                continue
            yield fname


def get_examples(pkg: str) -> Iterator[str]:
    fpath = Path(pkg, 'examples_test.go')
    assert fpath.is_file()
    content = fpath.read_text()
    for line in content.splitlines():
        if not line.startswith('func Example'):
            continue
        line = line.removeprefix('func Example')
        fname = line.split('(')[0]
        yield fname


@pytest.mark.parametrize('pkg', ['slices'])
def test_all_have_examples(pkg: str) -> None:
    funcs = set(get_funcs(pkg))
    examples = set(get_examples(pkg))
    assert funcs == examples
