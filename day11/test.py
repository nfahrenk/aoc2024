import functools

@functools.cache
def helper(num, blinks):
    strNum = str(num)
    if blinks == 0:
        return 1
    elif num == 0:
        return helper(1, blinks-1)
    elif len(strNum) % 2 == 0:
        mid = len(strNum) // 2
        left = int(strNum[:mid])
        right = int(strNum[mid:])
        middle = helper(left, blinks-1) + helper(right, blinks-1)
        return middle
    else:
        return helper(num * 2024, blinks-1)

def part1(line, blinks):
    nums = [helper(int(elem), blinks) for elem in line.split(" ")]
    print(nums)
    return sum(nums)

if __name__ == "__main__":
    print("solution", part1("125 17", 75))
