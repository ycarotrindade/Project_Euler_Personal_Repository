def identify_palindrome(n: int):
    str_reversed = "".join(reversed(str(n)))
    if str(n) == str_reversed:
        return True
    else:
        return False

largest_palindrome = 0
for i in range(999,100,-1):
    for j in range(999,100,-1):
        num = i * j
        if identify_palindrome(num) and num > largest_palindrome:
            largest_palindrome = num
print(largest_palindrome)