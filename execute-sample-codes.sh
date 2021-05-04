set -e

DIR=${1}

for f in `find ${DIR} -type f -name "*.go"`; do
    echo "---- ${f} ----"
    make `dirname ${f}`/result.txt
done
