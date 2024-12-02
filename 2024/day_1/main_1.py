left_list = []
right_list = []
while True:
    values = input().split()
    if not values:
        break
    left_list.append(int(values[0]))
    right_list.append(int(values[1]))

print(left_list)
print(right_list)
