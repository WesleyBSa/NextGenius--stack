gols_time1 = int(input("Digite os gols do time 1: "))
gols_time2 = int(input("Digite os gols do time 2: "))
if gols_time1 > gols_time2:
    print("Time 1 venceu!")
elif gols_time2 > gols_time1:
    print("Time 2 venceu!")
else:
    print("Empate!")
