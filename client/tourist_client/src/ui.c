// ui.c
#include <stdio.h>
#include "ui.h"
#include "socket_client.h"

void show_main_menu() {
    int choice;
    do {
        printf("\n=== Tourist Client ===\n");
        printf("1. Login\n");
        printf("2. Register\n");
        printf("3. View Packages\n");
        printf("4. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                // Call login function
                break;
            case 2:
                // Call registration function
                break;
            case 3:
                // Call view packages function
                break;
            case 4:
                printf("Exiting...\n");
                break;
            default:
                printf("Invalid choice. Try again.\n");
        }
    } while (choice != 4);
}
