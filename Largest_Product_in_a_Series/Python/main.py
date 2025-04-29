with open('../number.txt','r') as file:
    text = "".join(file.readlines()).replace("\n","")

max_length = 13
largest_product = 0


for i in range(0, len(text) - max_length):
    numbers = list(text[i:i + max_length])
    product = 1
    for v in numbers:
        product *= int(v)
    largest_product = product if product > largest_product else largest_product
print(largest_product)