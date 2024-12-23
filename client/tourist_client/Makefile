CC = gcc
CFLAGS = -Wall -Wextra -std=c11
SRC = src/main.c src/ui.c src/socket_client.c src/utils.c
INC = include
BIN = tourist_client

all: $(BIN)

$(BIN): $(SRC)
	$(CC) $(CFLAGS) -I$(INC) -o $@ $(SRC)

clean:
	rm -f $(BIN)

.PHONY: all clean
