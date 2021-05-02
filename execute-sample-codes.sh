set -e

DIR=${1}

ls ${DIR}/*/*.go > /dev/null

for f in `ls ${DIR}/*/*.go`; do
    echo ${f}
done
