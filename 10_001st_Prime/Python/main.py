def sieve_of_eratosthenes(n:int):
    sieve = [True] * (n + 1)
    primes = []
    p = 2
    while p * p <= n:
        if sieve[p]:
            for i in range(p * p,n + 1,p):
                sieve[i] = False
        p += 1
    
    for i, v in enumerate(sieve):
        if v:
            primes.append(i)
    return primes[2:]

print(sieve_of_eratosthenes(100000000)[10000])