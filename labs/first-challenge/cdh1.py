

def getTheLongestSubStringWithOutRepetingChar(word):
    letters = {}
    currentLengthSubString = 0
    maxLengthSubString = 0
    for currentLetter in word:
        if currentLetter in letters:
            maxLengthSubString = currentLengthSubString if currentLengthSubString >= maxLengthSubString else maxLengthSubString
            currentLengthSubString = 0
            letters = {}
        currentLengthSubString += 1
        letters[currentLetter] = None

    return maxLengthSubString


word = str(input())
print(getTheLongestSubStringWithOutRepetingChar(word))
