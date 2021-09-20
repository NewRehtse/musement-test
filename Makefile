# May be useful if you want to test with a different Go version
# (note that variables can also be overwritten in cli argument)
DEP_BIN = $(shell which dep)
GO_BIN = $(shell which go)
GO_BIN = $(shell which go)
GO_BIN_FOLDER = $(shell echo $(GOPATH)/bin)
MKDIR = mkdir -p

.PHONY: list
list:
	@echo ""
	@echo "Useful targets:"
	@echo ""
	@echo "  deps         > install dependencies"
	@echo "  build        > compile package"
	@echo "  clean-build  > clean && compile package"
	@echo "  install      > installs application"
	@echo ""
	@echo "  clean        > removes vendors and built package"
	@echo ""
	@echo "  test        > run test"
	@echo "  bench        > run benchmarks"
	@echo ""

#
# Setting features
#----------------------------------------------------------------------------

# deps dependencies with Glide
.PHONY: deps
deps:
	$(GO_BIN) mod tidy
	$(GO_BIN) mod vendor

# compile the package
.PHONY: build
build: compile

# cleans and compile the package
.PHONY: clean-build
clean-build: clean deps compile

.PHONY: compile
compile:
	$(RM) ./bin/app
	$(GO_BIN) build -o ./bin/app

.PHONY: install
install: compile copy_to_bin

.PHONY: copy_to_bin
copy_to_bin:
	cp ./bin/app $(GO_BIN_FOLDER)

#
# Testing features
#------------------------------------------------

.PHONY: test
test:
	@echo "\n\033[92m***** RUNNING TESTS *****\033[0m"
	@echo "\n\033[92m***** CLEAN SCENARIO *****\033[0m"
	$(RM) -r build
	$(MKDIR) build
	@echo "\n\033[92m***** RUN TESTS *****\033[0m"
	$(GO_BIN) test -race -coverprofile=build/coverage.out ./...
	@echo "\n\033[92m***** BUILD REPORT *****\033[0m"
	$(GO_BIN) tool cover -html=build/coverage.out -o build/coverage.html

.PHONY: bench
bench:
	@echo "\n\033[92m***** RUNNING TESTS *****\033[0m"
	$(GO_BIN) test -bench=. ./...

#
# Cleaning features
#------------------------------------------------

.PHONY: clean
clean:
	$(RM) -r vendor
