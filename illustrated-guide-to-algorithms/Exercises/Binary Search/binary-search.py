"""
_Binary Search_

This python module implements a binary search for two lists.

input: sorted list 

output: 
    - index of value if in list and the number of steps
    - -1 if the value is not in the list (this return value could vary depending on requirements) and the number of steps
"""

# create lists, these are unorderd
mylist_integers = [3,4,510, 45,60, 23, 303]
mylist_names = ["Andrew", "Jake", "Russell", "Aaron", "Darius"]

# order the lists with the sort method
mylist_integers.sort()
mylist_names.sort()

def binary_search(inputList, searchItem):
    # starting with the first and last index
    low = 0
    high = len(inputList) - 1
    steps = 0
    # while the two indices are less than or equal
    while low <= high:
        steps += 1
        # find the midpoint
        mid = (low + high) // 2
        
        # store the value of the element of the midpoint
        guess = inputList[mid]

        # if the searchItem equals the guess, we are done
        if guess == searchItem:
            return mid, steps

        # because they are sorted, if the guess is great than the searchItem
        # we reduce the high value to search lower
        if guess > searchItem:
            high = mid - 1 
        
        # if the guess is lower than the searchItem
        # we increase the low value to search higher
        else:   
            low = mid + 1
    # if the loop breaks without finding the element, return -1
    return -1, steps

print(binary_search(mylist_integers, 5)) 
# returns (-1, 3) so it took 3 steps to know the element wasn't present instead of 7 with simple search
print(binary_search(mylist_integers,303), mylist_integers[5])
# returns (5, 2) so it took 2 steps to know the element was present instead of 7 with simple search
print(binary_search(mylist_names, "Terry"))
# returns (-1, 3) so it took 3 steps to know the element wasn't present instead of 5 with simple search
print(binary_search(mylist_names, "Aaron"),  mylist_names[0])
# returns (0, 2) so it took 2 steps to know the element was present instead of 7 with simple search