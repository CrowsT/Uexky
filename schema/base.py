types = '''
scalar Time

# SliceInfo objects are generated by the server. Can be used in consecutive queries.
type SliceInfo {
    firstCursor: String!
    lastCursor: String!
}

# SliceQuery object is for selecting specific 'slice' of an object to return. Affects returned SliceInfo.
input SliceQuery {
    # Either this field or 'after' is required.
    # An empty string means slice from the beginning.
    before: String
    # Either this field or 'before' is required.
    # An empty string means slice to the end.
    after: String
    # Set the amount of returned items.
    limit: Int!
}
'''
