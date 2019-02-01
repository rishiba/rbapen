==================
Docker Cheat Sheet
==================

To extract a file from a docker image
=====================================

* ``docker run --rm --entrypoint cat /usr/bin/docker-entrypoint.sh 3da585a74e97``

Pulling centos6.7 machine
=========================

* ``docker pull IMAGENAME``
* ``docker pull centos:6.7``

List the images
===============

* ``docker images``
* ``docker image list``

::

  # docker image list
  REPOSITORY          TAG                 IMAGE ID      CREATED        SIZE
  minio/minio         latest              78a749dfdd01  12 days ago    36.6MB
  golang              1.11.4-alpine3.7    0eea6b1e32f1  2 weeks ago    309MB
  alpine              3.7                 9bea9e12e381  2 weeks ago    4.21MB
  fedora              29                  25e6809f6fab  3 weeks ago    274MB
  google/cadvisor     latest              75f88e3ec333  13 months ago  62.2MB

Remove images
=============


* ``docker image rm 36ab0c269d38``

run an image
============

* ``docker run -d myimage``

List the running images
=======================

* ``docker ps``




* ``docker container list``
  - list the running containers

Stop a running container
=========================

* ``docker stop CONTAINER_ID``


Get a bash shell in the running docker image
============================================

* ``docker exec -it container_id /bin/bash``



Starting the container with name as buildMachine
================================================

``docker run  -d --name myMachineName  -i -t 000c5746fa52``

Installing build related packages and other tool in the package
===============================================================

Docker Pass enviornment variable
================================

Install cockpit-docker for UI based management
==============================================

dnf install  -y   cockpit-docker

Docker make a new image
=======================

docker  commit -a "Rishi Agrawal <rishi.b.agrawal@gmail.com>" -m "Commit message" myImageName:myImageTag1


Docker remove all the containers ``docker rm $(docker ps -a -q)``

Docker remove all the images ``docker rmi $(docker images -q)``

Docker remove selected images docker image list| grep none | awk ' {print $3}'  | xargs docker image rm
https://stackoverflow.com/questions/21398087/how-can-i-delete-dockers-images

Some important command line  options
=====================================


``docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres``

We have provided several options to the docker run command:

— rm: Automatically remove the container and it’s associated file system upon exit. In general, if we are running lots of short term containers, it is good practice to to pass rm flag to the docker run command for automatic cleanup and avoid disk space issues. We can always use the v option (described below) to persist data beyond the lifecycle of a container
— name: An identifying name for the container. We can choose any name we want. Note that two existing (even if they are stopped) containers cannot have the same name. In order to re-use a name, you would either need pass the rm flag to the docker run command or explicitly remove the container by using the command docker rm [container name].
-e: Expose environment variable of name POSTGRES_PASSWORD with value docker to the container. This environment variable sets the superuser password for PostgreSQL. We can set POSTGRES_PASSWORD to anything we like. I just choose it to be docker for demonstration. There are additional environment variables you can set. These include POSTGRES_USER and POSTGRES_DB. POSTGRES_USER sets the superuser name. If not provided, the superuser name defaults to postgres. POSTGRES_DB sets the name of the default database to setup. If not provided, it defaults to the value of POSTGRES_USER.
-d: Launches the container in detached mode or in other words, in the background.
-p: Bind port 5432 on localhost to port 5432 within the container. This option enables applications running out side of the container to be able to connect to the Postgres server running inside the container.
-v: Mount $HOME/docker/volumes/postgres on the host machine to the container side volume path /var/lib/postgresql/data created inside the container. This ensures that postgres data persists even after the container is removed.
