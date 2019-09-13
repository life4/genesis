import subprocess
from pathlib import Path
from shutil import rmtree


TYPES = ('int32', 'int64', 'float32', 'float64', 'bool', 'string')


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
