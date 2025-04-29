def list_mult(n:int):
    mult = [x for x in range(1,n) if x % 3 == 0 or x % 5 == 0]
    return sum(mult)

if __name__ == '__main__':
    n = int(input('Type the max number:'))
    res = list_mult(n)
    print(f'The sum of the numbers below \033[32m{n}\033[0m that are multiples of 3 or 5 is \033[32m{res}\033[0m ')