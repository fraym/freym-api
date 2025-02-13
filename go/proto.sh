#! /bin/bash

cd ../api/grpc

for filename in ./auth/management/*.proto; do
	params="$params --go_opt=Mauth/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/auth/managementpb"
	params="$params --go-grpc_opt=Mauth/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/auth/managementpb"
done

for filename in ./projections/management/*.proto; do
	params="$params --go_opt=Mprojections/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/projections/managementpb"
	params="$params --go-grpc_opt=Mprojections/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/projections/managementpb"
done

for filename in ./projections/delivery/*.proto; do
	params="$params --go_opt=Mprojections/delivery/$(basename $filename)=github.com/fraym/freym-api/go/proto/projections/deliverypb"
	params="$params --go-grpc_opt=Mprojections/delivery/$(basename $filename)=github.com/fraym/freym-api/go/proto/projections/deliverypb"
done

for filename in ./crud/management/*.proto; do
	params="$params --go_opt=Mcrud/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/crud/managementpb"
	params="$params --go-grpc_opt=Mcrud/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/crud/managementpb"
done

for filename in ./crud/delivery/*.proto; do
	params="$params --go_opt=Mcrud/delivery/$(basename $filename)=github.com/fraym/freym-api/go/proto/crud/deliverypb"
	params="$params --go-grpc_opt=Mcrud/delivery/$(basename $filename)=github.com/fraym/freym-api/go/proto/crud/deliverypb"
done

for filename in ./streams/management/*.proto; do
	params="$params --go_opt=Mstreams/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/streams/managementpb"
	params="$params --go-grpc_opt=Mstreams/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/streams/managementpb"
done

for filename in ./sync/management/*.proto; do
	params="$params --go_opt=Msync/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/sync/managementpb"
	params="$params --go-grpc_opt=Msync/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/sync/managementpb"
done

for filename in ./deployments/management/*.proto; do
	params="$params --go_opt=Mdeployments/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/deployments/managementpb"
	params="$params --go-grpc_opt=Mdeployments/management/$(basename $filename)=github.com/fraym/freym-api/go/proto/deployments/managementpb"
done

cd ../../go/proto 

rm -rf ./sync
rm -rf ./streams
rm -rf ./projections
rm -rf ./crud
rm -rf ./deployments
rm -rf ./auth

cd ../../api/grpc

protoc \
  	--proto_path=./ \
	--go_out=../../go/proto \
	--go-grpc_out=../../go/proto \
	${params[@]} \
	--go_opt=module=github.com/fraym/freym-api/go/proto \
	--go-grpc_opt=module=github.com/fraym/freym-api/go/proto \
	--go_opt=default_api_level=API_OPAQUE \
	**/**/*.proto

cd ../../go 
go mod tidy
go get -u all
go mod tidy

