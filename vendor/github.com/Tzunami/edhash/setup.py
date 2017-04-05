#!/usr/bin/env python
import os
from distutils.core import setup, Extension
sources = [
    'src/python/core.c',
    'src/libedhash/io.c',
    'src/libedhash/internal.c',
    'src/libedhash/sha3.c']
if os.name == 'nt':
    sources += [
        'src/libedhash/util_win32.c',
        'src/libedhash/io_win32.c',
        'src/libedhash/mmap_win32.c',
    ]
else:
    sources += [
        'src/libedhash/io_posix.c'
    ]
depends = [
    'src/libedhash/edhash.h',
    'src/libedhash/compiler.h',
    'src/libedhash/data_sizes.h',
    'src/libedhash/endian.h',
    'src/libedhash/edhash.h',
    'src/libedhash/io.h',
    'src/libedhash/fnv.h',
    'src/libedhash/internal.h',
    'src/libedhash/sha3.h',
    'src/libedhash/util.h',
]
pyedhash = Extension('pyedhash',
                     sources=sources,
                     depends=depends,
                     extra_compile_args=["-Isrc/", "-std=gnu99", "-Wall"])

setup(
    name='pyedhash',
    author="Matthew Wampler-Doty",
    author_email="matthew.wampler.doty@gmail.com",
    license='GPL',
    version='0.1.23',
    url='https://github.com/Tzunami/ethash',
    download_url='https://github.com/Tzunami/ethash/tarball/v23',
    description=('Python wrappers for edhash, the ethereum proof of work'
                 'hashing function'),
    ext_modules=[pyedhash],
)
