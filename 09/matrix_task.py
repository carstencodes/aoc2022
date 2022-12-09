from enum import IntEnum


VERBOSE = False
DYNAMIC_BOARD = False


class State(IntEnum):
    Blank = 0
    Visited = 1
    Start = 5  # Visited & Tail
    Head = 2
    Tail = 4


def sgn(x):
    if x > 0:
        return 1
    elif x < 0:
        return -1
    else:
        return 0


def maybe_extend(current_move, position, projections: list[list[int]]):
    y0, x0 = position

    extend_at_end = sgn(current_move[0] > 0) or sgn(current_move[1] > 0)
    extend_row = current_move[0] != 0
    extend_col = current_move[1] != 0

    new_projections = projections

    max_row_length = len(projections)
    max_col_length = 0
    for row in projections:
        max_col_length = max(max_col_length, len(row))

    extended = False

    if extend_row and (y0 >= max_row_length or y0 < 0):  # new row
        new_projections = []
        for row in projections:
            new_projections.append(row)
        new_list = [0 for _ in range(max_col_length)]
        if extend_at_end:
            new_projections.append(new_list)
        else:
            new_projections.insert(0, new_list)
        if VERBOSE:
            print(f"New row: {projections} -> {new_projections} ({extend_at_end})")
        extended = not extend_at_end
    elif extend_col and (x0 >= max_col_length or x0 < 0):  # new column
        new_projections = []
        offset = int(not extend_at_end)
        for row in projections:
            next_row = [0 for _ in range(max_col_length + 1)]
            for i, item in enumerate(row):
                next_row[i + offset] = item
            new_projections.append(next_row)
        if VERBOSE:
            print(f"New column: {projections} -> {new_projections} ({extend_at_end})")
        extended = not extend_at_end

    return new_projections, extended


def format_board(projection: list[list[int]]):
    result = ""
    for row in projection:
        for cell in row:
            if cell in (0, 1):
                result = result + "."
            elif cell in (
                int(State.Head),
                int(State.Head) + int(State.Tail),
                int(State.Head) + int(State.Visited),
                int(State.Head) + int(State.Tail) + int(State.Visited),
            ):
                result = result + "H"
            elif cell in (int(State.Tail), int(State.Tail) + int(State.Visited)):
                result = result + "T"
            else:
                result = result + str(cell)
        result = result + "\n"

    return result


def format_visits(projection: list[list[int]]):
    result = ""
    for row in projection:
        for cell in row:
            if cell == 0:
                result = result + "."
            elif cell in (int(State.Head), int(State.Head) + int(State.Tail)):
                result = result + "."
            else:
                result = result + "#"
        result = result + "\n"

    return result


def is_adjacent(delta):
    set_of_direct_neighbors = (
        (0, 1),
        (1, 0),
        (-1, 0),
        (0, -1),
        (1, 1),
        (-1, -1),
        (-1, 1),
        (1, -1),
    )
    return delta in set_of_direct_neighbors


def is_hovering(delta):
    return delta == (0, 0)


def move_tail(head, tail, projections: list[list[int]]):
    if head == tail:
        return tail

    delta = ((head[0] - tail[0]), head[1] - tail[1])
    new_tail = tail

    if not (is_hovering(delta)) and not (is_adjacent(delta)):

        delta_row, delta_col = delta
        pos_dr = delta_row * sgn(delta_row)
        pos_dc = delta_col * sgn(delta_col)

        if pos_dc > 0 and pos_dr > 0:  # diagonal move
            if pos_dc > pos_dr:  # more horizontal than vertical
                new_tail = (head[0], tail[1] + sgn(delta_col))
            elif pos_dr > pos_dc:  # more vertical then horizontal
                new_tail = (tail[0] + sgn(delta_row), head[1])
            else:
                raise ValueError(f"Invalid position delta: ({pos_dr}, {pos_dc})")
        elif pos_dr > 0 or pos_dc > 0:
            # vertical move
            #               horizontal move
            new_tail = (tail[0] + sgn(delta_row), tail[1] + sgn(delta_col))
    else:
        if VERBOSE:
            print(f"T: T hovers H or is in direct neighborhood: {delta}")

    if VERBOSE:
        print(f"T: ({tail[0]}, {tail[1]}) -> ({new_tail[0]}, {new_tail[1]})")
    row, col = tail
    new_row, new_col = new_tail

    if head == new_tail:
        return new_tail

    projections[row][col] = int(State.Visited)
    projections[new_row][new_col] = (
        projections[new_row][new_col] + int(State.Tail) + int(State.Visited)
    )

    new_delta = ((head[0] - new_tail[0]), head[1] - new_tail[1])
    if not (is_adjacent(new_delta) or is_hovering(new_delta)):
        raise ValueError(f"Invalid tail move: {new_delta}")

    return new_tail


