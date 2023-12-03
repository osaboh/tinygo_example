TARGET=advertise
HEX=$(TARGET).hex
SRCS=main.go

.PHONY: mod

all: mod $(HEX)

$(HEX): $(TARGET).elf
	tinygo build -target=microbit-s110v8 -o=./$(HEX) $(SRCS)

$(TARGET).elf: $(SRCS)
	tinygo build -target=microbit-s110v8 $^

mac: $(SRCS)
	go build $^

flash:
	tinygo flash -target=microbit-s110v8 -programmer=cmsis-dap

mod: go.mod go.sum
go.mod:
	go mod init $(TARGET)
go.sum:
	go mod tidy

clean:
	rm -rf $(TARGET) $(HEX) go.mod go.sum *.elf
