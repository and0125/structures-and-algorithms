# Binary Search

## Visualization

Think about finding a name in the phone book when you know a person's name starts with K. It makes more sense to start in the middle rather than go beginning to end, because you know that K is near the middle.

Searching from start to finish, beginning to end, is called simple search.

## I/O

The input to a binary search algorithm is a sorted list.

The output is either the position of the element in the list if the element is in the list, or null if the element is not in the list.

## Performance

Binary search cuts half of the number of inputs in half at every step or iteration of the search.

If we look at the worst case scenario, binary search takes log_2(n) steps to run, whereas simple search has a worse case of n steps.

Note, for this book, big O notation always represents Log base 2 logarithms, and not base 10. But the 2 is dropped because changing from base 2 to base 10 is just multiplying by a constant, and we ignore constants in big O notation.

