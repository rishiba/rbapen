# for an interface, traffic to - from to a particular host and specific port.
sudo tcpdump -i ens1f0   -A  host 172.30.250.14 and port 4800

# For local host
sudo tcpdump -i lo  -A  host 172.30.250.14 and port 4800


# for allowing root login on ssh 

You also need to edit /etc/ssh/sshd_config, and comment out the following line:

PermitRootLogin without-password
Just below it, add the following line:

PermitRootLogin yes
Then restart SSH:

service ssh reload

