SHELL       = /bin/bash -o pipefail
PACKAGE     = github.com/lopezator/nexmo-go
NAME        =  $(shell echo $(PACKAGE) | rev | cut -d/ -f1 | rev)
PKGS        =  $(shell go list ./... | grep -v /vendor/)
REPORTS_DIR ?= .reports

.PHONY: prepare
prepare:
	go get -u github.com/jstemmer/go-junit-report
	go get -u github.com/golang/lint/golint
	go get -u honnef.co/go/tools/cmd/megacheck
	go get -u github.com/kisielk/errcheck
	go get -u golang.org/x/tools/cmd/goimports

.PHONY: clean
clean:
	rm -f ./bin/*
	@for app in $(APP) ; do \
		rm -f cmd/$$app/*-linux-amd64; \
	done

.PHONY: install
install:
	@echo "Installing nexmo-go dependencies into your GOPATH..."
	@go get ./...

.PHONY: sanity-check
sanity-check: goimports golint vet megacheck errcheck

.PHONY: test
test:
	@echo "Running unit tests..."
	mkdir -p $(REPORTS_DIR)
	2>&1 go test -race -cover -v $(shell echo $(PKGS) | tr " " "\n") | tee $(REPORTS_DIR)/report-unittests.out
	cat $(REPORTS_DIR)/report-unittests.out | go-junit-report -set-exit-code > $(REPORTS_DIR)/report-unittests.xml

.PHONY: goimports
goimports:
	@echo "Running goimports..."
	@test -z "`for pkg in $(PKGS); do goimports -l $(GOPATH)/src/$$pkg/*.go | tee /dev/stderr; done`"

.PHONY: golint
golint:
	@echo "Running golint..."
	@golint -set_exit_status $(PKGS)

.PHONY: vet
vet:
	@echo "Running go vet..."
	@go vet $(PKGS)

.PHONY: megacheck
megacheck:
	@echo "Running megacheck..."
	@megacheck $(PKGS)

.PHONY: errcheck
errcheck:
	@echo "Running errcheck..."
	@errcheck $(PKGS)
