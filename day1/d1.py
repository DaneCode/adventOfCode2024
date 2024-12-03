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

# Step 3: Sort both arrays
array1.sort()
array2.sort()

# Step 4: Calculate the absolute difference between corresponding elements of the two arrays
array3 = [abs(a - b) for a, b in zip(array1, array2)]

# Step 5: Sum all the values in the third array
total_difference = sum(array3)

print("Total difference:", total_difference)