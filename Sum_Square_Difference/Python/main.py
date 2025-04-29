def sum_of_arithmetic_series(n:int, a1:int, r:int):
    an = a1 + (n - 1) * r
    sn = (n * (a1 + an)) // 2
    return sn

def sum_of_square_arithmetic_series(n:int, a1:int, r:int):
    an = a1 + (n - 1) * r
    sn = 0
    for i in range(a1,an+1,r):
        sn += i * i
    return sn

res_1 = sum_of_arithmetic_series(100,1,1)**2
res_2 = sum_of_square_arithmetic_series(100,1,1)
print(res_1 - res_2)