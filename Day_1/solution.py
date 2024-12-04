from math import fabs

def firstPart(left, right):
    result = 0
    for f, b in zip(left, right):
        result += fabs(f - b)

    print(int(result))

def secondPart(left, right):
    result = 0
    for x in left:
        result += x * right.count(x)
    print(result)

def main():
    with open("input.txt", "r+") as file:
        input = file.read()            
        input = [x.split(" ")  for x in input.split("\n")]
        left, right = [], []
        for line in input:
            left.append(int(line[0]))
            right.append(int(line[-1]))
        left.sort()
        right.sort()
    
    firstPart(left, right)
    secondPart(left, right)

if __name__ == "__main__":
    main()