letra = input("Digite uma letra: ").lower()

if not letra.isalpha() or len(letra) != 1:
    print("Caractere não aceito no programa.")
elif letra in "aeiou":
    print("A letra é uma vogal.")
else:
    print("A letra é uma consoante.")
