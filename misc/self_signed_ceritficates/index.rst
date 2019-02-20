========================
Self signed certificates
========================






Generate a key
==============



Generate a certificate
======================

Use the following configuration file


::

  $ cat root_certificate.conf
  [req]
  default_bits = 2048
  prompt = no
  default_md = sha256
  distinguished_name = dn

  [dn]
  C=IN
  ST=MAHARASHTRA
  L=PUNE
  O=OneGreatOrg
  OU=OneGreatOrgUnit1
  emailAddress=user1@onegreatorg.com
  CN = oneCommonName


Command

::

  openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.pem -config <( cat root_certificate.conf )


Read the generated certificate
===============================

You will see that the certificate has the inputs you gave in the configuration file

::

  openssl x509 -in rootCA.pem -text -noout



Now generate a new certificate from the root certificcate
=========================================================

Reference : https://medium.freecodecamp.org/how-to-get-https-working-on-your-local-development-environment-in-5-minutes-7af615770eec


Start the server
================

go run https_web_server/main.go


Call the API 
=============

curl --cacert rootCA.pem  -v  https://localhost:9999
