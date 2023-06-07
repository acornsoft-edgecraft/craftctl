FROM golang:1.19-alpine AS builder

LABEL maintainer="acornsoft"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY edge-benchmarks/go.mod edge-benchmarks/go.sum ./
RUN go mod download

# Copy the code into the container.
COPY edge-benchmarks .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o bin/edge-summarize main.go

## edge-benchamrks container image 생성
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

COPY --from=builder  package/run.sh \
    package/run-kube-bench.sh \    
    /usr/bin/

RUN mkdir edge-benchmarks

# COPY bin/edge-summarize /edge-benchmarks
# COPY conf/ /edge-benchmarks/conf/
COPY --from=builder ["/build/bin/edge-summarize", "edge-benchmarks"]
COPY --from=builder ["/build/conf/", "edge-benchmarks/conf"]

CMD ["run.sh"]
