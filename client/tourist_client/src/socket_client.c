// socket_client.c
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <arpa/inet.h>
#include "socket_client.h"

#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 8080

static int sock;

int initialize_socket_client() {
    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1) {
        perror("Socket creation failed");
        return -1;
    }

    struct sockaddr_in server_addr;
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(SERVER_PORT);

    if (inet_pton(AF_INET, SERVER_IP, &server_addr.sin_addr) <= 0) {
        perror("Invalid address");
        return -1;
    }

    if (connect(sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        perror("Connection failed");
        return -1;
    }

    printf("Connected to server.\n");
    return 0;
}

void cleanup_socket_client() {
    close(sock);
    printf("Disconnected from server.\n");
}

int send_message_to_server(const char *message, char *response) {
    if (send(sock, message, strlen(message), 0) < 0) {
        perror("Send failed");
        return -1;
    }

    if (recv(sock, response, 1024, 0) < 0) {
        perror("Receive failed");
        return -1;
    }

    return 0;
}
