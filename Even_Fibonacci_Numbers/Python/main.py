def sum_even_fib(max:int):
    last_numbers = [1,2]
    sum_fib = 2
    while True:
        fib_num = sum(last_numbers)
        if fib_num <= max:
            last_numbers[0] = last_numbers[1]
            last_numbers[1] = fib_num
            sum_fib += fib_num if fib_num % 2 == 0 else 0
        else:
            return sum_fib

if __name__ == '__main__':
    max_target = int(input('Type max fibonacci value:'))
    sum_fib = sum_even_fib(max_target)
    print(f'The sum the even fibonacci numbers below \033[32m{max_target}\033[0m is \033[32m{sum_fib}\033[0m')