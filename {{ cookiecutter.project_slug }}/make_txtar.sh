tmp=$(mktemp -d {{ cookiecutter.project_slug }}.XXXXX)

set -e

if [ -z "${tmp+x}" ]; then
    echo "Error: \$tmp is not set."
    exit 1
fi

if [ -z "$tmp" ]; then
    echo "Error: \$tmp is set but is an empty string."
    exit 1
fi

rm -f $tmp/{{ cookiecutter.project_slug }}.tar
rm -f $tmp/filelist.txt

{
    rg --files . \
        | grep -vE '{{ cookiecutter.project_slug }}$' \
        | grep -v README.org \
        | grep -v make_txtar.sh \
        | grep -v go.sum \
        | grep -v go.mod \
        | grep -v Makefile \
        | grep -v logger.go \
        | grep -v options.go \
        | grep -vE hooks/*.py \
        # | grep -v cmd/main.go \
        # | grep -v {{ cookiecutter.project_slug }}.go \

} | tee $tmp/filelist.txt
tar -cf $tmp/{{ cookiecutter.project_slug }}.tar -T $tmp/filelist.txt
mkdir -p $tmp/{{ cookiecutter.project_slug }}
tar xf $tmp/{{ cookiecutter.project_slug }}.tar -C $tmp/{{ cookiecutter.project_slug }}
rg --files $tmp/{{ cookiecutter.project_slug }}
txtar-c $tmp/{{ cookiecutter.project_slug }} | pbcopy

rm -rf $tmp
