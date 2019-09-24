from typing import Dict, List, Set

import attr
import parse

from ._types import Type


_t = '// {docs}\ntype {name:w} struct {{\n{body}\n}}'
parser = parse.compile(_t)
TEMPLATE = _t.replace(':w}', '}')


@attr.s()
class Struct:
    name = attr.ib()
    body = attr.ib()
    docs = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Struct']:
        structs = []
        for match in parser.findall(text):
            structs.append(cls(**match.named))
        return structs

    @property
    def generics(self) -> Set[str]:
        result = set()
        for name in ('G', 'T'):
            if name in self.body:
                result.add(name)
        return result

    def render(self, types: Dict[str, Type]) -> str:
        struct_name = self.name
        body = self.body

        # substitute types instead of generics
        for generic, t in types.items():
            struct_name += t.title
            body = body.replace(generic, t.name)

        # insert modified values into template
        return TEMPLATE.format(
            docs=self.docs,
            name=struct_name,
            body=body,
        )
