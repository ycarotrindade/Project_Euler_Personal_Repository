def generate_pythagorean(n1:int, n2:int):
    a = n1**2 - n2**2
    b = 2 * n1 * n2
    c = n1**2 + n2**2
    return a, b, c

m = 1
is_found = False
while not is_found:
    for n in range(m):
        numbers = generate_pythagorean(m, n)
        if sum(numbers) == 1000:
            prod = numbers[0] * numbers[1] * numbers[2]
            print(prod)
            is_found = True
            break
    m += 1