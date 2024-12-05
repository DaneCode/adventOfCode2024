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
    if not update:
        raise ValueError("Update list is empty")
    return update[len(update) // 2]

def topological_sort(pages, rules):
    from collections import defaultdict, deque

    graph = defaultdict(list)
    in_degree = defaultdict(int)
    for x, y in rules:
        if x in pages and y in pages:
            graph[x].append(y)
            in_degree[y] += 1
            if x not in in_degree:
                in_degree[x] = 0

    queue = deque([page for page in pages if in_degree[page] == 0])
    sorted_pages = []

    while queue:
        node = queue.popleft()
        sorted_pages.append(node)
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    if len(sorted_pages) != len(pages):
        raise ValueError("Cycle detected or missing pages in topological sort")

    return sorted_pages

def sum_middle_pages_of_corrected_updates(input_str):
    rules, updates = parse_input(input_str)
    total = 0
    for update in updates:
        if not is_correct_order(update, rules):
            sorted_update = topological_sort(update, rules)
            print(f"Original update: {update}")
            print(f"Sorted update: {sorted_update}")
            total += find_middle_page(sorted_update)
    return total

# Read input from file
with open('d5input.txt', 'r') as file:
    input_str = file.read()

# Calculate the sum of middle pages of corrected updates
result = sum_middle_pages_of_corrected_updates(input_str)
print(result)