is_divisible_by_all = False
num = 21
while not is_divisible_by_all:
    is_divisible_by_all = True
    for i in range(2,21):
        if not num % i == 0:
            is_divisible_by_all = False
            num += 1
            break
print(num)