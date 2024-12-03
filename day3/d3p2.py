import re

# Function to extract and multiply numbers
def extract_and_multiply(filename):
    results = []
    # Updated pattern to strictly match valid mul(X,Y) instructions
    mul_pattern = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')
    do_pattern = re.compile(r'do\(\)')
    dont_pattern = re.compile(r"don't\(\)")

    with open(filename, 'r') as file:
        data = file.read()
        mul_enabled = True  # Initially, mul instructions are enabled

        i = 0
        while i < len(data):
            # Check for do() and don't() instructions
            if data[i:].startswith('do()'):
                mul_enabled = True
                i += 4  # Move index past the do() instruction
            elif data[i:].startswith("don't()"):
                mul_enabled = False
                i += 7  # Move index past the don't() instruction

            # Check for mul(X,Y) instructions
            if mul_enabled:
                match = mul_pattern.match(data[i:])
                if match:
                    n1, n2 = map(int, match.groups())
                    results.append(n1 * n2)
                    i += match.end()  # Move index past the mul(X,Y) instruction
                else:
                    i += 1
            else:
                i += 1

    return results

# File path
file_path = 'd3.txt'

# Get the results
multiplication_results = extract_and_multiply(file_path)

# Calculate the sum of the results
total_sum = sum(multiplication_results)

# Print the sum of the results
print('Sum of multiplication results:')
print(total_sum)