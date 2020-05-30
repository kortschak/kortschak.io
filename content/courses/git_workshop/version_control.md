---
# Course title, summary, and position.
linktitle: Version Control
summary: Learn how to use Git for collaboration.
weight: 1

# Page metadata.
title: Version Control
date: "2018-01-09T00:00:00Z"
draft: false  # Is this a draft? true/false
toc: true  # Show table of contents? true/false
type: docs  # Do not modify.

# Add menu entry to sidebar.
# - name: Declare this menu item as a parent with ID `name`.
# - weight: Position of link in menu.
menu:
  git_workshop:
    name: Version Control
    parent: Git Workshop
    weight: 2

# Prev/next pager order (if `docs_section_pager` enabled in `params.toml`)
weight: 2
---

Version control is an important part of text development, whether that text is a scientific paper,
instructional document, or the source code of a piece of software.
Version control at its best provides a history of how a document got to be how it is now and the rationale for that path.

Historically, and still commonly in much of the world, version control is as limited as maintaining a set of distinctly named files or archives with an _ad hoc_ revision annotation.
These systems may, but often don't, include the reasons for changes and which version was the immediate predecessor.
They never scale well over time or for number of people involved.
They cause inordinate levels of pain for those that use them.

To address these kinds of problems, dedicated software was developed to help maintain version histories.
A brief history of open source version control systems will help illustrate the kinds of approaches that have been taken.

The version control system was developed in 1972 at Bell Labs, Source Code Control System (SCCS),
and later the Revision Control System (RCS) was written, adding the ability to operate across platforms.
These tools constitute the first generation of version control tools.
They were designed for single file, single user use.

The second generation of version control systems start in 1986 with the development of the Concurrent Versions System (CVS).
This tool, still in use in some places, allows more than one person to edit files and for multiple files to be tracked.
CVS was followed by Subversion (SVN) in 2000 which added features and improved on CVS.
Both CVS and SVN host the version-controlled text on a centralised server, allowing users to checkout files,
make changes and check them back in again.

The current, third, generation of version control system are distributed rather than centralised like CVS and SVN.
What this means is that there is no intrinsic central store required for the version control system,
although the project may specify one in practice.
Although distributed version control systems have existed since the 1990s,
they came to prominance when the Linux kernel's development version control system was moved to git,
a distributed version control system originally written by Linus Torvald, in 2005.
Other commonly used open source distributed version control systems are mercurial (hg) and bazaar (bzr).

For an extensive, though not up-to-date, explanation and description of version control systems, see [this article](http://www.catb.org/~esr/writings/version-control/version-control.html) by Eric S. Raymond,
or the [Getting Started chapter of the git book](https://www.git-scm.com/book/en/v2/Getting-Started-About-Version-Control).

## Git

For a variety of reasons, the [most commonly used](https://rhodecode.com/insights/version-control-systems-2016) version control system is git.

As said above, git is a distributed version control system.
Git retains versions as snapshots of files that have been committed to a repository
(an alternative approach used by some systems retains differences, rather than snapshots).

![Git stores data as snapshots of the project over time.](https://www.git-scm.com/figures/18333fig0105-tn.png)

These snapshots are retained locally on the computer where the work is being done.
This means that the entire history of all the file changes are available whether the computer can access a network or not.
However, it is possible to communicate the state of a repository to other computers, allowing publication or sharing of work.

In order to be able to communicate state in a safe and secure way, version are given unique identifiers based on the
repository state, author, authoring data, commiter, the commit message and date, and the hash of the previous version(s).
This is done by producing a cryptographic hash of these values; this is a one way function that condenses a long input into a
single, large, number with a low probability of two given inputs giving the same result.
The cryptographic hash function used by git is [SHA-1](https://en.wikipedia.org/wiki/SHA-1), which produces a 160 bit number from an input stream.
The SHA-1 hash is represented in hexadecimal, like so, `f164acde21f01cf3b2ba4f7dd650d9bf2a699b96`.

#### Questions

1. It is quite clear what the function of including the repository state, commit message and the author and commiter and time information.
Why is the SHA-1 of the previous version(s) also included in the hashed stream? Why is this important? (*Hint: Merkle tree*)
{{< spoiler text="Click to see the answer" >}}
Including the SHA-1 of the parents in the hashed stream means that the entire history's state is considered in the current version's SHA-1 without having to hash the entire history.

This is based on the principle of a [Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree). You can see in this illustration that if you already know the value of **Hash 0** and **Hash 1**, then calculating **Top Hash** is cheaper than rehashing the whole history, particularly if the tree is much larger, but **Top Hash** would still depend on the contents of all the input.

![Merkle tree](https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Hash_Tree.svg/640px-Hash_Tree.svg.png)

This is important because ensuring the provenance of all the code is as important as the current state, but the hash must not be too expensive to calculate.
{{< /spoiler >}}

2. If you knew that the input that generated the SHA-1 hash above was a single word, how would you find the word?
{{< spoiler text="Click to see the answer" >}}
```bash
#!/bin/bash

# Iterate over all the words.
for w in $(cat /usr/share/dict/words); do
	# Calculate the SHA-1 sum of each word.
	sha1="$(echo $w | sha1sum | cut -d ' ' -f 1)"
	# Compare the result to the query and exit with success if correct.
	if [ "$sha1" = "f164acde21f01cf3b2ba4f7dd650d9bf2a699b96" ]; then
		echo $w
		exit 0
	fi
done
exit 1
```

Running this code takes some time and it's only a list of words, so finding a large text that matches a specific hash would be not be feasible.
{{< /spoiler >}}

The next section will go into greater details about how git works and how it is used.
