# Basic go commands

GOBUILD-PATH=./docker/deploy/cmd
GOBUILD=go build
GOCLEAN=go clean
CROSS-LINUX=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Binary names
bin=djg_api-linux

# Build
all: $(bin)
default: $(bin)

djg_api-linux:
	$(CROSS-LINUX) $(GOBUILD) -o $(GOBUILD-PATH)/djg_admin ./main.go
	cp -rf ./conf/ $(GOBUILD-PATH)/conf/
	cp -rf ./static/ $(GOBUILD-PATH)/static/
	cp -rf ./views/ $(GOBUILD-PATH)/views/

# Clean
clean:
	$(GOCLEAN)
	rm -rf $(GOBUILD-PATH)/*