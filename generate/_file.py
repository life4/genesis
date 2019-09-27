import re
from functools import partial
from pathlib import Path
from typing import Iterable, List, Union, Set

import attr

from ._error import Error
from ._exclude import is_excluded
from ._function import Function
from ._struct import Struct
from ._test import Test
from ._types import Type


REX_PACKAGE = re.compile(r'package (\w+)')


@attr.s()
class File:
    package = attr.ib(type=str)
    imports = attr.ib(type=Set[str])
    functions = attr.ib(type=List[Union[Function, Test]])
    structs = attr.ib(type=List[Struct])
    errors = attr.ib(type=List[Error])

    @classmethod
    def from_text(cls, text: str, test: bool = False, example: bool = False) -> 'File':
        if example:
            factory = partial(Test.from_text, example=True)
        elif test:
            factory = Test.from_text
        else:
            factory = Function.from_text

        return cls(
            package=REX_PACKAGE.search(text).groups()[0],
            imports=cls._get_imports(text),
            functions=factory(text),
            structs=Struct.from_text(text),
            errors=Error.from_text(text),
        )

    @classmethod
    def from_path(cls, path: Path, example: bool = False) -> 'File':
        test = path.stem.endswith('_test')
        return cls.from_text(path.read_text(), test=test, example=example)

    @staticmethod
    def _get_imports(text: str) -> Set[str]:
        _before, sep, after = text.partition('import (')
        if not sep:
            return set()
        imports = after.split(')')[0]
        return set(imports.split('\n'))

    @staticmethod
    def merge(*files) -> 'File':
        return sum(files[1:], files[0])

    def __add__(self, other):
        if self.package != other.package:
            raise ValueError('merging different packages')
        return type(self)(
            package=self.package,
            imports=self.imports | other.imports,
            functions=self.functions + other.functions,
            structs=self.structs + other.structs,
            errors=self.errors + other.errors,
        )

    def render(self, types: Iterable[Type]) -> str:
        result = 'package {package}'.format(package=self.package)
        if self.imports:
            result += '\n\nimport ({})'.format('\n'.join(sorted(self.imports)))

        for error in self.errors:
            result += '\n' + error.source

        for t in types:
            # render structs for type
            for struct in self.structs:
                result += '\n\n' + struct.render({'T': t})

            for func in self.functions:
                if is_excluded(func=func, t=t):
                    continue
                if 'G' in func.generics:
                    # render function with additional generics
                    for g in types:
                        if is_excluded(func=func, t=t, g=g):
                            continue
                        result += '\n\n' + func.render({'T': t, 'G': g})
                else:
                    result += '\n\n' + func.render({'T': t})

        return result + '\n'
