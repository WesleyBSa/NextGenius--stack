#include <stdio.h>

int main() {
    float valorDolar, cotacaoDolar, valorReal;
    printf("Digite o valor em dólares (US$): ");
    scanf("%f", &valorDolar);
    printf("Digite a cotação do dólar: ");
    scanf("%f", &cotacaoDolar);

    valorReal = valorDolar * cotacaoDolar;
    printf("O valor em reais (R$) é: %.2f\n", valorReal);
    return 0;
}
