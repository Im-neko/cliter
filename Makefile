PROTO_TARGETS = tweet

proto:
	@echo $(PROTO_TARGETS); \
	for t in $(PROTO_TARGETS); do \
		echo "${t} build"; \
		protoc $(PROJECT_PATH)/proto/$$t.proto \
			--go_out=./proto --proto_path=$(PROJECT_PATH) \
			--go-grpc_out=./proto; \
	done

clean:
	rm -rf proto/*.pb.go
