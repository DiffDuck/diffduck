def example_amend_commit_hook(path, type, sha) -> int:
    print("This is an example amend commit hook.")
    print(path, type, sha)
    return 0



def example_new_commit_hook(path) -> int:
    print("This is an example new commit hook.")
    print(path)
    return 0
