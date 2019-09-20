import re
from typing import Dict, List, Set

import attr
import parse

from ._types import Type


_t = 'func Test{func:w}(t *testing.T) {{\n{body}\n}}'
parser = parse.compile(_t)
TEMPLATE = _t.replace(':w}', '}')
REX_NAME = re.compile(r'(?P<struct>[A-Z][a-z]+)(?P<name>[A-Z]\w+)')


@attr.s()
class Test:
    name = attr.ib()
    body = attr.ib()
    struct = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Test']:
        functions = []
        for match in parser.findall(text):
            name_match = REX_NAME.search(match.named['func'])
            if not name_match:
                msg = 'cannot detect struct name for Test{}'
                raise ValueError(msg.format(match.named['func']))
            functions.append(cls(
                name=name_match.groupdict()['name'],
                struct=name_match.groupdict()['struct'],
                body=match.named['body'],
            ))
        return functions

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
            body = body.replace(' ' + generic, ' ' + t.name)
            body = body.replace(']' + generic, ']' + t.name)
            body = body.replace(generic + '(', t.name + '(')
            body = body.replace(generic + ',', t.name + ',')
            body = body.replace(generic + ']', t.name + ']')

            # add suffix to struct name
            if generic == 'T':
                for struct in {'Channel', 'Sequence', self.struct}:
                    body = body.replace(struct, struct + t.title)
            # add suffix to func name
            if generic == 'G':
                body = body.replace(self.name, self.name + t.title)

        return TEMPLATE.format(
            func=test_name,
            body=body,
        )
