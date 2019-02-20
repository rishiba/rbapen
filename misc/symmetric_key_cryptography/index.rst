==================================
Symmetric encryption using openssl
==================================

Generate a file
===============

::

  $ dd if=/dev/urandom bs=1M count=10 of=10mb_file
  10+0 records in
  10+0 records out
  10485760 bytes (10 MB, 10 MiB) copied, 0.0634235 s, 165 MB/s


Find the md5sum
===============

::

  $ md5sum 10mb_file
  5c454dfdbd43bd6e410b3aae46c86a94  10mb_file

Append the md5sum in the filename
==================================

::


  $ mv 10mb_file 10mb_file_5c454dfdbd43bd6e410b3aae46c86a94


Check the type of the data
==========================

::

  $ file 10mb_file_5c454dfdbd43bd6e410b3aae46c86a94
  10mb_file_5c454dfdbd43bd6e410b3aae46c86a94: data

Lets generate a new key for encryption
======================================

::

  $ echo "MYKEY" > key

* For generating a strong key use the following command ``openssl rand -out key -hex 48`` or ``openssl rand -out key -base64 48``

Use the key for encryting the file
===================================

::

  $ openssl enc -aes-256-cbc -salt -in 10mb_file_5c454dfdbd43bd6e410b3aae46c86a94 -out 10mb_file_encrypted  -pass file:./key

Command details are 


.. csv-table:: 
  :header: "Option", "Meaning", "Remarks"

  "enc", "encrypt", "ask openssl to encrypt"
  "-aes-256-cbc", "encryption algorithm", "the algorithm to use for encryption, there can be several others like des etc."
  "-salt", " ..todo ", ""
  "in", "the file to encrypt"
  "out", "the name of the encrypted file"
  "pass", "the key to be used for encryption"

Check the md5sum of the file
===========================

::

  $ md5sum 10mb_file_encrypted 
  538b47f034ca11652d9c3cb9dd97eaa8  10mb_file_encrypted

Check the data of the file
===========================

::

  $ file 10mb_file_* 
  10mb_file_5c454dfdbd43bd6e410b3aae46c86a94: data
  10mb_file_encrypted:                        openssl enc'd data with salted password


Decrypt the data of the file using the same key
===============================================

::

  $ openssl enc -d -aes-256-cbc  -in 10mb_file_encrypted -out 10mb_file_decrypted  -pass file:./key

Command details are 

.. csv-table:: 
  :header: "Option", "Meaning", "Remarks"

  "enc", "encrypt", "ask openssl to encrypt"
  "-aes-256-cbc", "encryption algorithm", "the algorithm to use for encryption, there can be several others like des etc."
  "-d", " decrypt the input file", ""
  "in", "the file to decrypt"
  "out", "the name of the decrypted file"
  "pass", "the key to be used for decryption"

Check the md5sum of the file
=============================

::

  $ md5sum 10mb_file_*
  5c454dfdbd43bd6e410b3aae46c86a94  10mb_file_5c454dfdbd43bd6e410b3aae46c86a94
  5c454dfdbd43bd6e410b3aae46c86a94  10mb_file_decrypted
  538b47f034ca11652d9c3cb9dd97eaa8  10mb_file_encrypted

