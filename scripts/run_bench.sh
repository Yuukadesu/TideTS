#!/usr/bin/env bash
set -euo pipefail

HOST="${TIDETS_HOST:-127.0.0.1}"
PORT="${TIDETS_PORT:-5556}"
METRICS_URL="${TIDETS_METRICS_URL:-http://127.0.0.1:9090/metrics}"
RESULT_DIR="${TIDETS_BENCH_RESULT_DIR:-./bench-results}"

mkdir -p "${RESULT_DIR}"

echo "Running TideTS benchmark against ${HOST}:${PORT}"
go run ./scripts/bench \
  -host "${HOST}" \
  -port "${PORT}" \
  -op "${1:-insert_batch}" \
  -points "${POINTS:-10000}" \
  -batch-size "${BATCH_SIZE:-100}" \
  -range-size "${RANGE_SIZE:-1000}" \
  -concurrency "${CONCURRENCY:-1}" \
  -warmup "${WARMUP:-100}" \
  -metrics "${METRICS_URL}" \
  -result-dir "${RESULT_DIR}"
