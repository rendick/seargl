.PHONY: all

all: build 

build:
	go build && echo "Success: Go application built successfully." || (echo "Error: Failed to build the Go application." >&2; exit 1)

.PHONY: clean

clean:
	go clean

.PHONY: run

run: all
	@echo "Success: Script completed successfully."
