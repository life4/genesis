import attr
from jinja2 import Environment, PackageLoader
from pathlib import Path

from ._file import File


env = Environment(
    loader=PackageLoader('generate', 'templates'),
)


@attr.s()
class Docs:
    file = attr.ib(type=File)

    def render_root(self) -> str:
        template = env.get_template('root.md.j2')
        return template.render(
            structs=self.file.structs,
        )

    def render_struct(self, struct) -> str:
        template = env.get_template('struct.md.j2')
        return template.render(
            struct=struct,
            funcs=self.file.functions,
        )

    def render(self, path: Path) -> None:
        path.mkdir(parents=True, exist_ok=True)
        (path / 'README.md').write_text(self.render_root())
        for struct in self.file.structs:
            subpath = path / struct.name.lower() / 'README.md'
            subpath.parent.mkdir(exist_ok=True)
            subpath.write_text(self.render_struct(struct=struct))
