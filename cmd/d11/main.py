from functools import cache

rocks = [int(x) for x in input().split()]


@cache
def count(rock: int, steps: int):
    if steps == 0:
        return 1
    if rock == 0:
        return count(1, steps - 1)
    str_rock = str(rock)
    length = len(str_rock)
    if length % 2 == 0:
        return count(int(str_rock[: length // 2]), steps - 1) + count(
            int(str_rock[length // 2 :]), steps - 1
        )
    return count(rock * 2024, steps - 1)


print(sum(count(rock, 75) for rock in rocks))
