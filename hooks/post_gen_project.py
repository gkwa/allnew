import pathlib
import stat


def set_script_permissions():
    for path in pathlib.Path(".").glob('**/*.sh'):
        path.chmod(path.stat().st_mode | stat.S_IEXEC)


if __name__ == "__main__":
    set_script_permissions()
