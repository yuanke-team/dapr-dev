rm -rf ./pkg
# rm -rf ./target/wasm32-wasi
rm -rf ../../image-api-go/lib/classify_bg.wasm

rustup override set 1.50.0
rustwasmc  build --enable-ext
cp ./pkg/classify_bg.wasm ../../image-api-go/lib/classify_bg.wasm

echo -e "finished build functions/classify ..."