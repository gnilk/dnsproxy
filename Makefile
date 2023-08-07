BIN = $(GOPATH)/bin
MODELGEN = $(BIN)/modelgenerator
GOIMPORTS = $(BIN)/goimports

NODE_BIN = $(shell npm bin)

TEST_FILES = \
	util_test.go \
	rulesengine_test.go \
	router_test.go

TEST_ROUTER_FILES = router_test.go
TEST_RULES_FILES = rulesengine_test.go

SRC_FILES = \
	dnsproxy.go \
	perflog.go \
	config.go \
	rulesengine.go \
	routerclient.go \
	unifirouterclient.go \
	devicecache.go \
	system.go \
	resolver.go \
	util.go

MODEL_SRC = config_datamodel.xml
MODEL_OUT = config.go

VERSION=0

all:	proxy install

proxy: $(SRC_FILES)
	go build -o dnsproxy $(SRC_FILES)

install: proxy
	cp dnsproxy $(BIN)



test: $(SRC_FILES) $(TEST_FILES)
	go test -v $(TEST_FILES) $(SRC_FILES)

test_router: $(SRC_FILES) $(TEST_ROUTER_FILES)
	go test -v $(TEST_ROUTER_FILES) $(SRC_FILES)

test_rules: $(SRC_FILES) $(TEST_RULES_FILES)
	go test -v $(TEST_RULES_FILES) $(SRC_FILES)


persistence: config_datamodel.xml
	$(MODELGEN) -v -c config_datamodel.xml -o config.go
	$(GOIMPORTS) -w $(MODEL_OUT)

