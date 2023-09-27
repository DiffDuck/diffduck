#!/usr/bin/env python3

import pathlib
import sys
import typing


### 1. Import your hook functions here ###

from __init__ import (
    example_amend_commit_hook,
    example_new_commit_hook,
    diffduck_commit_hook,
)

### END HOOK IMPORTS ###


### 2. Make sure that your hooks functions have the correct signature ###

AmendCommitHook = typing.Callable[[pathlib.Path, str, str], int]
# amend_hook(path, type, sha) -> int

NewCommitHook = typing.Callable[[pathlib.Path], int]

### END HOOK SIGNATURES ###


### 3. Add your hooks to the hooks lists ###

AMEND_COMMIT_HOOKS: typing.List[AmendCommitHook] = [
    example_amend_commit_hook,
    diffduck_commit_hook,
]


NEW_COMMIT_HOOKS: typing.List[NewCommitHook] = [
    example_new_commit_hook,
    diffduck_commit_hook,
]

### END HOOKS LISTS ###


### Hook execution ###
# Hooks are executed sequentially. If any hook returns a non-zero exit code,
# the remaining hooks are not executed and the script exits with the exit code.
def _amend_commit(path) -> int:
    for hook in AMEND_COMMIT_HOOKS:
        exit_code = hook(path)
        if exit_code > 0:
            return exit_code
    return 0


def _new_commit(path):
    for hook in NEW_COMMIT_HOOKS:
        exit_code = hook(path)
        if exit_code > 0:
            return exit_code
    return 0


if __name__ == "__main__":
    COMMIT_MESSAGE_TMP_FILE_PATH = sys.argv[1]

    if len(sys.argv) == 2:
        exit_code = _new_commit(COMMIT_MESSAGE_TMP_FILE_PATH)
        sys.exit(exit_code)

    if len(sys.argv) == 4:
        COMMIT_TYPE = sys.argv[2]
        COMMIT_SHA = sys.argv[3]
        exit_code = _amend_commit(COMMIT_MESSAGE_TMP_FILE_PATH)
        sys.exit(exit_code)

### END HOOK EXECUTION ###
