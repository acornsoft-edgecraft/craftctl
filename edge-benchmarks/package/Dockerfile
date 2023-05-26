FROM registry.suse.com/bci/bci-base:15.4 as intermediate

#Label the image for cleaning after build process
LABEL stage=intermediate

RUN zypper --non-interactive update \
    && zypper --non-interactive install \
    git

RUN git clone -b v0.6.12 --depth 1 https://github.com/aquasecurity/kube-bench.git

FROM registry.suse.com/bci/bci-base:15.4
ARG kube_bench_tag=0.6.12
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

# COPY package/cfg/cis-1.6-k3s /etc/kube-bench/cfg/cis-1.6-k3s

COPY package/run.sh \
    package/run-kube-bench.sh \    
    /usr/bin/

RUN mkdir edge-benchmarks

COPY bin/edge-summarize /edge-benchmarks
COPY conf/ /edge-benchmarks/conf/

CMD ["run.sh"]
