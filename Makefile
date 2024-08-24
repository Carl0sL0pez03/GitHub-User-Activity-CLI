# Nombre del binario
BINARY_NAME = github-activity

# Directorio donde se genera el binario
BIN_DIR = ./bin

# Variable for the Go executable
GO = go

# Comando para construir el binario
build:
	@echo "Building the application..."
	$(GO) build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_SRC)
	@echo "Build complete."

# Comando para ejecutar el binario
run: build
	@echo "Running the application..."
	$(BIN_DIR)/$(BINARY_NAME) $(ARGS)

# Command to run tests
test:
	$(GO) test ./...

# Comando para limpiar los archivos generados
clean:
	@echo "Cleaning up..."
	rm -f $(BIN_DIR)/$(BINARY_NAME)
	@echo "Cleanup complete."

# Comando por defecto
all: build

# Ejemplo de uso de argumentos
run-example:
	@$(MAKE) run ARGS="kamranahmedse"

# Help
help:
	@echo "Makefile commands:"
	@echo "  build        - Build the Go application"
	@echo "  run ARGS=... - Run the application with specified arguments"
	@echo "  clean        - Remove the built binary"
	@echo "  all          - Build the application (default)"
	@echo "  run-example  - Run the application with an example argument (username: kamranahmedse)"
	@echo "  help         - Show this help message"

.PHONY: build run clean all run-example help