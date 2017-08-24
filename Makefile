cnbm_co_version := 0.1.0

.PHONY: build test

###############################################################################
# commands related to Go testing and builds --> binaries (HTTP API and CLI tool)

build :
	@echo Building the CNBM CO CLI
	go build -ldflags "-X github.com/cnbm/container-orchestration/cli/cmd.releaseVersion=$(cnbm_co_version)" -o ./cnbm-co cli/main.go

test :
	@echo Testing the CNBM CO library
	go test -short -run Test* ./pkg
