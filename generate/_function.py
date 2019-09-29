from typing import Dict, List, Set

import attr
import parse

from ._types import Type, replace_type


_t = 'func ({pointer}{struct:w}) {name:w}{signature} {{\n{body}\n}}'
parser = parse.compile(_t)
TEMPLATE = '{docs}\n' + _t.replace(':w}', '}')


@attr.s()
class Function:
    pointer = attr.ib()
    struct = attr.ib()
    name = attr.ib()
    signature = attr.ib()
    body = attr.ib()
    docs = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Function']:
        functions = []
        for match in parser.findall(text):
            start = match.spans['pointer'][0] - 7
            docs = text[:start].rsplit('\n\n', maxsplit=1)[-1]
            functions.append(cls(docs=docs, **match.named))
        return functions

    @property
    def generics(self) -> Set[str]:
        result = {'T'}
        for name in ('G', 'T'):
            if name in self.signature:
                result.add(name)
        return result

    @property
    def public(self) -> bool:
        return self.name[0].isupper()

    @property
    def source(self) -> str:
        return TEMPLATE.format(
            pointer=self.pointer,
            struct=self.struct,
            name=self.name,
            signature=self.signature,
            body=self.body,
            docs=self.docs,
        )

    @property
    def clean_docs(self) -> str:
        return self.docs.replace('// ', '').replace('\n', ' ')

    def render(self, types: Dict[str, Type]) -> str:
        function_name = self.name
        signature = self.signature
        body = self.body

        # add the second generic type into func title if specified
        if 'G' in types:
            function_name += types['G'].title
        # substitute types instead of generics
        for generic, t in types.items():
            body = replace_type(text=body, type_from=generic, type_to=t.name)
            signature = replace_type(text=signature, type_from=generic, type_to=t.name)

        # insert modified values into template
        return TEMPLATE.format(
            pointer=self.pointer,
            struct=self.struct + types['T'].title,
            name=function_name,
            signature=signature,
            body=body,
            docs=self.docs,
        )
