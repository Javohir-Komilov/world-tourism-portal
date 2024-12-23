// socket_client.h
#ifndef SOCKET_CLIENT_H
#define SOCKET_CLIENT_H

int initialize_socket_client();
void cleanup_socket_client();
int send_message_to_server(const char *message, char *response);

#endif // SOCKET_CLIENT_H
