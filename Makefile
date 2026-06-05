export PATH := $(shell go env GOPATH)/bin:$(PATH)

PROTO_DIR := protocol/grpc-datanode
PB_DIR := $(PROTO_DIR)/pb

ANTLR_VERSION := 4.13.2
ANTLR_JAR := tools/antlr/antlr-$(ANTLR_VERSION)-complete.jar
ANTLR_URL := https://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar
GRAMMAR_DIR := antlr/grammar
PARSER_OUT := antlr/parser

.PHONY: antlr-download sql-gen
antlr-download:
	@mkdir -p tools/antlr
	@test -f $(ANTLR_JAR) || ( \
		echo "downloading $(ANTLR_JAR) from $(ANTLR_URL)..." && \
		curl -fsSL -o $(ANTLR_JAR) $(ANTLR_URL) \
	)

sql-gen: antlr-download
	@command -v java >/dev/null 2>&1 || (echo "JDK required: install java and retry make sql-gen" && exit 1)
	cd $(GRAMMAR_DIR) && java -jar ../../$(ANTLR_JAR) -Dlanguage=Go -visitor -package parser -o ../parser TideSQL.g4

.PHONY: proto
proto:
	@mkdir -p $(PB_DIR)
	protoc \
		-I $(PROTO_DIR) \
		--go_out=$(PB_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/client.proto

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ./cmd/datanode -data-dir ./data

.PHONY: build
build:
	@mkdir -p bin
	go build -o bin/datanode ./cmd/datanode
	go build -o bin/tidets-cli ./cmd/tidets-cli

.PHONY: build-cli
build-cli:
	@mkdir -p bin
	go build -o bin/tidets-cli ./cmd/tidets-cli

.PHONY: bench-help
bench-help:
	@echo 'Start DataNode first, then run one of:'
	@echo '  go run ./scripts/bench -op insert_batch -points 10000 -batch-size 100 -concurrency 4'
	@echo '  ./scripts/run_bench.sh insert_batch'
