##############################################
#
# Change this section to customize the output
# for your software
#
# MAIN SETTINGS

BINARY=automate
VERSION=1.0.0
CONTROL_VERSION=$(VERSION)
USERNAME=tbleuse
BUILD_TAGS= -tags "mastervtom"
##############################################

#GO SETTINGS
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOARCH=$(shell $(GOCMD) env GOARCH)

#SOFTWARE SETTINGS
COMPILE_DATE=$(shell date +"%d/%m/%Y - %H:%M")
LDFLAGS= -ldflags '-X "kdatapack/utils.VERSION_NUMBER=$(VERSION)" -X "kdatapack/utils.COMPILE_DATE=$(COMPILE_DATE)" -X "kdatapack/utils.USERNAME=$(USERNAME)"'
BINARY_NAME=$(BINARY)-$(VERSION)

all: deps build

build:
	@echo "> Building..."
	$(GOBUILD) -o $(BINARY_NAME) $(LDFLAGS) $(BUILD_TAGS)
	@echo ""

clean:
	@echo "> Cleaning folders..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -fr $(BINARY_NAME)
	@echo ""

fclean: clean
	@echo "> Cleaning potential binaries..."
	rm -rf $(BINARY)-[0-9].*
	@echo ""

deps:
	@echo "> Go getting project dependencies..."
	$(GOGET) $(BUILD_TAGS)
	@echo ""

re: clean all

.PHONY: all build clean deps re
