import subprocess
from pathlib import Path
from shutil import rmtree


TYPES = (
    'bool', 'byte', 'string',
    'float32', 'float64',
    'int', 'int8', 'int16', 'int32', 'int64',
    'uint', 'uint8', 'uint16', 'uint32', 'uint64',
)


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

print('# clean')
cmd = ['go_merge', '-o', 'generated.go']
cmd.extend(['generated/{}.go'.format(t) for t in TYPES])
subprocess.call(cmd)

rmtree('generated')
