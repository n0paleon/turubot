# Nama file output
BINARY_NAME := ../bin/turubot
# Nama file utama
MAIN_FILE := ../cmd/bot/main.go

# Variabel Go
GO := go
GO_BUILD := $(GO) build
GO_FMT := $(GO) fmt
GO_MOD := $(GO) mod
GO_RUN := $(GO) run

.PHONY: all install run clean

# Compile dan build binary
all: install build

# Install dependencies
install:
	$(GO_MOD) tidy

# Format kode
fmt:
	$(GO_FMT) ../...

# Build bot
build: fmt
	$(GO_BUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Jalankan bot
run:
	$(GO_RUN) $(MAIN_FILE)

# Hapus file binary
clean:
	rm -f $(BINARY_NAME)
