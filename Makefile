PROTO_TARGETS = tweet

proto:
	@echo $(PROTO_TARGETS); \
	for t in $(PROTO_TARGETS); do \
		echo "${t} build"; \
		protoc $(PROJECT_PATH)/proto/$$t.proto \
			--go_out=./proto --proto_path=$(PROJECT_PATH) \
			--go-grpc_out=./proto; \
	done

docker-build:
	docker build -t cliter -f api/Dockerfile .;\
	docker tag cliter gcr.io/voltaic-quest-176113/cliter:latest 

docker-push:
	docker push gcr.io/voltaic-quest-176113/cliter:latest 

clean:
	rm -rf proto/*.pb.go
