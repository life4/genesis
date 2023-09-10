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
        deprecated = False
        for line in content.splitlines():
            if not line.startswith('func '):
                deprecated = line.startswith('// DEPRECATED')
                continue
            line = line.removeprefix('func ')
            fname = line.split('[')[0].split('(')[0]
            if deprecated:
                continue
            if not fname[0].isupper():
                continue
            yield fname


def get_tests(pkg: str) -> Iterator[str]:
    for fpath in Path(pkg).iterdir():
        if not fpath.name.endswith('_test.go'):
            continue
        content = fpath.read_text()
        for line in content.splitlines():
            if not line.startswith('func Test'):
                continue
            line = line.removeprefix('func Test')
            fname = line.split('(')[0]
            assert fname[0].isupper()
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
    """Every function must have an example.
    """
    funcs = set(get_funcs(pkg))
    examples = set(get_examples(pkg))
    assert funcs == examples


@pytest.mark.parametrize('pkg', [
    'slices',
    'maps',
    'lambdas',
    # channels need tests for context-aware versions
    # 'channels',
])
def test_all_have_tests(pkg: str) -> None:
    """Every function must have unit tests.
    """
    funcs = set(get_funcs(pkg))
    tests = set(get_tests(pkg))
    assert funcs
    assert not funcs - tests


@pytest.mark.parametrize('func', get_funcs('slices'))
def test_slices_func_linked_in_docs(func: str) -> None:
    """Every function in the slices package must be listed in the package docs.
    """
    docs = Path('slices', 'doc.go').read_text()
    assert f'//   - [{func}](' in docs