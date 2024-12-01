from collections import defaultdict

def read_input(filename):
    list1 = []
    list2 = []
    with open(filename) as f:
        for line in f:
            line = line.rstrip()
            results = line.split(' ')
            results = [elem for elem in results if elem != '']
            list1.append(int(results[0]))
            list2.append(int(results[1]))
    return list1, list2

def find_map(filename):
    list1,list2 = read_input(filename)
    list1.sort()
    list2.sort()
    return sum([abs(list2[i] - list1[i]) for i in range(0, len(list1))])

def find_map2(filename):
    list1,list2 = read_input(filename)
    counter = defaultdict(int)
    output = 0
    for i in list2:
        counter[i] += 1
    for i in list1:
        output += i * counter[i]
    return output

if __name__ == "__main__":
    print(find_map2('aoc2024/day1/input2.txt'))
