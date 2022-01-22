GC=go build
BUILD_DIR=build

all:build


build:
	$(GC) -o $(BUILD_DIR)/yaf

install: build
	@echo "Installing yaf"
	@cp $(BUILD_DIR)/yaf /usr/bin/yaf
	@echo "Setting permissions"
	@chmod go+x /usr/bin/yaf

clean:
	rm -rf $(BUILD_DIR)