#include <stdio.h>

int main() {
    float a, b;
    printf("Digite dois números: ");
    scanf("%f %f", &a, &b);
    printf("Soma: %.2f\n", a + b);
    printf("Subtração: %.2f\n", a - b);
    printf("Multiplicação: %.2f\n", a * b);
    printf("Divisão: %.2f\n", a / b);
    return 0;
}
