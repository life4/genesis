from __future__ import annotations
from pathlib import Path
from typing import Iterator
import pytest

ALL_PACKAGES = (
    'channels',
    'lambdas',
    'maps',
    'sets',
    'slices',
)


def get_funcs_for_pkg(pkg: str) -> Iterator[str]:
    for fpath in Path(pkg).iterdir():
        if fpath.suffix != '.go':
            continue
        if fpath.stem.endswith('_test'):
            continue
        yield from get_funcs_for_file(fpath)


def get_funcs_for_file(fpath: Path) -> Iterator[str]:
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


@pytest.mark.parametrize('pkg', [
    'slices',
    'sets',
])
def test_all_have_examples(pkg: str) -> None:
    """Every function must have an example.
    """
    funcs = set(get_funcs_for_pkg(pkg))
    examples = set(get_examples(pkg))
    assert funcs == examples


# channels need tests for context-aware versions
@pytest.mark.parametrize('pkg', sorted(set(ALL_PACKAGES) - {'channels'}))
def test_all_have_tests(pkg: str) -> None:
    """Every function must have unit tests.
    """
    funcs = set(get_funcs_for_pkg(pkg))
    tests = set(get_tests(pkg))
    assert funcs
    diff = funcs - tests
    assert not diff


@pytest.mark.parametrize('pkg', ALL_PACKAGES)
def test_all_funcs_sorted(pkg: str) -> None:
    for fpath in Path(pkg).iterdir():
        funcs = list(get_funcs_for_file(fpath))
        funcs = [func.split('_')[0] for func in funcs]
        assert funcs == sorted(funcs)


@pytest.mark.parametrize('pkg', ALL_PACKAGES)
def test_all_funcs_separated_by_newline(pkg: str) -> None:
    for fpath in Path(pkg).iterdir():
        content = fpath.read_text()
        funcs = content.split('}\n')
        for func in funcs:
            assert not func.startswith('func'), 'must have an empty line'


@pytest.mark.parametrize('func', get_funcs_for_pkg('slices'))
def test_slices_func_linked_in_docs(func: str) -> None:
    """Every func in the slices package must be listed in the package docs.
    """
    docs = Path('slices', 'doc.go').read_text()
    assert f'//   - [{func}](' in docs


@pytest.mark.parametrize('func', get_funcs_for_pkg('channels'))
def test_channels_func_linked_in_docs(func: str) -> None:
    """Every func in the channels package must be listed in the package docs.
    """
    docs = Path('channels', 'doc.go').read_text()
    assert f' [{func}]' in docs
