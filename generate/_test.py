from typing import Dict, List, Set

import attr
import parse

from ._types import Type


_t = 'func Test{name:w}(t *testing.T) {{\n{body}\n}}'
parser = parse.compile(_t)
TEMPLATE = _t.replace(':w}', '}')
parser_struct = parse.compile('{name:w}{{given}}.')


@attr.s()
class Test:
    name = attr.ib()
    body = attr.ib()
    struct = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Test']:
        functions = []
        for match in parser.findall(text):
            struct_match = parser_struct.search(match.named['body'])
            if not struct_match:
                msg = 'cannot detect struct name for Test{}'
                raise ValueError(msg.format(match.named['name']))
            struct = struct_match.named['name']
            functions.append(cls(struct=struct, **match.named))
        return functions

    @property
    def generics(self) -> Set[str]:
        result = set()
        for name in ('G', 'T'):
            if name in self.body:
                result.add(name)
        return result

    def render(self, types: Dict[str, Type]) -> str:
        test_name = self.name
        body = self.body
        for generic, t in types.items():
            test_name += t.title
            body = body.replace(' ' + generic, ' ' + t.name)
            body = body.replace(']' + generic, ']' + t.name)
            body = body.replace(generic + '(', t.name + '(')
            body = body.replace(generic + ',', t.name + ',')
            body = body.replace(generic + ']', t.name + ']')

            # add suffix to struct name
            if generic == 'T':
                body = body.replace(self.struct, self.struct + t.title)
            # add suffix to func name
            if generic == 'G':
                body = body.replace(self.name, self.name + t.title)

        return TEMPLATE.format(
            name=test_name,
            body=body,
        )
