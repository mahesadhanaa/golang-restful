# Simple RESTful API using Golang and Postgres

##### to run the serve locally
$ export DB_NAME=db_name DB_USERNAME=username DB_PASSWORD=password

$ make run

##### to run in Docker
$ export DB_NAME=db_name DB_USERNAME=username DB_PASSWORD=password

$ export CI_DOCKER_IMAGE_OUTPUT

$ make ci_docker_image