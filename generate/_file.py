import re
from itertools import product
from pathlib import Path
from typing import Iterable, List

import attr

from ._function import Function
from ._types import Type, TYPES


REX_PACKAGE = re.compile(r'package (\w+)')


@attr.s()
class File:
    package = attr.ib(type=str)
    functions = attr.ib(type=List[Function])

    @classmethod
    def from_text(cls, text: str) -> 'File':
        return cls(
            package=REX_PACKAGE.search(text).groups()[0],
            functions=Function.from_text(text),
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
        )

    def render(self, types: Iterable[Type] = None) -> str:
        if types is None:
            types = TYPES
        result = 'package {package}'.format(package=self.package)
        for t in types:
            for func in self.functions:
                if 'G' in func.generics:
                    continue
                result += '\n\n' + func.render({'T': t})

        for t, g in product(types, repeat=2):
            for func in self.functions:
                if 'G' not in func.generics:
                    continue
                result += '\n\n' + func.render({'T': t, 'G': g})

        return result + '\n'
