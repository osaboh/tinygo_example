TARGET=microbit-blink
HEX=$(TARGET).hex
SRCS=microbit-blink.go

all: $(HEX)


$(HEX): $(TARGET)
	tinygo build -target=microbit -o=./$(HEX) $(SRCS)

$(TARGET): $(SRCS)
	tinygo build -target=microbit $^

mac: $(SRCS)
	go build $^

clean:
	rm -rf $(TARGET) $(HEX) *.elf
