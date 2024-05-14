#!/bin/bash

CONTRACT_PATH=/app/api/sols
ABI_PATH=/app/idl/contracts

# solc --help
for f in ${CONTRACT_PATH}/*.sol; do
  solc --abi $f -o ${ABI_PATH} --overwrite --ignore-missing
  solc --bin $f -o ${ABI_PATH} --overwrite --ignore-missing
done

for f in ${ABI_PATH}/*.abi; do
  p=$(echo $f | sed -r 's/\.[^.]*$//')
  name=$(echo $p | sed -r 's:.*/::')
  abigen --abi $f --bin ${p}.bin --pkg contract --type $name --out $p.go
done

#! remove permission
chmod -R 777 ${ABI_PATH}
