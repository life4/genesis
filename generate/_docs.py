import attr
from jinja2 import Environment, PackageLoader
from pathlib import Path
from typing import Sequence

from ._file import File
from ._types import Type
from ._exclude import is_excluded


env = Environment(
    loader=PackageLoader('generate', 'templates'),
)


@attr.s()
class Docs:
    code_file = attr.ib(type=File)
    test_file = attr.ib(type=File)
    types = attr.ib(type=Sequence[Type])

    def render_root(self) -> str:
        template = env.get_template('root.md.j2')
        return template.render(
            structs=self.code_file.structs,
        )

    def render_struct(self, struct) -> str:
        template = env.get_template('struct.md.j2')
        return template.render(
            struct=struct,
            funcs=self.code_file.functions,
        )

    def render_func(self, func) -> str:
        test = None
        for t in self.test_file.functions:
            if t.name == func.name and t.struct == func.struct:
                test = t
                break

        template = env.get_template('func.md.j2')
        rendered = template.render(
            func=func,
            test=test,
            types=self.types,
            structs=self.code_file.structs,
            is_excluded=is_excluded,
            sorted=sorted,
        )
        while '\n\n\n' in rendered:
            rendered = rendered.replace('\n\n\n', '\n\n')
        return rendered

    def render(self, path: Path) -> None:
        path.mkdir(parents=True, exist_ok=True)
        (path / 'README.md').write_text(self.render_root())

        for struct in self.code_file.structs:
            subpath = path / struct.name.lower() / 'README.md'
            subpath.parent.mkdir(exist_ok=True)
            subpath.write_text(self.render_struct(struct=struct))

            for func in self.code_file.functions:
                if not func.public:
                    continue
                if func.struct != struct.name:
                    continue
                func_path = (subpath.parent / func.name.lower()).with_suffix('.md')
                func_path.write_text(self.render_func(func=func))
