import random

def if_else():
    i = random.randint(10,50)
    if i == 45:
        print('i is 45')
    elif i == 40:
        print('i is 40')
    else:
        print(f"i is {i}")

if_else()

def for_loop():
    for i in range(20):
        print(i)

for_loop()

def while_loop():
    count = 0
    while count < 3:
        print(f"The count is {count}")
        count += 1
while_loop()