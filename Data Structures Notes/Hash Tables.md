# Chapter 1: Arrays and Strings

## Hash Tables

A hash table is a data structure that maps keys to values for highly efficient lookups.

You can think of this as building an indexed list from a list of values.

You can think of the common implementation of the hash map as just adding an index to data values, where the index has a limited range. This can result in creating an index that refers to a list of values if some of the values have the same index.

A hashmap or table is like a drawer that stores things in bins and labels the items by which drawer they belong in.

The difference between an array and a hashtable or map is that the key for looking up a value in an array is always a numeric index, whereas for a hashmap the lable could be a string, number, object, or anything else.

The hashmap uses the array, and maps the labels to array indexes with the hash function.

### Language Specific Implementations

- Python: the dictionary data type represents a hash table.
- Javascript: Javascript objects and the Map class are both implementations of hash maps in Javascript.
- Go: has a map implementation that is the same as a hash table.

### Implementation

1. first compute the key's hash code, which will usually be an integer (or Big Integer, or long). Note that two differetne keys could have the same hash code, as there may be an infinite number of keys and a finite number of ints.
2. Then map the hash code to an index in the array. This could be done with something like `hash(key) % array_length`. Two different hash codes could map to the same index though.
3. At this index, there is a linked list of keys and values. Store the key and value in this index. The linked list is required here because you may have two different keys with the same hash code, or two different hash codes that map to the same index.

Then to retrieve the value pair by its key, you repeat this process:

1. compute the hash code from the key,
2. then compute the index from the hash code.
3. Then search through the linked list for the value with this key.

### Run Time

If the number of collisions is very high, the worst case runtime is O(n), where N is the number of keys. A good implementation should keep the number of collisions as small as possible (collisions being the number of keys mapped to the same index), so the amortized lookup time is O(1).

### Hash Tables and Hash Maps in Python
