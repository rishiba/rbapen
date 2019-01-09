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

