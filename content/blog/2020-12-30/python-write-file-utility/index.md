---
title: "Python write_file utility"
slug: python-write_file-utility
author: Christopher Loverich
date: "2020-12-30T00:00:00.000Z"
tags: devlog
description: "When simple functions take a bit of elbow grease"
---

Recently I wrote a python script for syncing data to a 3rd party service. I needed to output various text files during its run, both for debugging and audit history. Ideally, I wanted a simple function which would allow the following:

```python
# Set some base directory via config
working_dir = 'data'

# Pass filenames and contents, with any directories
# automatically created as necessary:

# should output to: data/job_2/stats.csv
write_file('/job_2/stats.csv.', raw_policy_data)

# should output to: data/job_2/applications/results.csv
write_file('/job_2/applications/results.csv', app_results)

# should raise -- we don't want to silently overwrite a file
write_file('/job_2/applications/results.csv', app_results)

```

I ran into a few snags implementing this simple behavior; this post documents what I learned along the way.

## Using Pathlib

My starting point used python's [pathlib](https://docs.python.org/3/library/pathlib.html) library. Its been a while since I've used python, and this was not part of the standard library previously. Initially, I thought it might give me what I wanted out of the box:

```python
from pathlib import Path

path = Path(working_dir).joinpath('job_2/results.csv')
path.write_file(app_resuls)
```

But this throws the following error:

```python
Traceback (most recent call last):
...
FileNotFoundError: [Errno 2] No such file or directory: 'data/job_2/results.csv'
```

Rats. We'll need to explicitly make some directories first, which justifies writing a little utility function to encapsulate it.

## Creating directories

For me, the expected behavior of directory creation is always that of the `mkdir -p`: Create all directories at this path if they do not exist. Wrapping this behavior and file writing into a utility, with `Path` as its base, we arrive at:

```python
working_dir = 'data'

def write_file(filepath: str, contents: str) -> None:
  path = Path(working_dir).joinpath(filepath)
  path.parent.mkdir()
  path.write_text(contents)

```

Depending on the existing directories, this may work at first. However, you'll encounter errors in two specific scenarios that make this untenable as written:

- If the file path has nested directories, this won't create them (and, confusingly, you'll get `FileNotFoundError: [Errno 2] No such file or directory: ...`)
- If the directories already exists (`FileExistsError: [Errno 17] File exists: ...`

While it is easy enough to reason about why `mkdir` raises these exceptions, it is not the behavior we want. Reading through the pathlib documentation, we find these scenarios can be handled with a few additional parameters:

> If parents is true, any missing parents of this path are created as needed; they are created with the default permissions without taking mode into account (mimicking the POSIX mkdir -p command).

> If parents is false (the default), a missing parent raises FileNotFoundError.

> If exist_ok is false (the default), FileExistsError is raised if the target directory already exists.

> If exist_ok is true, FileExistsError exceptions will be ignored (same behavior as the POSIX mkdir -p command), but only if the last path component is not an existing non-directory file.

The elusive `mkdir -p` behavior is actually called out specifically, which I found amusing. Ultimately I imagine it is the behavior _most_ users would expect. Perhaps there are complexities I am glossing over (and notably on Python versions < 3.5, there is [quite a bit of baggage](https://stackoverflow.com/questions/273192/how-can-i-safely-create-a-nested-directory) and the solution is not as simple). At any rate, updating the utility function looks like this:

```python
def write_file_demo(filepath: str, contents: str) -> None:
  path = Path(working_dir).joinpath(filepath)
  path.parent.mkdir(parents=True, exist_ok=True)
  path.write_text(contents)
```

## Permissive input and path segments

There lurks a sneaky issue with the path joining function as written.

```python

base_dir = 'data'

Path(base_dir).joinpath('/job_1/results/output.csv')

# prints -- where did the 'data' directory go?
PosixPath('/job_1/results/output.csv')
```

A leading slash at any point in the arguments _silently erases_ prior path segments. I... am not sure under what scenario this would be desire-able. I do not see this documented in pathlib, but it is called out in the docs for [`os.path.join`](https://docs.python.org/3/library/os.path.html#os.path.join)

> Join one or more path components intelligently... If a component is an absolute path, **all previous components are thrown away** and joining continues from the absolute path component.

Most languages I have worked with include a library for joining path segments, usually like `join(path_one, ...other_paths)`and giving you a "clean" result. I don't remember other languages behaving like Python — though it strikes me as sufficiently weird that I second guess my intuition. I quickly try in a few other languages — here's a test in Go:

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Join("a/b", "/c"))
}

// prints 'a/b/c'
```

No issue there. What about Node?

```jsx
const path = require("path")
path.join("a/b", "/c")

// prints 'a/b/c'
```

Nope. Python is weird.

At any rate, if we do not change this behavior, calling code may unintentionally write to root directories. This would likely be confusing and result in either misplaced files or seemingly random permissions errors. As your program grows, and path parts begin being passed around or through helper functions, it is natural that some leading slashes may begin slipping through. While it is tempting to raise an error, I prefer to simply fix it since we have to _detect_ it either way. In this case the solution is straight forward:

```python
if filepath.startswith('/'):
	filepath = filepath.strip('/')
```

## File Overwriting

Lastly, I want to prevent file overwriting. Given the way directory creation works, I found it surprising that by default, Python will (silently) overwrite files when you provide a path that already exists. From reading the docs and some light testing, I cannot figure out a way to get `write_text` to throw on its own. Instead, we'll need to add another check:

```python
if path.is_file():
      raise FileExistsError(f'{path} already exists')
```

One callout here is that `is_file` (and `is_dir`) look for an actual, existing file (or directory) rather than the structure of the path you've provided. I think a name like `file_exists` could clarify this distinction but it doesn't change the implementation.

## Recap

Putting it all together, we get the behavior we want:

```python
def write_file(filepath: str, contents: str) -> None:
    """ Write a file as text, creating directories along the way """
    if filepath.startswith('/'):
        filepath = filepath.strip('/')

    path = Path(data_dir).joinpath(filepath)
    path.parent.mkdir(parents=True, exist_ok=True)
		if path.is_file():
      raise FileExistsError(f'{path} already exists')

		path.write_text(contents)
```

I hadn't thought writing a text file would warrant a write up, but there are enough surprising behaviors in Python that snags are inevitable — indeed this stackoverflow question on [creating nested directories](https://stackoverflow.com/questions/273192/how-can-i-safely-create-a-nested-directory) alone has a whopping 5500+ votes. Reflecting on Python's standard library design choices, I imagine there is a difficult balance when writing standard library code. The implementation must support a wide variety of use cases and its impossible to provide defaults that satisfy all users. Still, I would be interested in the motivations for these particular default behaviors, especially whether they are rooted in common use cases or merely concerns over backwards compatibility.

See also:

- [Why you should be using pathlib](https://treyhunner.com/2018/12/why-you-should-be-using-pathlib/)
- [Go by Example: Writing Files](https://gobyexample.com/writing-files)
- [Read and Write files](https://deno.land/manual/examples/read_write_files) (Deno)
