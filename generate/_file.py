import re
from pathlib import Path
from typing import Iterable, List

import attr

from ._function import Function
from ._struct import Struct
from ._types import Type, TYPES


REX_PACKAGE = re.compile(r'package (\w+)')


@attr.s()
class File:
    package = attr.ib(type=str)
    functions = attr.ib(type=List[Function])
    structs = attr.ib(type=List[Struct])

    @classmethod
    def from_text(cls, text: str) -> 'File':
        return cls(
            package=REX_PACKAGE.search(text).groups()[0],
            functions=Function.from_text(text),
            structs=Struct.from_text(text),
        )

    @classmethod
    def from_path(cls, path: Path) -> 'File':
        return cls.from_text(path.read_text())

    @staticmethod
    def merge(*files) -> 'File':
        return sum(files[1:], files[0])

    def __add__(self, other):
        if self.package != other.package:
            raise ValueError('merging different packages')
        return type(self)(
            package=self.package,
            functions=self.functions + other.functions,
            structs=self.structs + other.structs,
        )

    def render(self, types: Iterable[Type] = None) -> str:
        if types is None:
            types = TYPES
        result = 'package {package}'.format(package=self.package)

        for t in types:
            # render structs for type
            for struct in self.structs:
                result += '\n\n' + struct.render({'T': t})

            for func in self.functions:
                if 'G' in func.generics:
                    # render function with additional generics
                    for g in types:
                        result += '\n\n' + func.render({'T': t, 'G': g})
                else:
                    result += '\n\n' + func.render({'T': t})

        return result + '\n'
