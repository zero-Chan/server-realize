docker_respository=zero-chan/base/centos7-users
resposition_tag=latest

docker_full_respository=${docker_respository}:${resposition_tag}

all:
	docker build -t "${docker_full_respository}" src

.PHONY: clean
clean:
	docker rmi "${docker_full_respository}"
