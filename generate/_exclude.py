from ._test import Test
from ._types import Type


EXCLUDED_TESTS = (
    'bool', 'interface{}', 'byte', 'string',
    'float32', 'float64',
    'uint', 'uint8', 'uint16', 'uint32', 'uint64',
)
EXCLUDED = {
    'Sequence': {
        'Count': ['bool', 'byte', 'string', 'interface{}'],
        'Exponential': ['bool', 'byte', 'string', 'interface{}'],
        'Range': ['bool', 'byte', 'string', 'interface{}'],
    },
    'Slice': {
        'Max': ('bool', 'interface{}'),
        'Min': ('bool', 'interface{}'),
        'Sum': ('bool', 'interface{}'),
        'Sort': ('bool', 'interface{}'),
        'Sorted': ('bool', 'interface{}'),
    },
    'Channel': {
        'Max': ('bool', 'interface{}'),
        'Min': ('bool', 'interface{}'),
        'Sum': ('bool', 'interface{}'),
    },
}


def is_excluded(func, t: Type, g: Type = None) -> bool:
    if isinstance(func, Test):
        if t.name in EXCLUDED_TESTS:
            return True
        if g and g.name in EXCLUDED_TESTS:
            return True
    return t.name in EXCLUDED.get(func.struct, {}).get(func.name, ())
