#!/bin/bash

error_exit() {
  echo "Error: $1" >&2
  exit 1
}

success_message() {
  echo "Success: $1"
}

go build && success_message "Go application built successfully." || error_exit "Failed to build the Go application."

echo "Enter your default browser: "
read var1

if [ -z "$var1" ]; then
  error_exit "Browser command cannot be empty."
fi

$var1 localhost:8000 && success_message "Browser opened successfully." || error_exit "Failed to open the browser with the specified URL."

./seargl && success_message "seargl executed successfully." || error_exit "Failed to execute seargl."

success_message "Script completed successfully."
