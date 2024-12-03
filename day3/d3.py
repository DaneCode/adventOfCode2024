import re

# Function to extract and multiply numbers
def extract_and_multiply(filename):
    results = []
    # Updated pattern to strictly match valid mul(X,Y) instructions
    pattern = re.compile(r'mul\((\d{1,3}),(\d{1,3})\)')
    
    with open(filename, 'r') as file:
        for line in file:
            matches = pattern.findall(line)
            for match in matches:
                n1, n2 = map(int, match)
                results.append(n1 * n2)
    
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