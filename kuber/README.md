###
curl -LO https://dl.k8s.io/release/v1.23.0/bin/linux/amd64/kubectl

sudo install -o root -g root -m 0755 kubectl /usr/bin/kubectl

curl -LO https://dl.k8s.io/release/v1.23.0/bin/linux/amd64/kubeadm

sudo install -o root -g root -m 0755 kubeadm /usr/bin/kubeadm


kubeadm init \
--apiserver-advertise-address=192.168.20.10 \
--image-repository registry.aliyuncs.com/google_containers \
--pod-network-cidr=10.244.0.0/16