from __future__ import annotations
from pathlib import Path


def count_funcs(pkg: str) -> int:
    count = 0
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
            count += 1
    return count


total = 0
pkgs = ('channels', 'lambdas', 'maps', 'sets', 'slices')
for pkg in pkgs:
    count = count_funcs(pkg)
    print(f'{pkg}: {count}')
    total += count
print(f'total: {total}')
