# Paths (relative to server/)
RUSTLIB_DIR = bindings/rustlib
MERKLE_DIR = merkle
SHARED_DIR = $(MERKLE_DIR)/shared
LIB_NAME = librustlib.dylib

.PHONY: all build clean run

all: run

build:
	@echo "Building Rust library..."
	cd $(RUSTLIB_DIR) && cargo build --release
	cp $(RUSTLIB_DIR)/target/release/$(LIB_NAME) $(SHARED_DIR)

clean:
	@echo "Clean Build"
	cd $(RUSTLIB_DIR) && cargo clean
	$(MAKE) build

run:
	@echo "Running Go program..."
	cd $(MERKLE_DIR) && \
	ls ./shared && \
	install_name_tool -id @rpath/$(LIB_NAME) ./shared/$(LIB_NAME) && \
	DYLD_LIBRARY_PATH=./shared go run main.go

