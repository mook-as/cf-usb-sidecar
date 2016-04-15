#!/bin/sh -x

CSM_IMAGE_NAME="amazon-rds-mysql:release"
MYSQL_DB_CLASS="db.t2.micro"
DB_SIZE="5"
MULTIAZ="false"
MYSQL_ROOT_USER="root"
MYSQL_ROOT_PASSWORD="rootsecret123"
CSM_API_KEY="csm-auth-token"


if [ -f /sbin/md5 ]
then
	MYSQL_RDS_INSTANCE=`date | md5 -r | head -c 8`
else
	MYSQL_RDS_INSTANCE=`date | md5sum | head -c 8`
fi


: ${AWS_ACCESS_KEY_ID:?"Need to set AWS_ACCESS_KEY_ID"}
: ${AWS_SECRET_ACCESS_KEY:?"Need to set AWS_SECRET_ACCESS_KEY non-empty"}
: ${AWS_RDS_REGION:?"Need to set AWS_RDS_REGION non-empty"}

# export AWS_ACCESS_KEY_ID=<add yours>
# export AWS_SECRET_ACCESS_KEY=<add yours>
# export VPC_ID=<add yours>
# export AWS_RDS_REGION=<add yours>

docker run --name amazon-rds-mysql -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} -e AWS_RDS_REGION=${AWS_RDS_REGION} -e MYSQL_DB_CLASS=${MYSQL_DB_CLASS} -e DB_SIZE=${DB_SIZE} -e MULTIAZ=${MULTIAZ} -e MYSQL_ROOT_USER=${MYSQL_ROOT_USER} -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -e MYSQL_RDS_INSTANCE=rds${MYSQL_RDS_INSTANCE} -e CSM_API_KEY=${CSM_API_KEY} -p 8081:8081 -i -t ${CSM_IMAGE_NAME}