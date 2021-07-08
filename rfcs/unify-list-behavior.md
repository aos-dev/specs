- Author: JinnyYi <github.com/JinnyYi>
- Start Date: 2021-07-07
- RFC PR: [beyondstorage/specs#0](https://github.com/beyondstorage/specs/pull/0)
- Tracking Issue: [beyondstorage/go-storage#0](https://github.com/beyondstorage/go-storage/issues/0)

# GSP-0: Unify List Behavior

Previous Discussion:

- [Specify the behavior of List](https://github.com/beyondstorage/specs/issues/135)

## Background

[go-storage](https://github.com/beyondstorage/go-storage) exposes a list operation that lets you enumerate the keys contained in a bucket, or the files contained in a directory.

The list operation has been changed many times and tended to be stable since [Proposal: Unify List Operation](./23-unify-list-operation.md)

```go
func List(path string, pairs ...Pair) (oi *ObjectIterator, err error)
```

**Parameters**

path - The directory path for file system, or a file hosting service like dropbox. Also, it could be a prefix filter for object storage.

pairs - It contains all the filters for `List`. Usually we can specify `ListMode` for list operation.

`ListMode` is the type for `List`.

- `ListModePrefix` means this list will use prefix type. The returned file or object names must contain the prefix.
- `ListModeDir` means this list will use dir type. That means list files or objects hierarchically.
- `ListModePart` means this list will use part type. Generally, it's used to retrieve a list of in-progress multipart uploads. Only services that support multipart upload could list keys with `ListModePart`.
- `ListModeBlock` means this list will use block type. Generally, it's used to retrieve the list of blocks. It's used for block related operations.

**Returns**

oi - An object iterator. You can retrieve all the object by `Next` and `IterateDone` will be returned while there are no items anymore.

err - It's nil if no error.


For the current implementation:

- Services have no default value for `ListMode` and will get `ListModeInvalidError` when no `ListMode` passed in which is not user-friendly.
- Whether we need to check `VirtualDir` while `ListMode` is `ListModeDir`. In [GSP-109: Redesign Features](./109-redesign-features.md), we introduced a new feature called `VirtualDir` to support simulated dir behavior and any operation related to dir should check `VirtualDir` in services that don't have native support for dir.

## Proposal

So I propose to specify the behavior of `List`.

### Support default `ListMode` for services

Service SHOULD support default `ListMode` if there's no `ListMode` passed in.

- File hosting services could be `ListModeDir`.
- Object storage services could be `ListModePrefix`.

### Implement `ListModeDir` without checking `VirtualDir`

Service SHOULD support `ListModeDir` for `List` and implement it without the check for `VirtualDir`.

Services like s3, oss or azblob, have a flat structure instead of a hierarchy like file system. However, for the sake of organizational simplicity, they support the `folder` concept as a means of grouping objects. The purpose of the prefix and delimiter is to help you organize and then browse the keys hierarchically. If prefix is specified and delimiter is set to a forward slash (/), only the objects in the "folder" are listed.

## Rationale

N/A

## Compatibility

Service should `List` with default `ListMode` instead of returning `ListModeInvalidError` if no `ListMode` passed in.

## Implementation

- All services should support default `ListMode`.
- All services should implement `ListModeDir`.
