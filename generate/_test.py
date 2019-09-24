import re
from typing import Dict, List, Set, Tuple

import attr
import parse

from ._types import Type, replace_type


_t = 'func Test{func:w}(t *testing.T) {{\n{body}\n}}'
parser = parse.compile(_t)
TEMPLATE = _t.replace(':w}', '}')
REX_NAME = re.compile(r'(?P<struct>[A-Z][a-z]+)(?P<name>[A-Z]\w+)')
KNOWN_STRUCTS = ('AsyncSlice', )


@attr.s()
class Test:
    name = attr.ib()
    body = attr.ib()
    struct = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Test']:
        functions = []
        for match in parser.findall(text):
            name, struct = cls._get_name_and_struct(match.named['func'])
            functions.append(cls(
                name=name,
                struct=struct,
                body=match.named['body'],
            ))
        return functions

    @staticmethod
    def _get_name_and_struct(text: str) -> Tuple[str, str]:
        for struct in KNOWN_STRUCTS:
            if text.startswith(struct):
                return text[len(struct):], struct

        match = REX_NAME.search(text)
        if not match:
            msg = 'cannot detect struct name for Test{}'
            raise ValueError(msg.format(text))
        g = match.groupdict()
        return g['name'], g['struct']

    @property
    def generics(self) -> Set[str]:
        result = set()
        for name in ('G', 'T'):
            if name in self.body:
                result.add(name)
        return result

    def render(self, types: Dict[str, Type]) -> str:
        test_name = self.struct + self.name
        body = self.body
        for generic, t in types.items():
            test_name += t.title
            body = replace_type(text=body, type_from=generic, type_to=t.name)

            # add suffix to struct name
            if generic == 'T':
                for struct in {'Channel', 'Sequence', 'Slice', self.struct}:
                    body = replace_type(text=body, type_from=struct, type_to=struct + t.title)
            # add suffix to func name
            if generic == 'G':
                body = replace_type(text=body, type_from=self.name, type_to=self.name + t.title)

        return TEMPLATE.format(
            func=test_name,
            body=body,
        )
