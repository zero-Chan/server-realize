docker_respository=zero-chan/redis-cluster

resposition_tag=1

docker_full_respository="${docker_respository}:${resposition_tag}"

manager_docker_full_respository="${docker_respository}-manager:${resposition_tag}"

.PHONY: manager
all: manager
	docker build -t "${docker_full_respository}" src

manager:
	docker build -t "${manager_docker_full_respository}" manager

.PHONY: clean
clean:
	docker rmi "${docker_full_respository}"
	docker rmi "${manager_docker_full_respository}"
