from typing import List

import attr
import parse


_t = '\nvar {name:w} = errors.New("{message}")\n'
parser = parse.compile(_t)
TEMPLATE = _t.replace(':w}', '}')


@attr.s()
class Error:
    name = attr.ib()
    message = attr.ib()

    @classmethod
    def from_text(cls, text: str) -> List['Error']:
        structs = []
        for match in parser.findall(text):
            structs.append(cls(**match.named))
        return structs

    @property
    def source(self) -> str:
        return TEMPLATE.format(
            name=self.name,
            message=self.message,
        )
