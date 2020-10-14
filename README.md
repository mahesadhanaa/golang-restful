# Simple RESTful API using Golang and Postgres

##### to run the serve locally
- install postgres on locally for depedencies

- Create env vars

    $ export DB_NAME=db_name DB_USERNAME=username DB_PASSWORD=password

- run makefile

    $ make run

##### to run in Docker
- run postgress image in docker

- Create env vars

    `$ export CI_DOCKER_IMAGE_OUTPUT`

- Build Docker image

    `$ make ci_docker_image`

- Run in docker

    `$ docker run --name web_server -d -p 8000:8000 $CI_DOCKER_IMAGE_OUTPUT`