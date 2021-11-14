proto:
	@echo "-- generate proto"
	protoc ./schema/user/user_stream/* --go_out=plugins=grpc:./pkg/grpc
	protoc ./schema/task/task_stream/* --go_out=plugins=grpc:./pkg/grpc