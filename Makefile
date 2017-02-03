PROTEIN_PATH=${GOPATH}/src/github.com/znly/protein
PROTOBUF_SCHEMA=protobuf_schema

all: protobuf

protobuf-clean:
	rm -f ${PROTEIN_PATH}/*.pb.go

protobuf: protobuf-clean
	protoc --proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		   --proto_path=${GOPATH}/src/github.com/gogo/ \
		   --proto_path=${PROTEIN_PATH}/protobuf/ \
		   --gogofaster_out=${PROTEIN_PATH}/ \
		   ${PROTEIN_PATH}/protobuf/*.proto
	sed -i '' 's#google_protobuf "google/protobuf"#google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"#g' ${PROTEIN_PATH}/*.pb.go
	sed -i '' 's#google_protobuf1 "google/protobuf"#google_protobuf1 "github.com/gogo/protobuf/types"#g' ${PROTEIN_PATH}/*.pb.go
	sed -i '' 's#"protobuf/gogoproto"#"github.com/gogo/protobuf/gogoproto"#g' ${PROTEIN_PATH}/*.pb.go
