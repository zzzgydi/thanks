HOMEDIR := $(shell pwd)
PROJECTNAME := $(shell basename $(HOMEDIR))
OUTDIR  := $(HOMEDIR)/output

GOBUILD := go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"

build:
	mkdir -p $(OUTDIR)
	$(GOBUILD) -o $(OUTDIR)/$(PROJECTNAME) ./cmd/main.go


dev: build
	mkdir -p $(OUTDIR)/log
	@if [ -f .env ]; then cp .env $(OUTDIR)/.env; fi
	cd $(OUTDIR) && ./$(PROJECTNAME)


include .env
export $(shell sed 's/=.*//' .env)

.PHONY: test
test:
	go test -cover ./service/parser
	go test -cover ./service/thk
