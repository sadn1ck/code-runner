GO := go
DEV := run
BUILD := build
FILE := main.go
EXEC := judge

all:
	$(GO) $(DEV) $(FILE)

build:
	$(GO) $(BUILD) -o ./$(BUILD)/$(EXEC)