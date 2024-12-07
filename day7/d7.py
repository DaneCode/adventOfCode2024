import itertools

def evaluate_expression(numbers, operators):
    result = numbers[0]
    for i in range(len(operators)):
        if operators[i] == '+':
            result += numbers[i + 1]
        elif operators[i] == '*':
            result *= numbers[i + 1]
    return result

def can_be_true(test_value, numbers):
    if len(numbers) == 1:
        return numbers[0] == test_value
    
    operators = ['+', '*']
    for ops in itertools.product(operators, repeat=len(numbers) - 1):
        if evaluate_expression(numbers, ops) == test_value:
            return True
    return False

def main():
    total_calibration_result = 0
    
    with open('d7Input.txt', 'r') as file:
        for line in file:
            test_value_str, numbers_str = line.split(':')
            test_value = int(test_value_str.strip())
            numbers = list(map(int, numbers_str.strip().split()))
            
            if can_be_true(test_value, numbers):
                total_calibration_result += test_value
    
    print(total_calibration_result)

if __name__ == "__main__":
    main()