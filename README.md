# Linked List Practice with Golang

This is a practice exercise in implementing a simple Linked List using Golang. The methods implemented for this version of the Linked List are as follows:

- `Find`: Traverses the list and searches for the given value. If the value is found, the method returns `true` and `false` otherwise.
- `Append`: Adds the given node to the end of the list
- `InsertLeft`: Inserts the given node to the left of the first occurrence of the given value, if it exists in the list. It does nothing otherwise.
- `InsertRight`: Inserts the given node to the right of the first occurrence of the given value, if it exists in the list. Otherwise it does nothing.
- `Delete`: Removes the first occurrence of the node with that has the given value, if it exists. Otherwise it does nothing.
- `Reverse`: Reverses the list in-place
- `Print`: A convenience function that prints the values of each node to show the list (e.g. `1 -> 2 -> 3 -> 4`)