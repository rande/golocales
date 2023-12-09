
#!/bin/sh

mkdir -p data

if [ ! -d "data/common" ]; then
  echo "Download the CLDR data"
  
  curl https://unicode.org/Public/cldr/44/core.zip -o data/core.zip
  (cd data && unzip core.zip)
else
  echo "CLDR data already exists"
fi

export PATH=$PATH:`go env GOPATH`/bin
export CLDR_DIR=`pwd`/data/common
export LOCALE_DIR=`pwd`/locales

mkdir -p locales
rm -rf locales/*

(cd generator && go run .)

cat locales/fr/fr.go

du -sh locales

go test .