def move_head_tail(head, tail, current_move, projections: list[list[int]]):
    delta_row, delta_col = current_move
    assert sgn(delta_row) + sgn(delta_col) != 0

    new_head = head
    new_tail = tail

    if delta_row != 0:
        assert delta_col == 0
        while delta_row != 0:
            row, col = new_head
            move_row = sgn(delta_row)
            delta_row = delta_row - move_row

            old_head = new_head
            new_head = (row + move_row, col)

            if DYNAMIC_BOARD:
                projections, extended = maybe_extend(
                    (move_row, 0), new_head, projections
                )
                if extended:
                    old_head = (old_head[0] + 1, old_head[1])
                    new_head = (new_head[0] + 1, new_head[1])
                    new_tail = (new_tail[0] + 1, new_tail[1])
                row, col = old_head

            if VERBOSE:
                print(
                    f"H: ({old_head[0]}, {old_head[1]}) -> ({new_head[0]}, {new_head[1]})"
                )

            projections[row][col] = projections[row][col] - int(State.Head)
            projections[row + move_row][col] = projections[row + move_row][col] + int(
                State.Head
            )

            if VERBOSE:
                print(format_board(projections))
                print("~~~~~~~~~~~~~~~~~~~~~~~~")

            if old_head != new_tail:
                new_tail = move_tail(new_head, new_tail, projections)

            if VERBOSE:
                print(format_board(projections))
                print("~~~~~~~~~~~~~~~~~~~~~~~~")
    elif delta_col != 0:
        assert delta_row == 0
        while delta_col != 0:
            row, col = new_head
            move_col = sgn(delta_col)
            delta_col = delta_col - move_col

            old_head = new_head
            new_head = (row, col + move_col)

            if DYNAMIC_BOARD:
                projections, extended = maybe_extend(
                    (0, move_col), new_head, projections
                )
                if extended:
                    old_head = (old_head[0], old_head[1] + 1)
                    new_head = (new_head[0], new_head[1] + 1)
                    new_tail = (new_tail[0], new_tail[1] + 1)
                row, col = old_head

            if VERBOSE:
                print(
                    f"H: ({old_head[0]}, {old_head[1]}) -> ({new_head[0]}, {new_head[1]})"
                )

            projections[row][col] = projections[row][col] - int(State.Head)
            projections[row][col + move_col] = projections[row][col + move_col] + int(
                State.Head
            )

            if VERBOSE:
                print(format_board(projections))
                print("~~~~~~~~~~~~~~~~~~~~~~~~")

            if old_head != new_tail:
                new_tail = move_tail(new_head, new_tail, projections)

            if VERBOSE:
                print(format_board(projections))
                print("~~~~~~~~~~~~~~~~~~~~~~~~")
    else:
        raise ValueError(f"Invalid move: {move}")

    return (new_head, new_tail, projections)


