TARGET=hello
UF2=$(TARGET).uf2
SRCS=hello.go

all: $(UF2)


$(UF2): $(TARGET)
	tinygo build -target pico -o $(UF2) hello.go

$(TARGET): $(SRCS)
	tinygo build $^

mac: $(SRCS)
	go build $^

clean:
	rm -rf $(TARGET) $(UF2)
