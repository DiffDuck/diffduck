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
    print("Runing diffduck commit hook")
    p = subprocess.run(
        args=["./diffduck " + path],
        shell=True,
        capture_output=True,
    )
    if len(p.stdout) > 0:
        print(p.stdout.decode("utf-8"))
    if len(p.stderr) > 0:
        print(p.stderr.decode("utf-8"))
    return p.returncode
