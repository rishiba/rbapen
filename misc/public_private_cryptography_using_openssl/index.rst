============================================
Encrypt/Decrypt a file using Assymetric Keys
============================================

Generate a small file say with 64 bit of data
=============================================

Encryption with assymetric encryption is highly computer intenseive as compared to the symmetric keys

.. todo:: write this section

::

  $ openssl rand -out data  -hex 64

See the data
------------

::

  $ cat data
  1c8fef6273bd947cc3bd2828d2d23d4850d778834e38ca54e0faae013a18c4ae3a4a56cf564ba6fa465b963834573872d82680bd9d0fae6bdc5f3232dfcbf99b

Generate a private key
======================

::

  $ openssl genrsa -out private_rsa.key 2048
  Generating RSA private key, 2048 bit long modulus
  .................+++
  ...................................+++
  e is 65537 (0x010001)

See the private key
-------------------

::

  rishi@rishilappy:~/11/rbapen/misc/public_private_cryptography_using_openssl$ cat private_rsa.key
  -----BEGIN RSA PRIVATE KEY-----
  MIIEpAIBAAKCAQEAq8oXbx3A4eQjySoz0P6BfB1riUAc9QbEdOHRnU7nxMJgm9aD
  t0jbFbHSRrsxZUmaRCG1m+HUPZmOr4UImOKkZ7nKqBl7dLcRECgw/2yNdS6aFd8d
  gX9BzwlV2z3LCjmfgeOgbWZSqqcshiukdZXe1HZNagYN+7Yd1ImCvj9sp/omYc2g
  YEr+fd4Mbx9sRiD+aXAI4tNmubLZLHdfLxTaRBU1nkCdKg/DIgBIxuznbSA+aB8L
  AtwGXf6cDktVSSWbcH602l0PmqFFNT62lb26SK09TYR65eqbtZIeNlsFe6/5kfaH
  qGo4fdeXF/YQ9WM0DTg0CLGcYBUaYgBYjsxSswIDAQABAoIBABXnOn2saX9juzxW
  6SxohtHwXlmtlwOWzU2pzTQlb1+i+PZlr32bCHQHrCvgijfD3qt5MNKjRJBmF89B
  QwtbOmMPSUuNlDQZc0+AEF9A0/hw7KyFOhncw+NF6XB+vHidSD0jbL0GItwnXIHu
  5pUnerXUL6KpCvdhDQSgU881/wCbquCekQufOhp/e5EoaPryyWEbX/7PsfIw75BJ
  Bwc+vZhUiGqbd/2n217KDkEY/3VAH469ItdCLBkR94ddiRPKCk4NPGx+Rz86DbpU
  uroIOwydZedOtDApRzkGk2z3FEMYo9RmFjmlTQNSTaZCQzadTzuDAanyn2bKqb9k
  g/RI6wECgYEA3g+6akU4zS7nFWVBG7kZVRID6gvBdpXrVHRHYRmapU1ByR8odaVm
  5h+QCP/6+8Uu1+dQlKMZ+Erz9DTggx0uCnq6/83K2E7SMvTF5OUewupWuy+miMax
  P08JpC4/sSOQrFFW7xaSlExA09xJFPjcB0riAr3zG9Rks0cdwH496pMCgYEAxgtx
  pNIlv0mGKKZ72BtvsY4tdyEZBcXSAPHELCcxRBDc9acMT28YmC4FeqAe4M2T3154
  4F1mCsnwx6usvM7bznnC5WbthW9OvFW4RBBcoqmwv32uC6V9hAj5kszXCIsyVFQM
  tGVfV3XeTNX29RjR/oir0RBH1erYNsI27RZpa2ECgYEAobtq1kOcXzSt8ZNRnFNv
  89mvKCNvwGLohxY5dqsjSwm9xDIBUc4p56lNSoDF1+GBgJZkxh0UqRZOH6rzagsy
  oUdKM007U9f0mFBWYaPOh8ANcz/9Vtj/91AlgYJ9uRJiKF8FphUWpRP0k1l5kOC/
  wD/15HFcQwSqYziEntJTXikCgYBb2AjPECozSOyvw22hri+hXzdUjQPmunQGvhhl
  BZQCZfOi9OJvoCz+vjEKSmGnjY8rQsu5+XaOQFKqtsH4GdK0MaGh85bpj9Iq43H+
  FrxvL1TpqWZivk+0QNlFLDv5GVyw4P23V1/c1qvOZFgQGH+ilE1iOjOmUDAPTQhd
  8UvYYQKBgQCqUCCIAUPraT3hR/fdZ8qM1iBbSGDIuK83qAcQKiWuEKRZrdAiPG5X
  LQV+Lh1enUQvYjKusdR2dQZ/XiGG1TvbQbZY7XolvifEN3R6CjY+/25Lq1n5wpq8
  vBbqcxAw29cNEQNsBBVx1+E5livUd/Iwj7vZZI93RZKgMagatmOzAg==
  -----END RSA PRIVATE KEY-----

Generate a public key from that private key
===========================================

::

  $ openssl rsa -pubout -in private_rsa.key -out public_rsa.key
  writing RSA key

See the public key
-------------------

::

  $ cat public_rsa.key
  -----BEGIN PUBLIC KEY-----
  MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAq8oXbx3A4eQjySoz0P6B
  fB1riUAc9QbEdOHRnU7nxMJgm9aDt0jbFbHSRrsxZUmaRCG1m+HUPZmOr4UImOKk
  Z7nKqBl7dLcRECgw/2yNdS6aFd8dgX9BzwlV2z3LCjmfgeOgbWZSqqcshiukdZXe
  1HZNagYN+7Yd1ImCvj9sp/omYc2gYEr+fd4Mbx9sRiD+aXAI4tNmubLZLHdfLxTa
  RBU1nkCdKg/DIgBIxuznbSA+aB8LAtwGXf6cDktVSSWbcH602l0PmqFFNT62lb26
  SK09TYR65eqbtZIeNlsFe6/5kfaHqGo4fdeXF/YQ9WM0DTg0CLGcYBUaYgBYjsxS
  swIDAQAB
  -----END PUBLIC KEY-----

Encrypt the file using the public key
======================================

::

  openssl rsautl -encrypt -inkey public_rsa.key -pubin -in data -out data_encrypted_using_public

Decrpyt the file using the private key
======================================

::

  openssl rsautl -decrypt -inkey private_rsa.key  -in data_encrypted_using_public -out data_decrypted_using_private_key

Check the md5sum of the files
=============================

::

  $ md5sum data*
  3454a4c64059bb1e799efd1a5a16ada2  data
  3454a4c64059bb1e799efd1a5a16ada2  data_decrypted_using_private_key
  584c77dd702945e95bc035bebbaad724  data_encrypted_using_public
