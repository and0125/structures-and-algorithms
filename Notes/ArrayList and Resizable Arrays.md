# ArrayLists and Resizable Arrays

Arrays are sometimes called lists, and can be automatically resized in some languages:

- python: python's lists only resize when necessary, and over-allocate space as the size increases. Also has slicing for making different size copies of the lists.
- javascript: has a resizable array structure, but, if a list is provided with holes in the indices, javascript will re-render the object as a hash table (or a dictionary), but this will be slower to access because it has to generate a hash function that accounts for the gap. javascript also supports slicing.
  - example: if you define list elements 1 through 5, then jump to define the element with index 400, the structure will switch to a hash map, but also has to develop a hash function that can handle those 400 potential indices.
- go: arrays can have a defined size or the compiler will find the size for you at run time and will resize as needed. Arrays cannot be re-sized in go, so slices are the best way to approach making smaller copies of a list.
