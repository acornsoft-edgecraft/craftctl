#!/bin/bash

set -x
set -eE

ERROR_LOG_FILE="/tmp/edge.error.log"
SONOBUOY_OUTPUT_DIR=${SONOBUOY_OUTPUT_DIR:-/tmp/sonobuoy}

EDGE_NS="edge-benchmarks"
EDGE_DIR="/tmp/edge"

BENCHMARKS_ID="${BENCHMARKS_ID}"

# Clean up the output directory, just in case
rm -rf "${SONOBUOY_OUTPUT_DIR}"/*.tar.gz

# Run sonobuoy aggregator
if ! sonobuoy aggregator -v 3
then
  echo "error running sonobuoy" | tee -a ${ERROR_LOG_FILE}

  /edge-benchmarks/edge-summarize \
      --benchmarks-id "${BENCHMARKS_ID}" \
      --reason "error running sonobuoy" 2> ${ERROR_LOG_FILE}

  exit 1
fi

SONOBUOY_OUTPUT_FILE=$(ls -1t "${SONOBUOY_OUTPUT_DIR}"/*.tar.gz | head -1)

echo "SONOBUOY OUTPUT FILE :: ${SONOBUOY_OUTPUT_FILE}"

# Extract the result
mkdir -p "${EDGE_DIR}"

if ! tar -C "${EDGE_DIR}" \
         -xvf "${SONOBUOY_OUTPUT_FILE}" \
         --warning=no-timestamp 2> ${ERROR_LOG_FILE}
then
  echo "error extracting output file: \"${SONOBUOY_OUTPUT_FILE}\"" | tee -a ${ERROR_LOG_FILE}
  exit 1
fi

# Run edge-summarize
if ! /edge-benchmarks/edge-summarize \
      --benchmarks-id "${BENCHMARKS_ID}" \
      --results-dir "${EDGE_DIR}" 2> ${ERROR_LOG_FILE}
then
  echo "error running edge-summarize" | tee -a ${ERROR_LOG_FILE}
  exit 1
fi

if [[ "${DEBUG}" == "true" ]]; then
    sleep infinity
fi
