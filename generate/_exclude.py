from ._types import Type

EXCLUDED = {
    # 'Count': ['bool', 'byte', 'string', 'interface{}'],
    # 'Exponential': ['bool', 'byte', 'string', 'interface{}'],
    # 'Rasnge': ['bool', 'byte', 'string', 'interface{}'],
    'Slice': {
        'Max': ('bool', 'interface{}'),
        'Min': ('bool', 'interface{}'),
    },
}


def is_excluded(struct: str, func: str, t: Type) -> bool:
    return t.name in EXCLUDED.get(struct, {}).get(func, ())
