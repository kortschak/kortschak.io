---
title: Computer Setup
linktitle: Computer Setup
toc: true
type: docs
date: "2019-05-05T00:00:00+01:00"
draft: false
menu:
  git_workshop:
    parent: Git Workshop
    name: Computer Setup
    weight: 1

# Prev/next pager order (if `docs_section_pager` enabled in `params.toml`)
weight: 1
---

For those running OSX or Ubuntu we will be able to easily run today’s session on your own machines. Many bioinformaticians use these operating systems so the suite of tools is relatively mature. Windows will require some additional work. In order to correctly configure your computer for today’s session, follow the section that is relevant for your operating system.

## OSX

To set up your own computer for today's session, follow these instructions.
Copying and pasting the given code may be the easiest way to make sure everything works.

1. Open a Terminal `Applications -> Utilities -> Terminal`.
2. If you don't have `homebrew` installed, type the following to your terminal. If you *do* have `homebrew` installed, please skip to the next step.

```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Install git.

```
brew update
brew install git
```

## Ubuntu

To set up your own computer for today's session, follow these instructions.
Copying and pasting the given code may be the easiest way to make sure everything works.

Open a Terminal and enter these lines one at a time. Enter `y` where required.

```
sudo apt update
sudo apt install -y git
```

## Windows

To install git on Windows navigate to the [git for windows page](https://git-for-windows.github.io/), click on the download button and follow the installer instructions.
