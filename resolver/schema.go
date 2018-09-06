package resolver

// Schema for api
const Schema = `
schema {
    query: Query
    mutation: Mutation
}

type Query {
    # A user profile object.
    profile(): User!
    # A slice of thread.
    threadSlice(tags: [String!], query: SliceQuery!): ThreadSlice!
    # A thread object.
    thread(id: String!): Thread!
    # A post object.
    post(id: String!): Post!
    # Containing mainTags and tagTree.
    tags(): Tags!
}

type Mutation {
    # Register/Login via email address. An email containing login info will be sent to the provided email address.
    auth(email: String!): Boolean!
    # Set the Name of user.
    setName(name: String!): User!
    # Save/Add/Del tags subscribed by user.
    syncTags(tags: [String]!): User!
	addSubbedTags(tags: [String!]!): User!
	delSubbedTags(tags: [String!]!): User!
    # Publish a new thread.
    pubThread(thread: ThreadInput!): Thread!
    # Publish a new post.
    pubPost(post: PostInput!): Post!
}

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

scalar Time

// Data Type Defines

type User {
    email: String!
    # The Name of user. Required when not posting anonymously.
    name: String
    # Tags saved by user.
    tags: [String!]
}

# Construct a new thread.
input ThreadInput {
    # Toggle anonymousness. If true, a new ID will be generated in each thread.
    anonymous: Boolean!
    content: String!
    # Required. Only one mainTag is allowed.
    mainTag: String!
    # Optional, maximum of 4. 
    subTags: [String!]
    # Optional. If not set, the title will be '无题'.
    title: String
}

type Thread {
    # UUID with 8 chars in length, and will increase to 9 after 30 years.
    id: String!
    # Thread was published anonymously or not.
    anonymous: Boolean!
    # Same format as id if anonymous, name of User otherwise.
    author: String!
    content: String!
    createTime: Time!
    # Only one mainTag is allowed.
    mainTag: String!
    # Optional, maximum of 4.
    subTags: [String!]
    # Default to '无题'.
    title: String
    replies(query: SliceQuery!): PostSlice!
	countOfReplies: Int!
}

type ThreadSlice {
    threads: [Thread]!
    sliceInfo: SliceInfo!
}


input PostInput {
    threadID: String!
    anonymous: Boolean!
    content: String!
    # Set referring PostIDs.
    refers: [String!]
}

type Post {
    id: String!
    anonymous: Boolean!
    author: String!
    content: String!
    createTime: Time!
    refers: [Post!]
	countOfRefered: Int!
}

type PostSlice {
    posts: [Post]!
    sliceInfo: SliceInfo!
}

type Tags {
    # Main tags are predefined manually.
    mainTags: [String!]!
    # Recommended tags are picked manually.
    recommended: [String!]!
    tree(query: String): [TagTreeNode!]
}

type TagTreeNode {
    mainTag: String!
    subTags: [String!]
}
`
