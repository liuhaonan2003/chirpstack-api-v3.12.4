.PHONY: requirements as-external-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

all: requirements as-external-api

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

as-external-api:
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/application.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/frameLog.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/device.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/deviceProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/deviceQueue.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/gateway.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/gatewayProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/internal.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/multicastGroup.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/networkServer.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/organization.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/profiles.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/serviceProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/user.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/test.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportDetail.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/smsTemplate.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/role.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/modules.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/logs.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/warning.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/warningRule.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/workOrder.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportCheckDay.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportCheckMonth.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportDay.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportDetail.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportGwDay.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportMonth.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportOrganizationDay.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportOrganizationMonth.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportOrganizationWeek.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/reportWeek.proto


