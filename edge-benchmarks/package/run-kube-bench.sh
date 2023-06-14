#!/bin/bash

set -eEx
set -o errexit

echo "SONOBUOY RESULTS DIRECTORY :: ${SONOBUOY_RESULTS_DIR}"

HOSTNAME=$(cat /etc/hostname)
ERROR_LOG_FILE="${SONOBUOY_RESULTS_DIR}/error.log"
KB_OUTPUT_FILE="${SONOBUOY_RESULTS_DIR}/${HOSTNAME}.json"

targets="node"
benchmark=""
version=""

if [ "$IS_MASTER" = true ]; then
    targets="master,controlplane,etcd,policies,node"    
fi

if [ "$BENCHMARK" != "" ]; then
    benchmark=" --benchmark $BENCHMARK"
fi 

if [ "$VERSION" != "" ]; then
    version=" --version $VERSION"
fi


# Run kube-bench
kube-bench run --config-dir "/etc/kube-bench/cfg" --targets $targets $benchmark $version --json --outputfile ${KB_OUTPUT_FILE} -v=3 > "${ERROR_LOG_FILE}"

echo "KUBE_BENCH OUTPUT FILE :: ${KB_OUTPUT_FILE}"

# Inform sonobuoy worker about completion of the job
echo -n "${KB_OUTPUT_FILE}" > ${SONOBUOY_RESULTS_DIR}/done
