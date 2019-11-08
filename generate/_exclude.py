from ._test import Test
from ._types import Type


EXCLUDED_TESTS = (
    'bool', 'interface{}', 'byte', 'string', 'error',
    'float32', 'float64',
    'uint', 'uint8', 'uint16', 'uint32', 'uint64',
)
EXCLUDED = {
    'Sequence': {
        'Count': ['bool', 'byte', 'error', 'string', 'interface{}'],
        'Exponential': ['bool', 'byte', 'error', 'string', 'interface{}'],
        'Range': ['bool', 'byte', 'error', 'string', 'interface{}'],
    },
    'Slice': {
        'Join': ('float32', 'float64', 'bool', 'interface{}'),
        'Max': ('bool', 'error', 'interface{}'),
        'Min': ('bool', 'error', 'interface{}'),
        'Sum': ('bool', 'error', 'interface{}'),
        'Sort': ('bool', 'error', 'interface{}'),
        'Sorted': ('bool', 'error', 'interface{}'),
    },
    'Channel': {
        'Max': ('bool', 'error', 'interface{}'),
        'Min': ('bool', 'error', 'interface{}'),
        'Sum': ('bool', 'error', 'interface{}'),
    },
    'Pair': {
        'Max': ('bool', 'error', 'interface{}'),
        'Min': ('bool', 'error', 'interface{}'),
    },
}  # type: dict


def is_excluded(func, t: Type, g: Type = None) -> bool:
    if isinstance(func, Test):
        if t.name in EXCLUDED_TESTS:
            return True
        if g and g.name in EXCLUDED_TESTS:
            return True
    return t.name in EXCLUDED.get(func.struct, {}).get(func.name, ())
