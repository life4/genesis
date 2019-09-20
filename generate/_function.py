from typing import Dict, List, Set

import attr
import parse

from ._types import Type


_t = 'func ({pointer}{struct:w}) {name:w}{signature} {{\n{body}\n}}'
parser = parse.compile(_t)
TEMPLATE = _t.replace(':w}', '}')


@attr.s()
class Function:
    pointer = attr.ib()
    struct = attr.ib()
    name = attr.ib()
    signature = attr.ib()
    body = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Function']:
        functions = []
        for match in parser.findall(text):
            functions.append(cls(**match.named))
        return functions

    @property
    def generics(self) -> Set[str]:
        result = set()
        for name in ('G', 'T'):
            if name in self.signature:
                result.add(name)
        return result

    def render(self, types: Dict[str, Type]) -> str:
        function_name = self.name
        signature = self.signature
        body = self.body

        # add the second generic type into func title if specified
        if 'G' in types:
            function_name += types['G'].title
        # substitute types instead of generics
        for generic, t in types.items():
            signature = signature.replace(generic, t.name)
            body = body.replace(generic, t.name)
            body = body.replace(t.name + 'roup', 'Group')  # restore WaitGroup

        # insert modified values into template
        return TEMPLATE.format(
            pointer=self.pointer,
            struct=self.struct + types['T'].title,
            name=function_name,
            signature=signature,
            body=body,
        )
