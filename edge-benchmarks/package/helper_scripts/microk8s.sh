#!/bin/bash

set -eEx

# Check 1.1.1
chmod 600 /var/snap/microk8s/current/args/kube-apiserver

# Check 1.1.2
chown root:root /var/snap/microk8s/current/args/kube-apiserver

# Check 1.1.3
chmod 600 /var/snap/microk8s/current/args/kube-controller-manager

# Check 1.1.4
chown root:root /var/snap/microk8s/current/args/kube-controller-manager

# Check 1.1.5
chmod 600 /var/snap/microk8s/current/args/kube-scheduler

# Check 1.1.6
chown root:root /var/snap/microk8s/current/args/kube-scheduler

# Check 1.1.7
chmod 600 /var/snap/microk8s/current/args/k8s-dqlite
chmod 600 /var/snap/microk8s/current/args/etcd

# Check 1.1.8
chown root:root /var/snap/microk8s/current/args/k8s-dqlite
chown root:root /var/snap/microk8s/current/args/etcd

# Check 1.1.9
chmod -R 600 /var/snap/microk8s/current/args/cni-network/

# Check 1.1.10
chown -R root:root /var/snap/microk8s/current/args/cni-network/

# Check 1.1.11
chmod -R 700 /var/snap/microk8s/current/var/kubernetes/backend/

# Check 1.1.12
chown -R root:root /var/snap/microk8s/current/var/kubernetes/backend/

# Check 1.1.13
chmod 600 /var/snap/microk8s/current/credentials/client.config

# Check 1.1.14
chown root:root /var/snap/microk8s/current/credentials/client.config

# Check 1.1.15
chmod 600 /var/snap/microk8s/current/credentials/scheduler.config

# Check 1.1.16
chown root:root /var/snap/microk8s/current/credentials/scheduler.config

# Check 1.1.17
chmod 600 /var/snap/microk8s/current/credentials/controller.config

# Check 1.1.18
chown root:root /var/snap/microk8s/current/credentials/controller.config

# Check 1.1.19
chown -R root:root /var/snap/microk8s/current/certs/

# Check 1.1.20
chmod -R 600 /var/snap/microk8s/current/certs/

# Check 1.1.21
chmod -R 600 /var/snap/microk8s/current/certs/


# Check 1.2.6
# Check 1.2.7
# Check 1.2.8
# microk8s enable rbac


# Check 4.1.1
chmod 600 /etc/systemd/system/snap.microk8s.daemon-kubelite.service

# Check 4.1.2
chown root:root /etc/systemd/system/snap.microk8s.daemon-kubelite.service

# Check 4.1.3
chmod 600 /var/snap/microk8s/current/credentials/proxy.config

# Check 4.1.4
chown root:root /var/snap/microk8s/current/credentials/proxy.config

# Check 4.1.5
chmod 600 /var/snap/microk8s/current/credentials/kubelet.config

# Check 4.1.6
chown root:root /var/snap/microk8s/current/credentials/kubelet.config

# Check 4.1.7
chmod 600 /var/snap/microk8s/current/certs/ca.crt

# Check 4.1.8
chown root:root /var/snap/microk8s/current/certs/ca.crt
