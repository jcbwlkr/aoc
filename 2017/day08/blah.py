def isAnagram(a, b):
    tester = list(a)
    tested = list(b)
    print("Tester: ", tester)
    print("Tested: ", tested)
    print("testing...")
    i = 0
    while i < len(tester):
        j = 0
        while j < len(tested):
            print(i,j)
            if tester[i] == tested[j]:
                if len(tester) == 1:
                    return 1
                tester.pop(i)
                tested.pop(j)
                if isAnagram(tester, tested) == 1:
                    print("It's an anagram!")
                    return 1
            else:
                j += 1
        i += 1
    return 0


print(isAnagram("abc", "cba"))
