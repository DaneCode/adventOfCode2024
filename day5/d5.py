def parse_input(input_str):
    rules_str, updates_str = input_str.strip().split('\n\n')
    rules = [tuple(map(int, rule.split('|'))) for rule in rules_str.split('\n')]
    updates = [list(map(int, update.split(','))) for update in updates_str.split('\n')]
    return rules, updates

def is_correct_order(update, rules):
    index_map = {page: i for i, page in enumerate(update)}
    for x, y in rules:
        if x in index_map and y in index_map and index_map[x] > index_map[y]:
            return False
    return True

def find_middle_page(update):
    return update[len(update) // 2]

def sum_middle_pages(input_str):
    rules, updates = parse_input(input_str)
    total = 0
    for update in updates:
        if is_correct_order(update, rules):
            total += find_middle_page(update)
    return total

# Read input from file
with open('d5input.txt', 'r') as file:
    input_str = file.read()

# Calculate the sum of middle pages
result = sum_middle_pages(input_str)
print(result)