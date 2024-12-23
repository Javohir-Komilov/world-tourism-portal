#include <stdio.h>
#include "ui.h"
#include "socket_client.h"

int main() {
    if (initialize_socket_client() != 0) {
        fprintf(stderr, "Failed to initialize client.\n");
        return 1;
    }

    show_main_menu();
    cleanup_socket_client();
    return 0;
}