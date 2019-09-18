from pathlib import Path

from ._file import File


def entrypoint():
    files = []
    for path in Path('implementation').iterdir():
        files.append(File.from_path(path))
    file = File.merge(*files)
    print(file.render())
