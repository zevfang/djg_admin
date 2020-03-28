# Basic go commands
# #############/Services/RTSS/layui/RTSS/bin/
GOBUILD-PATH=../../public/bin
GOBUILD=go build
GOCLEAN=go clean
CROSS-LINUX=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Binary names
bin=\
rtss_local-linux \

# Build
all: $(bin)
default: $(bin)

rtss_local-linux:
	$(CROSS-LINUX) $(GOBUILD) -o $(GOBUILD-PATH)/rtss_local ./main.go
	mkdir -p $(GOBUILD-PATH)/static/css/ $(GOBUILD-PATH)/static/js/ $(GOBUILD-PATH)/views/
	cp -rf ./static/css/ $(GOBUILD-PATH)/static/css/
	cp -rf ./static/js/ $(GOBUILD-PATH)/static/js/
	cp -rf ./views/ $(GOBUILD-PATH)/views/

# Clean
clean:
	$(GOCLEAN)
	rm -f $(GOBUILD-PATH)/rtss_local