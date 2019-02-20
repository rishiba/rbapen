
Reference : https://linuxconfig.org/how-to-install-kubernetes-on-ubuntu-18-04-bionic-beaver-linux
--------------------------------------------
Configure Network to have a hostonly network
--------------------------------------------


---------------------------
Remove swap from /etc/fatab
---------------------------



Run the following commands to have your machine install all the required
packages


Take a snapshot

sudo apt-get update
sudo apt-get install -y docker.io
sudo systemctl enable docker
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
sudo apt-add-repository "deb http://apt.kubernetes.io/ kubernetes-xenial main"
sudo apt-get install -y kubeadm
sudo swapoff -a


Take a snapshot
Run the following command to PUll the required images
kubeadm config images pull

Take a snapshot

Run the command 



``kubeadm init --apiserver-advertise-address=192.168.56.101 --pod-network-cidr=10.244.0.0/16``



::

  rishi@kubemaster:~$ sudo kubeadm init --apiserver-advertise-address=192.168.56.101 --pod-network-cidr=10.244.0.0/16
  [init] Using Kubernetes version: v1.13.1
  [preflight] Running pre-flight checks
  [preflight] Pulling images required for setting up a Kubernetes cluster
  [preflight] This might take a minute or two, depending on the speed of your internet connection
  [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
  [kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
  [kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
  [kubelet-start] Activating the kubelet service
  [certs] Using certificateDir folder "/etc/kubernetes/pki"
  [certs] Generating "ca" certificate and key
  [certs] Generating "apiserver" certificate and key
  [certs] apiserver serving cert is signed for DNS names [kubemaster kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 192.168.56.101]
  [certs] Generating "apiserver-kubelet-client" certificate and key
  [certs] Generating "etcd/ca" certificate and key
  [certs] Generating "etcd/server" certificate and key
  [certs] etcd/server serving cert is signed for DNS names [kubemaster localhost] and IPs [192.168.56.101 127.0.0.1 ::1]
  [certs] Generating "etcd/peer" certificate and key
  [certs] etcd/peer serving cert is signed for DNS names [kubemaster localhost] and IPs [192.168.56.101 127.0.0.1 ::1]
  [certs] Generating "etcd/healthcheck-client" certificate and key
  [certs] Generating "apiserver-etcd-client" certificate and key
  [certs] Generating "front-proxy-ca" certificate and key
  [certs] Generating "front-proxy-client" certificate and key
  [certs] Generating "sa" key and public key
  [kubeconfig] Using kubeconfig folder "/etc/kubernetes"
  [kubeconfig] Writing "admin.conf" kubeconfig file
  [kubeconfig] Writing "kubelet.conf" kubeconfig file
  [kubeconfig] Writing "controller-manager.conf" kubeconfig file
  [kubeconfig] Writing "scheduler.conf" kubeconfig file
  [control-plane] Using manifest folder "/etc/kubernetes/manifests"
  [control-plane] Creating static Pod manifest for "kube-apiserver"
  [control-plane] Creating static Pod manifest for "kube-controller-manager"
  [control-plane] Creating static Pod manifest for "kube-scheduler"
  [etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
  [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
  [apiclient] All control plane components are healthy after 21.009607 seconds
  [uploadconfig] storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
  [kubelet] Creating a ConfigMap "kubelet-config-1.13" in namespace kube-system with the configuration for the kubelets in the cluster
  [patchnode] Uploading the CRI Socket information "/var/run/dockershim.sock" to the Node API object "kubemaster" as an annotation
  [mark-control-plane] Marking the node kubemaster as control-plane by adding the label "node-role.kubernetes.io/master=''"
  [mark-control-plane] Marking the node kubemaster as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
  [bootstrap-token] Using token: gw3h6o.ko6yzgcpgjb0ri3y
  [bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
  [bootstraptoken] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
  [bootstraptoken] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
  [bootstraptoken] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
  [bootstraptoken] creating the "cluster-info" ConfigMap in the "kube-public" namespace
  [addons] Applied essential addon: CoreDNS
  [addons] Applied essential addon: kube-proxy

  Your Kubernetes master has initialized successfully!

  To start using your cluster, you need to run the following as a regular user:

    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config

  You should now deploy a pod network to the cluster.
  Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
    https://kubernetes.io/docs/concepts/cluster-administration/addons/

  You can now join any number of machines by running the following on each node
  as root:

    kubeadm join 192.168.56.101:6443 --token gw3h6o.ko6yzgcpgjb0ri3y --discovery-token-ca-cert-hash sha256:e61f19d4afe71963c3e3d28b5c5199ae9f473619dedca6b30f56e78baa6f58ec




=====================
    Start the cluster
=====================


    $  mkdir -p $HOME/.kube
rishi@kubemaster:~$     sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
rishi@kubemaster:~$     sudo chown $(id -u):$(id -g) $HOME/.kube/config
rishi@kubemaster:~$ sudo kubectl apply -f kube-flannel.yml



======================
Create overlay network
======================


sudo kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml


=================
Check the cluster
=================


    rishi@kubemaster:~$ kubectl get pods --all-namespaces
NAMESPACE     NAME                                 READY   STATUS    RESTARTS   AGE
kube-system   coredns-86c58d9df4-5vzxh             1/1     Running   0          4m13s
kube-system   coredns-86c58d9df4-6nxr2             1/1     Running   0          4m13s
kube-system   etcd-kubemaster                      1/1     Running   0          3m26s
kube-system   kube-apiserver-kubemaster            1/1     Running   0          3m20s
kube-system   kube-controller-manager-kubemaster   1/1     Running   0          3m38s
kube-system   kube-flannel-ds-amd64-j6lzl          1/1     Running   0          22s
kube-system   kube-proxy-cf8lk                     1/1     Running   0          4m13s
kube-system   kube-scheduler-kubemaster            1/1     Running   0          3m25s


Installing slave

Run ``sudo tar -cf  archives.tar /var/cache/apt/archives/`` on master

Copy the file to the slave using scp

Untar the file

Check the files

ls var/cache/apt/archives/
cd  var/cache/apt/archives/

Install the debs 

dpkg -i `*.deb


Start make the slave join the master - note that the token is valid for 24 hours only. 

.. todo:: generate a new token after 24 hours and add a node.

``kubeadm join 192.168.56.101:6443 --token gw3h6o.ko6yzgcpgjb0ri3y --discovery-token-ca-cert-hash sha256:e61f19d4afe71963c3e3d28b5c5199ae9f473619dedca6b30f56e78baa6f58ec``


------------------
List all the nodes
------------------


$ kubectl get nodes
NAME          STATUS     ROLES    AGE     VERSION
kubemaster    Ready      master   33m     v1.13.1
kubeslave01   NotReady   <none>   9m14s   v1.13.1




k8s.gcr.io/kube-proxy                v1.13.1             fdb321fd30a0        4 weeks ago         80.2MB
k8s.gcr.io/kube-apiserver            v1.13.1             40a63db91ef8        4 weeks ago         181MB
k8s.gcr.io/kube-scheduler            v1.13.1             ab81d7360408        4 weeks ago         79.6MB
k8s.gcr.io/kube-controller-manager   v1.13.1             26e6f1db2a52        4 weeks ago         146MB
k8s.gcr.io/coredns                   1.2.6               f59dcacceff4        2 months ago        40MB
k8s.gcr.io/etcd                      3.2.24              3cab8e1b9802        3 months ago        220MB
k8s.gcr.io/pause                     3.1                 da86e6ba6ca1        12 months ago       742kB


Create web ui

kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml



Install only docker
sudo dpkg -i docker.io_18.06.1-0ubuntu1~16.04.2_amd64.deb  libltdl7_2.4.6-0.1_amd64.deb


sudo dpkg -i bridge-utils_1.5-9ubuntu1_amd64.deb cgroupfs-mount_1.2_all.deb cri-tools_1.12.0-00_amd64.deb  ebtables_2.0.10.4-3.4ubuntu2.16.04.2_amd64.deb kubeadm_1.13.1-00_amd64.deb kubectl_1.13.1-00_amd64.deb kubelet_1.13.1-00_amd64.deb kubernetes-cni_0.6.0-00_amd64.deb  socat_1.7.3.1-1_amd64.deb ubuntu-fan_0.12.8~16.04.2_all.deb
