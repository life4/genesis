import subprocess
from pathlib import Path
from shutil import rmtree


TYPES = (
    'bool', 'byte', 'string',
    'float32', 'float64',
    'int', 'int8', 'int16', 'int32', 'int64',
    'uint', 'uint8', 'uint16', 'uint32', 'uint64',
)
EXCLUDE = ('MaxBool', 'MinBool')


def generate():
    Path('generated').mkdir(exist_ok=True)
    for t in TYPES:
        print('#', t)
        subprocess.call([
            'go_generics',
            '-i', 'implementation/functional.go',
            '-o', 'generated/{}.go'.format(t),
            '-p', 'genesis',
            '-suffix', t.title(),
            '-t', 'T={}'.format(t),
        ])


def merge():
    cmd = ['go_merge', '-o', 'generated.go']
    cmd.extend(['generated/{}.go'.format(t) for t in TYPES])
    subprocess.call(cmd)


def clean(path: Path):
    lines = path.read_text().split('\n')
    cleaned = []
    inside_of_bad = False
    for line in lines:
        if line[5:].startswith(EXCLUDE):
            # drop documentation for excluded function
            if cleaned[-1].startswith('//'):
                cleaned.pop(-1)
            # drop empty lines before
            while not cleaned[-1]:
                cleaned.pop(-1)
            inside_of_bad = True
            continue

        if line and line[0] == '}' and inside_of_bad:
            inside_of_bad = False
            continue

        if not inside_of_bad:
            cleaned.append(line)

    path.write_text('\n'.join(cleaned))


if __name__ == '__main__':
    generate()
    merge()
    rmtree('generated')
    clean(Path('generated.go'))
