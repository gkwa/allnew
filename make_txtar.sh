rm -rf /tmp/allnew

rm -f /tmp/allnew.tar
rm -f /tmp/filelist.txt

{
    rg --files . \
        | grep -vE 'allnew$' \
        | grep -v README.org \
        | grep -v make_txtar.sh \
        | grep -v go.sum \
        | grep -v go.mod \
        | grep -v Makefile \
        # | grep -v cmd/main.go \
        # | grep -v options/options.go \
        # | grep -v allnew.go \

} | tee /tmp/filelist.txt
tar -cf /tmp/allnew.tar -T /tmp/filelist.txt
mkdir -p /tmp/allnew
tar xf /tmp/allnew.tar -C /tmp/allnew
rg --files /tmp/allnew
txtar-c /tmp/allnew | pbcopy
