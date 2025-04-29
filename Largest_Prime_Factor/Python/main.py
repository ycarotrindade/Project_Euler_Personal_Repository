def sieve_of_eratosthenes(n:int):
    sieve = [True] * (n + 1)
    primes = []
    p = 2
    while p * p <= n:
        if sieve[p]:
            for i in range(p * p, n + 1, p):
                sieve[i] = False
        p += 1
    
    for i, v in enumerate(sieve):
        if v == True:
            primes.append(i)
            
    return primes[2:]


max_number = int(input("Type max number:"))
primes = sieve_of_eratosthenes(10000)
aux = 0
while max_number > 1:
    prime = primes[aux]
    if max_number % prime == 0:
        max_number /= prime
    else:
        aux += 1
    
    if aux >= len(primes):
        break
print(max_number, primes[aux])