.PHONY: all

all: build execute_seargl

build:
	go build && echo "Success: Go application built successfully." || (echo "Error: Failed to build the Go application." >&2; exit 1)

execute_seargl:
	@output=$$(./seargl --help2>&1); \
	if echo "$$output" | grep -q "Seargl executed."; then \
		echo "Success: seargl executed successfully."; \
	else \
		echo "Error: Failed to execute seargl. Output: $$output" >&2; \
		exit 1; \
	fi

.PHONY: clean

clean:
	# Add clean-up commands if needed

.PHONY: run

run: all
	@echo "Success: Script completed successfully."
