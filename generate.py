import subprocess
from itertools import product
from pathlib import Path
from shutil import rmtree
from typing import Sequence


TYPES = (
    'bool', 'byte', 'string',
    'float32', 'float64',
    'int', 'int8', 'int16', 'int32', 'int64',
    'uint', 'uint8', 'uint16', 'uint32', 'uint64',
    'interface{}',
)
EXCLUDE = ('MaxBool', 'MinBool', 'Max', 'Min')
GENERICS = ('T', 'G')


def _make_suffix(types: Sequence[str]) -> str:
    if len(types) == 1 and types[0] == 'interface{}':
        return ''
    suffixes = []
    for t in types:
        if t == 'interface{}':
            t = 'Interface'
        suffixes.append(t.title())
    return ''.join(suffixes)


def generate(in_path: Path, out_path: Path):
    out_path.mkdir(exist_ok=True)

    # read generic types declaration
    types_content = (in_path / 'types.go').read_text() + '\n\n'
    types_content = '\n'.join(types_content.split('\n')[1:])

    for module in in_path.iterdir():
        print(module.stem)
        if module.stem == 'types':
            continue
        tmp_path = out_path / 'tmp.go'
        with tmp_path.open('w') as stream:
            lines = module.read_text().split('\n')
            # write package name
            stream.write(lines[0])
            # write generic types declaration
            stream.write(types_content)
            # write functions definition
            stream.write('\n'.join(lines[1:]))

        # generate function definition for every type
        generics_count = int(module.stem[-1])
        for types in product(TYPES, repeat=generics_count):
            print('   ', ', '.join(types))
            output_name = module.stem + ''.join(types).replace('{', '').replace('}', '')
            cmd = [
                'go_generics',
                '-i', str(tmp_path),
                '-o', out_path.joinpath(output_name).with_suffix('.go'),
                '-p', 'genesis',
                '-suffix', _make_suffix(types),
            ]
            for g, t in zip(GENERICS, types):
                cmd.extend(['-t', '{}={}'.format(g, t)])
            subprocess.call(cmd)

        tmp_path.unlink()


def merge(in_path: Path, out_path: Path):
    cmd = ['go_merge', '-o', out_path]
    cmd.extend([str(path) for path in in_path.iterdir()])
    subprocess.call(cmd)


def _is_excluded_type(line: str) -> bool:
    # exclude generic type declaration
    if line.startswith('type') and line.rstrip().endswith('int8'):
        return True
    return False


def _is_excluded_func(line: str) -> bool:
    # exclude blacklisted functions
    for name in EXCLUDE:
        prefix = 'func {}('.format(name)
        if line.startswith(prefix):
            return True
    return False


def clean(path: Path):
    lines = path.read_text().split('\n')
    cleaned = []
    inside_of_bad = False
    for line in lines:
        excluded_type = _is_excluded_type(line)
        excluded_func = _is_excluded_func(line)
        if excluded_type or excluded_func:
            # drop documentation for excluded function
            if cleaned[-1].startswith('//'):
                cleaned.pop(-1)
            # drop empty lines before
            while not cleaned[-1]:
                cleaned.pop(-1)
            if excluded_func:
                inside_of_bad = True
            continue

        if line and line[0] == '}' and inside_of_bad:
            inside_of_bad = False
            continue

        if not inside_of_bad:
            cleaned.append(line)

    path.write_text('\n'.join(cleaned))


if __name__ == '__main__':
    generate(in_path=Path('implementation'), out_path=Path('generated'))
    merge(in_path=Path('generated'), out_path=Path('generated.go'))
    rmtree('generated')
    clean(Path('generated.go'))
