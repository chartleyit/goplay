# heap
data structure that can be expressed as a complete tree

# binary heap

# max heap
parent has a higher key then it's children

used when removing the largest key

note: min heap is the reverse parent has smaller than children

heaps can be expressed as a tree but heaps are expressed as an array that operates like a tree

children / parents based on parent / child index

39x2 + 1 = index of left
39X2 + 2 = index of right

keys added to bottom/right (last index)

parent key larger than children
- swap if new node is higher
- repeat until it reaches its proper place

- heapify
  - done when adding or removing an item from tree

- extract... last goes to root, then swap down