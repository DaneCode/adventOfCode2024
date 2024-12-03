def is_safe(report):
    increasing = all(report[i] < report[i + 1] and 1 <= report[i + 1] - report[i] <= 3 for i in range(len(report) - 1))
    decreasing = all(report[i] > report[i + 1] and 1 <= report[i] - report[i + 1] <= 3 for i in range(len(report) - 1))
    return increasing or decreasing

# Step 1: Read the data from d2input.txt
with open('d2input.txt', 'r') as file:
    lines = file.readlines()

# Step 2: Split each line into a list of numbers
reports = [list(map(int, line.split())) for line in lines]

# Step 3-5: Check each report and count the number of safe reports
safe_count = sum(1 for report in reports if is_safe(report))

print("Number of safe reports:", safe_count)