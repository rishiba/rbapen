======================
Simple REST API Server
======================

Introduction
============

Writing a simple REST API server is very easy in golang. In this article we will write a simple REST API server.

Problem Statement
=================

Write an REST API Server which will provide the following functionalities. The code should be commented well so that reader can understand the code by reading the comments.

Requirements
-------------

.. csv-table:: ""
  :header: "Verbs", "EndPoint", "Functionality", "OnSucess", "OnError"

  "POST", "/users", "add a user to the server", "return errorCode as 0, errMsg as nil", "if the user is already there return an error, with ``errorCode=1`` and the error string as well."
  "GET", "/users/{username}", "get the details of a single user", "return errorCode as 0, errMsg as nil", "if the user is not there return an error with ``errorCode=2`` and the error string as well."
  "GET", "/users", "get the details of all the users", "return errorCode as 0, errMsg as nil", "if no users are there return error with ``errorCcode=3`` and error string as well."
  "DELETE", "/users/{username}","delete a specific user", "return errorCode as 0, errMsg as nil", "delete the user if it is there, else return ``errorCode=2``"
  "PATCH", "/users/{username}","change any specific field of a user, the user should pre-exist", "return errorCode as 0, errMsg as nil", "if the user is not present return ``errorCode=2``"
  "PUT", "/users/username","update the full data of the user.", "return errorCode as 0, errMsg as nil", "if the user not present return ERR_USER_ABSENT"

ErrorCodes and Error Messages
+++++++++++++++++++++++++++++

.. csv-table:: 
  :header: "ErrorConstant", "ErrorCode", "ErrorString"

  "ERR_SUCCESS", "0", "Success"
  "ERR_USER_PRESENT", "1", "User {username} already present"
  "ERR_USER_ABSENT", "2", "User {username} not present"
  "ERR_USERS_DB_EMPTY", "3", "No users present"

.. csv-table:: Frozen Delights!
   :header: "Treat", "Quantity", "Description"
   :widths: 15, 10, 30

   "Albatross", 2.99, "On a stick!"
   "Crunchy Frog", 1.49, "If we took the bones out, it wouldn't be
   crunchy, now would it?"
   "Gannet Ripple", 1.99, "On a stick!"

Code
====


References
==========
