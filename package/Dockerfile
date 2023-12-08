FROM registry.suse.com/bci/bci-base:15.4 as intermediate

LABEL stage=intermediate

RUN zypper --non-interactive update \
    && zypper --non-interactive install \
    git

RUN git clone -b v0.6.15 --depth 1 https://github.com/aquasecurity/kube-bench.git

FROM registry.suse.com/bci/bci-base:15.4
ARG kube_bench_tag=0.6.15
ARG sonobuoy_version=0.56.15

RUN zypper --non-interactive update \
    && zypper --non-interactive install \
    curl \
    jq \
    vim \
    systemd \
    tar \
    awk \
    gzip
    
RUN curl -sLf https://github.com/vmware-tanzu/sonobuoy/releases/download/v${sonobuoy_version}/sonobuoy_${sonobuoy_version}_linux_amd64.tar.gz | tar -xvzf - -C /usr/bin sonobuoy
RUN curl -sLf https://github.com/aquasecurity/kube-bench/releases/download/v${kube_bench_tag}/kube-bench_${kube_bench_tag}_linux_amd64.tar.gz | tar -xvzf - -C /usr/bin

COPY --from=intermediate /kube-bench/cfg /etc/kube-bench/cfg/

COPY package/cfg/config.yaml /etc/kube-bench/cfg/
COPY package/cfg/k3s-1.23/ /etc/kube-bench/cfg/k3s-1.23/
COPY package/cfg/mk8s-1.23/ /etc/kube-bench/cfg/mk8s-1.23/

COPY package/run.sh \
    package/run-kube-bench.sh \  
    package/helper_scripts/check_for_k3s_etcd.sh \
    package/helper_scripts/microk8s.sh \
    /usr/bin/

RUN mkdir edge-benchmarks

COPY bin/edge-summarize /edge-benchmarks
COPY conf/ /edge-benchmarks/conf/

CMD ["run.sh"]
