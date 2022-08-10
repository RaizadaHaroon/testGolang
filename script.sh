echo "source director $1"
BUILD_DIR=$1
cd BUILD_DIR
git clone git@github.com:RaizadaHaroon/testGolang.git
cd testGolang
go run main.go
