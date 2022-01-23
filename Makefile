GC=go build
BUILD_DIR=build
DESTDIR=/usr/bin

all:build


build:
	$(GC) -o $(BUILD_DIR)/yaf

install: build
	@echo "Installing yaf"
	@cp $(BUILD_DIR)/yaf $(DESTDIR)/yaf
	@echo "Setting permissions"
	@chmod go+x $(DESTDIR)/yaf

clean:
	rm -rf $(BUILD_DIR)
