valor_dolar = float(input("Digite o valor em dólares (US$): "))
cotacao_dolar = float(input("Digite a cotação do dólar: "))

valor_real = valor_dolar * cotacao_dolar
print(f"O valor em reais (R$) é: {valor_real:.2f}")
