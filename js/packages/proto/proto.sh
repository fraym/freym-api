#! /bin/bash
rm -rf ./src/*

cd ../../../api/grpc

protoc \
	--plugin=../../js/node_modules/.bin/protoc-gen-ts_proto \
	--ts_proto_out=../../js/packages/proto/src \
	--ts_proto_opt=esModuleInterop=true \
	--ts_proto_opt=env=node \
	--ts_proto_opt=oneof=unions \
	--ts_proto_opt=forceLong=string \
	--ts_proto_opt=outputServices=grpc-js \
	--ts_proto_opt=useExactTypes=false \
	--ts_proto_opt=stringEnums=true \
	--ts_proto_opt=enumsAsLiterals=true \
	--ts_proto_opt=unrecognizedEnum=false \
	--ts_proto_opt=exportCommonSymbols=false \
	--ts_proto_opt=outputIndex=true \
	**/**/*.proto
