from dataclasses import dataclass, field
from typing import Callable


SAMPLE = """Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1"""

PUZZLE="""Monkey 0:
  Starting items: 64
  Operation: new = old * 7
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 1:
  Starting items: 60, 84, 84, 65
  Operation: new = old + 7
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 7

Monkey 2:
  Starting items: 52, 67, 74, 88, 51, 61
  Operation: new = old * 3
  Test: divisible by 5
    If true: throw to monkey 5
    If false: throw to monkey 7

Monkey 3:
  Starting items: 67, 72
  Operation: new = old + 3
  Test: divisible by 2
    If true: throw to monkey 1
    If false: throw to monkey 2

Monkey 4:
  Starting items: 80, 79, 58, 77, 68, 74, 98, 64
  Operation: new = old * old
  Test: divisible by 17
    If true: throw to monkey 6
    If false: throw to monkey 0

Monkey 5:
  Starting items: 62, 53, 61, 89, 86
  Operation: new = old + 8
  Test: divisible by 11
    If true: throw to monkey 4
    If false: throw to monkey 6

Monkey 6:
  Starting items: 86, 89, 82
  Operation: new = old + 2
  Test: divisible by 7
    If true: throw to monkey 3
    If false: throw to monkey 0

Monkey 7:
  Starting items: 92, 81, 70, 96, 69, 84, 83
  Operation: new = old + 4
  Test: divisible by 3
    If true: throw to monkey 4
    If false: throw to monkey 5"""


data = PUZZLE

PART1 = 20
PART2 = 10000


def get_indentation_level(line: str):
    return len(line) - len(line.lstrip(" "))

@dataclass
class Monkey:
    id: int = field(default=-1)
    items: list[int] = field(default_factory=list)
    operation: Callable[[int], int] = field(default=lambda x: x)
    test_modulus: int = field(default=0)
    next_monkey: dict[bool, int] = field(default_factory=dict)
    inspections: int = field(default=0)
    operation_is_multiplication: bool = field(default=False)


monkeys = []
instructions = []

    
current_monkey = None
for line in data.splitlines():
    indent = get_indentation_level(line)
    indent = indent // 2
    if indent == 0:
        if len(line) == 0:
            monkeys.append(current_monkey)
            current_monkey = None
        elif line.startswith("Monkey"):
            current_monkey = Monkey()
            monkey_id = int(line.removeprefix("Monkey ").removesuffix(":"))
            current_monkey.id = monkey_id
        else:
            raise ValueError(f"Invalid line: {line}")
    elif indent == 1:
        line_stripped = line.lstrip(" ")
        assert current_monkey is not None
        if line_stripped.startswith("Starting items:"):
            line_stripped = line_stripped.removeprefix("Starting items:")
            items = [int(x) for x in line_stripped.split(",")]
            current_monkey.items = items
        elif line_stripped.startswith("Operation:"):
            line_stripped = line_stripped.removeprefix("Operation: ")
            line_stripped = line_stripped.strip()
            instruction = f"""def instruction_monkey_{current_monkey.id}(old):
    {line_stripped}
    return new

instructions.append(instruction_monkey_{current_monkey.id})"""
            exec(instruction)
            current_monkey.operation = instructions[-1]
            current_monkey.operation_is_multiplication = "*" in line_stripped
        elif line_stripped.startswith("Test: divisible by "):
            line_stripped = line_stripped.removeprefix("Test: divisible by ")
            current_monkey.test_modulus = int(line_stripped)
        else:
            raise ValueError(f"Invalid line: {line}")
    elif indent == 2:
        line_stripped = line.strip(" ")
        assert current_monkey is not None
        if line_stripped.startswith("If true:"):
            line_stripped = line_stripped.removeprefix("If true:")
            next_monkey = True
        elif line_stripped.startswith("If false:"):
            line_stripped = line_stripped.removeprefix("If false:")
            next_monkey = False
        else:
            raise ValueError(f"Invalid line: {line}")
        
        line_stripped = line_stripped.strip()
        
        if line_stripped.startswith("throw to monkey"):
            line_stripped = line_stripped.removeprefix("throw to monkey").strip()
            current_monkey.next_monkey[next_monkey] = int(line_stripped)
        else:
            raise ValueError(f"Invalid line: {line}")
    else:
        raise ValueError(f"Invalid line: {line}")
    
monkeys.append(current_monkey)

number_of_rounds = PART2
modify_worry_level_before_test = False


def get_maximum_item(my_monkeys):
    maximum = 1
    for m in monkeys:
        maximum = max([maximum] + m.items)
    
    product = 1
    for i in range(1, maximum + 1):
        product = product * i
        
    return product


def get_next_level_bounded_by_three(level):
    return level // 3


def get_next_level_if_dotted(level):
    return level % MAX_VALUE


MAX_VALUE = get_maximum_item(monkeys)


for monkey_round in range(number_of_rounds):
    # print(f"Round {monkey_round + 1}")
    for monkey in monkeys:
        monkey.inspections = monkey.inspections + len(monkey.items)
        for item in monkey.items[:]:
            worry_level = item
            worry_level = monkey.operation(worry_level)
            if modify_worry_level_before_test:
                worry_level = get_next_level_bounded_by_three(worry_level)
            result = (worry_level % monkey.test_modulus) == 0
            next_monkey_id = monkey.next_monkey[result]
            next_monkey = monkeys[next_monkey_id]
            next_worry_level = worry_level
            if not modify_worry_level_before_test:
                next_worry_level = get_next_level_if_dotted(next_worry_level)
            next_monkey.items.append(next_worry_level)
            monkey.items.remove(item)
            
monkey_inspections = [monkey.inspections for monkey in monkeys]

first = max(monkey_inspections)
monkey_inspections.remove(first)
second = max(monkey_inspections)

monkey_business = first * second

print([m.inspections for m in monkeys])
print(monkey_business)
