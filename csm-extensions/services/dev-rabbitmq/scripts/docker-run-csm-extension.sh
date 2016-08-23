#!/bin/sh

DOCKER_IMAGE="rabbitmq"
DOCKER_IMAGE_TAG="3.6.0-management"
SIDECAR_LOG_LEVEL="debug"
SIDECAR_DEV_MODE="true"
SIDECAR_API_KEY=${SIDECAR_EXTENSION_TOKEN}

if [ ! -z ${DOCKER_HOST} ]
then
    export DOCKER_HOST_IP=`echo ${DOCKER_HOST} | cut -d "/" -f 3 | cut -d ":" -f 1`
else
    export DOCKER_HOST_IP=`ip route get 8.8.8.8 | awk 'NR==1 {print $NF}'`
fi

DOCKER_ENDPOINT=http://${DOCKER_HOST_IP}:4445

docker run --name ${SIDECAR_EXTENSION_IMAGE_NAME} \
	-p 8094:8081 \
	-e DOCKER_ENDPOINT=${DOCKER_ENDPOINT} \
	-e DOCKER_IMAGE=${DOCKER_IMAGE} \
	-e RABBIT_SERVICE_PORTS_POOL_START=${SIDECAR_EXTENSION_SVC_PORTS_START} \
	-e RABBIT_SERVICE_PORTS_POOL_END=${SIDECAR_EXTENSION_SVC_PORTS_END} \
	-e DOCKER_IMAGE_TAG=${DOCKER_IMAGE_TAG} \
	-e SIDECAR_LOG_LEVEL=${SIDECAR_LOG_LEVEL} \
	-e SIDECAR_API_KEY=${SIDECAR_API_KEY} \
	-e SIDECAR_DEV_MODE=${SIDECAR_DEV_MODE} \
	-d ${SIDECAR_EXTENSION_IMAGE_NAME}:${SIDECAR_EXTENSION_IMAGE_TAG}
