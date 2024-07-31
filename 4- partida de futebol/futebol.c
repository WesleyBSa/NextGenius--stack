#include <stdio.h>

int main() {
    int golsTime1, golsTime2;
    printf("Digite os gols do time 1: ");
    scanf("%d", &golsTime1);
    printf("Digite os gols do time 2: ");
    scanf("%d", &golsTime2);

    if (golsTime1 > golsTime2) {
        printf("Time 1 venceu!\n");
    } else if (golsTime2 > golsTime1) {
        printf("Time 2 venceu!\n");
    } else {
        printf("Empate!\n");
    }
    return 0;
}
