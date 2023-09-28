import subprocess

def example_amend_commit_hook(path) -> int:
    print("This is an example amend commit hook.")
    print(path)
    return 0



def example_new_commit_hook(path) -> int:
    print("This is an example new commit hook.")
    print(path)
    return 0

def diffduck_commit_hook(path) -> int:
    print("Running diffduck commit hook")
    p = subprocess.Popen(
        args=["go run ./cmd/diffduck " + path],
        shell=True,
    )
    return p.wait()