SAMPLE = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2"""

PUZZLE = """U 1
L 1
D 2
U 2
R 2
D 1
L 1
D 2
R 2
D 2
U 1
L 1
D 2
U 1
D 2
L 2
R 1
U 1
L 1
R 1
U 1
R 2
L 1
D 2
U 1
R 1
L 2
R 2
L 2
R 2
L 2
U 1
L 2
U 1
D 2
L 2
R 2
D 2
L 2
U 2
L 1
U 1
D 2
L 1
D 1
R 1
U 1
L 2
D 2
U 2
L 1
D 1
R 2
D 1
R 2
D 1
R 2
U 2
R 2
L 2
R 1
D 1
U 1
R 2
D 1
U 1
R 2
L 2
U 2
R 2
U 1
D 1
R 2
L 1
U 1
L 2
D 1
R 1
U 2
L 1
D 2
L 1
R 2
D 2
L 2
R 2
U 2
R 2
D 2
L 1
R 2
U 1
R 2
U 1
L 1
R 1
L 2
U 2
L 1
D 2
L 2
U 2
D 2
R 2
D 2
R 1
U 2
L 1
U 1
L 2
R 1
D 1
R 1
L 1
D 2
L 2
R 2
D 3
R 2
D 2
L 3
R 1
L 2
U 3
D 1
U 1
L 2
D 2
L 2
R 2
L 1
U 2
D 2
U 2
D 1
L 1
D 1
R 2
D 3
U 3
D 3
L 2
R 2
U 1
L 3
R 2
U 2
D 2
L 3
D 2
R 3
D 3
U 2
R 2
D 2
U 1
L 3
R 3
U 3
L 3
D 3
L 3
D 2
L 1
R 2
L 1
R 2
D 1
L 3
R 2
D 3
L 1
R 1
L 2
U 1
L 2
R 2
U 2
L 1
D 1
R 3
L 1
R 2
U 2
D 1
R 2
U 1
R 3
U 1
D 3
R 3
L 1
D 1
U 1
D 3
R 2
U 1
R 1
U 1
L 3
D 2
U 3
R 2
D 1
R 1
U 2
L 1
U 1
R 3
U 2
D 1
R 2
U 3
D 2
U 1
R 2
L 2
U 3
L 1
R 1
L 2
D 2
U 1
D 2
R 3
D 3
U 4
R 3
D 2
U 4
L 4
U 3
L 1
U 2
L 2
R 3
U 4
D 1
L 1
D 2
L 2
R 1
L 2
U 1
L 3
D 1
L 2
D 3
R 3
D 4
L 2
U 4
L 3
R 4
D 4
R 2
D 2
L 2
R 2
U 4
L 3
D 1
R 1
L 2
D 1
U 1
L 2
U 1
L 1
U 2
D 4
L 4
U 2
R 2
L 2
R 3
U 2
D 3
U 4
D 3
R 2
L 4
U 2
R 2
U 3
D 4
R 4
L 2
R 4
D 4
U 3
L 3
D 4
L 4
R 2
L 1
R 3
D 2
U 4
R 3
U 1
R 2
U 1
D 1
R 1
L 4
U 4
R 1
D 4
L 3
U 1
D 2
R 2
L 2
D 2
U 4
R 3
L 4
D 1
R 1
U 3
R 4
L 2
R 2
U 3
L 4
R 1
L 3
D 3
R 2
D 2
L 2
D 2
U 4
D 5
L 4
U 2
D 3
L 1
U 4
R 5
D 5
U 4
R 5
L 2
D 4
R 2
L 2
R 1
U 1
D 5
U 1
D 2
R 5
L 4
D 4
R 2
L 3
R 3
L 1
U 4
R 2
L 4
U 5
R 5
D 3
U 4
L 2
U 5
R 3
U 5
L 4
R 2
U 2
R 3
L 5
D 2
R 5
L 2
R 5
D 4
U 4
L 5
U 1
D 4
U 5
L 5
D 2
L 2
D 4
L 2
D 2
R 4
U 4
D 2
L 1
D 3
U 3
R 2
D 2
L 5
D 3
R 2
L 2
U 3
D 5
R 1
L 4
D 1
L 5
R 3
L 3
D 3
L 2
U 5
L 5
R 4
D 5
U 2
R 2
L 3
U 1
D 1
U 2
R 3
U 2
L 4
U 4
L 3
U 3
D 3
R 2
L 5
R 5
U 2
D 2
U 4
D 3
U 1
L 1
D 4
R 3
L 1
D 5
R 3
U 1
L 6
U 5
D 4
R 4
L 5
U 2
D 6
L 4
R 2
U 4
L 2
R 1
U 6
D 6
R 1
L 2
U 4
L 2
D 6
L 6
R 3
L 3
D 5
U 6
D 5
U 5
L 6
R 6
U 3
R 6
U 2
D 1
U 3
R 6
D 5
L 3
R 1
D 6
R 3
D 2
U 2
D 3
L 6
U 3
L 3
U 1
R 6
U 2
D 3
R 1
U 3
R 6
L 2
R 4
L 1
R 5
L 1
R 4
U 6
D 3
U 2
L 4
R 1
L 2
R 4
U 5
R 1
L 2
U 1
R 2
U 6
R 1
D 2
L 1
U 2
D 2
L 1
U 1
R 4
D 5
L 4
D 6
L 2
U 3
D 3
L 1
R 3
D 1
L 6
D 1
U 6
L 1
D 6
U 3
R 5
L 2
R 4
D 5
L 4
D 6
U 1
D 6
L 6
U 3
D 3
U 2
L 1
R 5
U 5
L 6
U 5
R 5
D 4
U 7
R 7
L 2
D 3
R 3
L 1
D 6
R 3
L 2
U 2
D 2
U 2
L 7
R 4
U 3
L 3
R 2
L 5
U 3
R 6
U 5
L 7
U 7
L 4
U 5
L 7
D 1
L 1
D 2
U 3
R 5
D 1
R 6
L 3
U 5
D 6
L 4
D 3
U 3
R 3
U 1
R 2
L 7
D 7
R 4
L 6
D 7
L 7
D 2
U 1
R 2
U 3
L 2
D 2
L 5
D 4
R 5
D 7
U 5
D 1
L 4
U 2
R 5
D 6
L 4
R 7
U 6
D 2
L 4
D 4
L 1
R 4
L 1
R 5
U 5
R 5
L 1
D 2
R 1
U 5
L 4
D 2
U 3
L 2
R 5
U 4
L 1
R 5
L 1
D 5
R 7
L 5
U 5
L 7
U 3
D 6
U 2
L 2
R 5
U 7
D 1
R 2
U 4
R 7
U 6
L 3
U 5
D 3
U 5
D 6
R 2
U 8
R 7
L 5
U 1
D 3
L 3
U 6
L 6
U 5
D 7
U 3
L 6
U 1
D 3
R 5
L 2
R 1
D 6
R 2
L 1
R 5
D 3
L 1
R 3
L 6
D 5
U 5
L 5
D 3
L 4
R 7
D 6
L 8
R 6
U 8
D 6
R 3
U 4
R 1
U 6
R 3
L 6
D 2
U 5
R 2
D 5
U 7
D 2
U 2
R 3
L 3
U 4
R 3
U 7
L 2
D 3
R 2
D 4
L 4
D 8
R 2
U 2
L 7
R 8
U 8
L 4
U 3
R 3
U 8
D 2
U 5
D 8
R 8
L 8
R 4
D 8
L 7
U 6
L 4
D 6
U 2
L 4
U 8
R 8
U 6
D 1
R 5
U 5
L 8
U 2
R 3
U 7
L 5
U 5
L 7
D 8
R 7
U 4
D 7
U 4
L 3
R 3
D 5
U 1
L 6
D 3
U 5
D 3
R 1
U 2
L 3
R 8
U 1
D 6
R 6
U 9
D 4
U 4
D 2
L 1
R 2
L 6
U 6
L 1
R 1
U 9
D 1
R 4
L 3
D 8
U 7
D 1
R 7
L 6
U 3
R 2
L 2
D 6
U 5
D 6
L 6
U 8
L 5
R 4
U 6
L 1
U 8
L 4
R 1
D 7
R 9
U 2
D 3
R 6
L 8
U 6
R 4
D 1
U 7
R 1
D 9
L 8
D 9
R 1
U 6
R 7
D 2
U 7
R 4
D 9
R 7
D 2
U 1
R 3
D 4
L 5
U 7
R 7
D 2
U 8
R 8
L 4
U 6
L 3
D 8
L 8
D 9
R 7
L 5
U 5
R 4
U 9
R 3
U 9
R 1
U 2
R 9
U 8
L 8
U 7
R 1
D 9
U 4
L 7
R 3
U 5
L 6
D 1
L 4
D 4
U 6
R 7
L 7
U 3
L 8
D 8
L 7
U 3
D 4
R 8
D 1
R 8
U 10
L 1
U 2
D 4
U 3
R 8
U 3
D 5
R 8
D 10
R 8
D 4
R 9
D 1
L 3
U 7
R 7
L 1
D 8
R 6
D 4
U 5
L 2
U 3
R 7
U 10
L 1
U 7
R 1
U 7
D 2
L 10
U 2
L 4
D 6
R 10
D 4
U 7
L 6
R 4
U 6
D 1
U 6
D 5
U 6
L 4
D 2
L 2
U 7
R 2
U 9
L 9
D 2
U 9
R 9
U 9
D 3
R 5
D 3
L 10
U 3
D 10
L 10
D 1
U 2
D 3
R 8
L 1
U 6
L 2
R 9
U 3
D 5
L 8
U 4
R 8
D 5
L 6
U 7
D 8
R 9
U 9
R 8
D 4
R 5
U 6
L 8
R 2
D 3
L 8
D 6
R 7
L 10
U 9
R 6
L 7
U 8
R 5
L 1
D 1
R 6
D 1
R 9
L 1
U 5
R 3
D 7
U 8
D 6
L 6
U 11
D 8
L 6
D 6
L 7
R 10
U 5
D 7
R 5
D 3
L 4
R 3
D 2
L 6
U 6
D 3
L 4
D 9
L 6
U 9
L 3
U 10
R 5
U 8
R 2
D 3
L 6
R 2
U 6
R 8
D 2
U 2
R 5
U 9
L 5
D 8
R 2
U 1
R 1
D 6
U 4
R 7
L 7
D 7
U 11
R 6
U 11
D 11
R 6
L 2
D 10
R 6
D 5
R 7
D 11
R 2
U 10
L 6
D 4
R 1
L 3
D 9
U 8
L 6
U 1
L 7
D 10
U 1
L 1
D 2
U 3
L 4
R 6
U 11
R 6
U 3
L 8
D 3
R 7
L 3
D 4
L 5
U 8
L 11
R 1
U 1
L 9
D 5
U 3
R 1
L 6
U 1
L 3
D 1
L 8
U 1
D 9
R 7
U 11
R 10
U 8
R 5
U 3
L 7
R 3
L 4
U 6
D 3
U 3
R 3
D 1
L 7
U 5
D 1
U 6
L 8
D 3
R 6
U 10
R 6
D 1
U 7
R 5
L 6
R 5
D 3
R 9
U 11
L 8
R 9
U 5
R 5
L 1
U 7
L 11
U 5
R 3
U 7
R 9
L 11
D 11
L 3
U 11
L 2
U 3
R 5
D 1
L 10
R 8
L 11
U 1
D 11
L 4
R 8
L 9
D 6
U 10
L 1
R 12
U 2
D 8
L 8
U 1
L 12
D 9
R 3
U 3
D 10
U 11
D 9
U 8
L 7
D 10
L 3
U 9
D 5
R 12
D 3
R 9
L 8
R 12
U 11
D 5
R 10
D 8
L 12
U 7
L 1
R 7
D 1
R 12
L 3
R 7
L 1
R 2
D 4
U 3
L 8
U 2
R 8
U 5
L 5
R 7
L 7
R 7
U 8
R 3
D 4
L 7
U 5
L 5
U 1
D 12
R 11
L 5
D 10
R 8
U 12
L 11
U 2
R 2
D 1
U 4
D 1
U 1
R 11
D 1
L 8
U 5
L 11
D 11
R 12
L 6
R 8
L 13
U 7
R 7
D 1
R 13
D 6
L 5
R 10
U 4
R 11
D 6
U 1
D 6
R 1
U 1
L 12
U 8
R 4
U 9
R 12
D 1
R 8
D 2
U 5
L 7
U 12
L 5
D 10
R 8
D 6
R 13
D 7
U 7
L 9
D 12
U 5
L 7
D 4
U 2
R 13
U 11
L 3
R 6
D 5
R 6
L 11
U 6
D 5
U 6
L 2
D 8
R 4
U 4
L 12
D 6
U 8
D 8
L 5
D 12
R 8
U 11
D 8
L 7
U 13
L 11
R 6
U 5
D 9
L 11
U 10
D 13
R 5
D 8
R 7
D 6
R 7
D 2
L 13
R 11
U 1
D 7
U 13
R 9
U 12
R 13
L 2
U 13
D 2
U 6
L 6
D 5
R 5
L 2
U 1
R 4
U 6
L 10
R 12
L 3
D 3
L 8
D 10
U 2
L 11
U 2
L 11
R 8
U 12
R 3
U 5
L 7
D 6
U 6
L 10
R 4
L 1
U 8
D 4
R 13
L 7
U 6
L 9
U 13
D 13
L 8
U 3
L 13
D 14
L 1
U 9
D 1
U 9
R 12
D 7
L 9
R 13
D 12
U 10
D 4
L 13
R 11
L 2
D 9
R 4
U 10
L 2
D 8
U 11
L 14
R 1
U 2
D 9
U 8
L 4
D 14
U 10
R 3
D 12
U 12
L 1
U 1
D 7
U 8
R 9
D 7
R 11
D 11
R 13
D 14
L 5
D 1
U 12
D 7
U 13
D 9
U 6
R 5
U 13
L 4
D 8
L 3
U 8
D 4
R 4
D 2
L 1
D 12
R 11
U 11
D 6
R 12
U 3
D 2
R 6
D 1
U 11
L 7
D 9
L 8
R 1
U 13
D 7
L 9
U 4
D 8
L 12
D 6
R 4
D 5
U 13
L 2
D 11
R 4
D 5
L 4
U 9
L 12
U 10
D 3
R 15
L 3
D 10
L 4
R 2
D 10
R 1
U 10
L 9
D 7
L 5
R 8
U 3
D 9
R 5
D 8
U 5
D 5
L 8
U 2
R 5
U 4
D 6
R 3
L 1
U 15
R 12
D 14
R 12
U 14
D 8
U 10
R 9
D 8
L 10
R 4
U 5
D 15
U 6
D 4
R 7
D 1
R 15
U 15
D 4
R 6
L 1
R 7
L 12
U 3
R 1
L 6
D 6
R 5
D 15
R 1
U 10
R 4
D 13
L 5
D 10
L 7
R 5
U 15
L 11
U 4
D 9
U 12
R 12
L 9
D 10
U 7
R 3
L 10
R 12
D 2
U 8
L 7
R 3
U 1
D 10
L 11
R 12
L 8
U 1
D 12
L 4
D 4
L 6
D 7
R 3
L 7
U 1
R 4
L 1
U 2
L 6
D 4
L 7
R 7
L 15
U 4
R 7
D 8
U 9
D 3
L 5
U 6
R 2
D 10
L 9
U 16
L 15
U 1
R 15
U 13
L 12
D 10
R 1
D 6
R 3
D 1
L 14
R 10
D 2
R 10
U 12
R 10
L 13
U 14
D 11
U 16
R 6
U 4
R 6
U 13
R 5
D 1
L 12
U 14
D 11
R 11
D 7
L 3
R 7
L 9
U 11
D 6
U 14
D 2
L 1
R 16
L 14
D 10
U 4
R 15
U 11
R 10
D 5
R 8
U 13
D 11
U 9
R 8
D 4
L 11
U 8
R 7
D 14
U 8
D 6
L 13
R 2
L 16
R 10
U 7
L 10
U 2
L 13
U 8
D 5
U 3
R 14
D 15
R 10
D 4
U 10
D 15
R 6
U 13
R 1
U 16
L 2
U 6
L 8
D 2
L 11
U 15
D 5
U 1
R 3
D 16
L 3
R 4
D 8
R 10
L 8
D 1
U 2
R 4
U 3
D 16
U 10
R 11
L 11
R 7
L 3
U 17
D 1
L 10
D 6
L 6
D 6
R 8
U 16
D 15
U 12
L 12
D 17
U 9
R 7
U 12
L 6
D 5
R 2
L 17
U 6
R 1
D 13
R 11
D 17
R 5
L 3
D 3
U 4
D 4
U 4
L 11
R 4
L 15
D 7
U 13
D 12
U 1
L 16
D 2
L 12
R 15
U 11
R 15
D 12
R 4
L 3
R 3
U 3
L 13
U 4
L 16
U 7
D 17
R 14
U 11
D 4
L 11
U 9
D 8
U 3
R 10
U 16
R 5
L 10
D 3
L 7
U 9
R 2
U 14
L 8
R 9
L 6
D 14
L 14
U 3
D 9
L 8
D 14
U 6
L 4
R 7
D 4
L 6
U 11
R 8
D 8
R 5
U 17
R 15
U 1
L 17
D 14
L 2
D 3
L 16
D 7
L 16
R 10
D 6
R 11
U 10
L 14
R 2
U 3
D 1
R 16
L 5
R 2
D 15
U 6
L 3
D 9
R 12
D 13
U 17
R 16
D 15
U 3
D 13
L 7
R 12
U 18
L 12
U 7
L 18
D 12
L 17
D 6
L 8
D 16
L 9
R 8
U 9
D 9
U 16
L 10
U 12
L 9
U 10
R 13
L 7
U 17
L 13
D 18
L 6
R 14
L 9
U 1
R 8
U 3
R 13
D 3
L 14
D 8
L 12
R 3
L 11
R 2
D 7
R 14
L 9
U 1
D 5
U 10
R 5
D 8
U 2
D 17
L 16
D 3
U 15
R 17
D 10
R 16
L 3
U 2
R 5
L 7
U 7
D 12
L 8
R 4
L 18
R 13
L 18
R 11
U 1
D 1
L 12
U 15
L 4
D 12
U 13
L 14
D 14
L 12
D 14
U 14
L 6
R 5
D 6
L 13
U 16
L 11
U 7
R 10
D 6
U 17
L 12
R 11
D 11
U 7
R 2
U 9
R 16
L 5
D 16
R 7
U 4
D 19
L 2
D 18
R 1
U 15
R 13
D 5
R 2
U 2
D 12
L 17
D 11
U 8
R 13
U 11
R 2
U 13
R 11
U 19
L 17
R 17
D 16
R 18
U 9
L 13
D 1
R 16
D 13
R 6
D 9
R 7
U 10
D 10
U 8
R 8
L 2
R 17
D 13
L 9
U 19
D 6
L 14
R 9
L 19
D 17
R 17
D 10
L 1
U 15
D 17
L 12
R 10
L 14
U 16
L 17
U 8
R 8
U 4
R 11
U 14
R 2
L 9
U 11
D 14
U 14
D 9
L 10
U 19
D 10
R 5
D 1
L 3
U 14
R 11
L 17
R 7
D 8
R 1
U 2
R 4
L 2
R 13
U 16
D 2
L 2
R 2
U 16
L 3
D 7
R 7
U 2
L 5
U 7
R 17
U 6
D 12
U 17
D 13
L 19
R 19
L 9
D 14
R 15
L 7
U 11
R 2
D 16
U 6
L 12"""

data = PUZZLE

PUZZLE_DIM = 640
SAMPLE_DIM = 10

HEIGHT: int = PUZZLE_DIM
WIDTH: int = PUZZLE_DIM

Y0: int = (HEIGHT // 2) - 1
X0: int = (WIDTH // 2) - 1

two_dim_projection: list[list[int]] = []
for _ in range(HEIGHT):
    two_dim_projection.append([0 for _ in range(WIDTH)])


two_dim_projection[X0][Y0] = int(State.Start) + int(State.Head)

head_ptr = (X0, Y0)
tail_ptr = (X0, Y0)

# print(two_dim_projection)

movement = data.splitlines()

moves = []
for move in movement:
    dir_and_count = move.split(" ")
    assert len(dir_and_count) == 2
    direction = dir_and_count[0]
    count = int(dir_and_count[1])

    if direction == "R":
        moves.append((0, count))
    elif direction == "L":
        moves.append((0, 0 - count))
    elif direction == "U":
        moves.append((0 - count, 0))
    elif direction == "D":
        moves.append((count, 0))
    else:
        raise ValueError(f"Invalid direction: {direction}")


# print(moves)

if VERBOSE:
    print(format_board(two_dim_projection))

for next_move in moves:
    head_ptr, tail_ptr, two_dim_projection = move_head_tail(
        head_ptr, tail_ptr, next_move, two_dim_projection
    )

# print(two_dim_projection)

one_dim_projection = []
for dim in two_dim_projection:
    one_dim_projection.extend(dim)

visits = sum(one_dim_projection) - int(State.Head) - int(State.Tail)

if VERBOSE:
    print(format_board(two_dim_projection))
    print(format_visits(two_dim_projection))

print(visits)
