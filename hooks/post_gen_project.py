import pathlib
import stat
import subprocess
import sys


def run_go_mod_tidy() -> bool:
    try:
        subprocess.run(["go", "mod", "tidy"], capture_output=True, check=True)
        return True
    except Exception:
        return False


def install_git_pre_commit() -> None:
    cmd = ["pre-commit", "sample-config"]
    proc = subprocess.Popen(
        cmd,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )

    out, err = proc.communicate()
    return_code = proc.poll()
    out = out.decode(sys.stdin.encoding)
    err = err.decode(sys.stdin.encoding)

    _dir = pathlib.Path("support")
    config = _dir / ".pre-commit-config.yaml"
    config.write_text(out)

    if err:
        msg = f"pre-commit sample-config >{config} failed"
        raise ValueError(msg)


def set_script_permissions() -> None:
    for path in pathlib.Path(".").glob("**/*.sh"):
        path.chmod(path.stat().st_mode | stat.S_IEXEC)


if __name__ == "__main__":
    set_script_permissions()
    if not run_go_mod_tidy():
        print("ERROR: go mod tidy failed.")
        sys.exit(1)
    install_git_pre_commit()
