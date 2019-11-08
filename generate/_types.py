import re


class Type:
    __slots__ = ('name', '_title')

    def __init__(self, name: str, title: str = None):
        self.name = name
        self._title = title

    @property
    def title(self) -> str:
        if self._title is not None:
            return self._title
        return self.name.title()

    def __repr__(self) -> str:
        return '{}({})'.format(type(self).__name__, self.name)


TYPES = (
    Type('bool'), Type('byte'), Type('string'), Type('rune'), Type('error'),
    Type('float32'), Type('float64'),
    Type('int'), Type('int8'), Type('int16'), Type('int32'), Type('int64'),
    Type('uint'), Type('uint8'), Type('uint16'), Type('uint32'), Type('uint64'),
    Type('interface{}', 'Interface'),
)


def replace_type(text: str, type_from: str, type_to: str) -> str:
    text = re.sub(r'(\W){}(\W)'.format(type_from), r'\1{}\2'.format(type_to), text)
    text = re.sub(r'^{}(\W)'.format(type_from), r'{}\1'.format(type_to), text)
    text = re.sub(r'(\W){}$'.format(type_from), r'\1{}'.format(type_to), text)
    return text
