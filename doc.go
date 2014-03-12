package gocct

/* Centralised Composite Tree

A Composite Tree that store the tree relationships separate to the nodes.

NOTE: Since interface{} allows any type, this code checks if the parent and children nodes
provided are pointers. The implementation allows this to be turned off, but the client of
the collection needs to ensure that the key of the objects is not changes.
Also structure may not have a key, eg in the case of one that contains a slice
Structure without a key can't be used as Nodes
*/
