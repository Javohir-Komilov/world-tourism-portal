// utils.c
#include <stdio.h>
#include "utils.h"

void clear_screen() {
    printf("\033[H\033[J"); // ANSI escape code to clear screen
}

void wait_for_keypress() {
    printf("Press Enter to continue...");
    while (getchar() != '\n');
}
