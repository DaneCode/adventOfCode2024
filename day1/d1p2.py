from collections import Counter

# Step 1: Read the data from d1Input.txt
with open('d1Input.txt', 'r') as file:
    lines = file.readlines()

# Step 2: Split each line into two numbers and store them in two separate arrays
array1 = []
array2 = []
for line in lines:
    num1, num2 = map(int, line.split())
    array1.append(num1)
    array2.append(num2)

# Step 3: Count the occurrences of each number from the first array in the second array
counter = Counter(array2)

# Step 4: Multiply each number from the first array by its count in the second array and store the result in a third array
array3 = [num * counter[num] for num in array1]

# Step 5: Sum all the values in the third array to get the total
total = sum(array3)

print("Total:", total)