# Parallel workpool

A toy project to run functions in parallel using the concept of a workpool.

It uses generics and converts a list into a channel, applies a function to the elements in parallel and then collects the result back into a list.

If you want to see the progress of the repo, check the tags.
- v1 builds the mechanisms, but only works with strings
- v2 replaces strings with empty interfaces (interface{})
- v3 uses go1.18 generics
- v4 and v5 are minor improvements, making the go compiler to infer the types