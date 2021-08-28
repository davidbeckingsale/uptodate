# Up to Date

This is a library and GitHub action to automatically update different assets in your
repository. To start, we will be updating the `FROM` images of Dockerfiles.

**under development! not ready for use!**

## Usage

### Install

To build the library:

```bash
$ make
```

This will create a binary executable, `uptodate` that you can use directly or
copy into a directory on your path.

### dockerfile

The updtodate dockerfile command will update a single Dockerfile:

```bash
$ ./uptodate dockerfile /path/to/Dockerfile
```

or an entire directory of Dockerfile:

```bash
$ ./uptodate dockerfile /path/to/directory
```

To update your `Dockerfile`s we will use:

 - [containerspec](https://github.com/vsoch/containerspec): for LABELS
 - [lookout](https://github.com/alecbcs/lookout) for updated versions
 

Specifically, I'd like to have commands that can read a Dockerfile, or
a directory / repository of `Dockerfile`s, and be able to tell us:

1. Is the digest up to date?
2. Are there new tags we might want to build?

For base images and updating them this means:

1. User can target a Dockerfile directly for one off update, or a folder with tags for scaled
2. Read in Dockerfile, keep track of labels and FROMS (add dockerfile parser)
3. Look at labels to see if a tag is there for the hash
4. For each FROM, look up list of tags, update hash (use lookout)
5. For each FROM, if a label exists after it for opencontainers, delete it
6. Update label to use new tag


## Previous Art

 - [binoc](https://github.com/autumus/binoc): is another updating tool that uses lookout, and the main difference will be in the design.
