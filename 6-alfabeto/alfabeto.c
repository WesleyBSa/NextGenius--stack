#include <stdio.h>
#include <ctype.h>

int main() {
    char letra;
    printf("Digite uma letra: ");
    scanf(" %c", &letra);

    letra = tolower(letra);
    if (!isalpha(letra)) {
        printf("Caractere não aceito no programa.\n");
    } else if (letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u') {
        printf("A letra é uma vogal.\n");
    } else {
        printf("A letra é uma consoante.\n");
    }

    return 0;
}
