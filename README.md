# line-edit: Gap Buffer Interactive Demo

* Author: a-tk, Andre Keys

## Overview

This project provides an **interactive terminal editor demo** using a
gap buffer data structure. A gap buffer is often used in text editors to
make insertions and deletions near the cursor efficient.

The program starts with an initial string and allows the user to edit it
directly in the terminal. Keystrokes insert and delete characters, and
special commands (entered after `:`) move the cursor or quit the program.

## Compiling and Using

To build:

```bash
go build ./...
```
To run the demo:

```bash
go run . [-d]
```
The optional ```-d``` flag enables debug mode, which displays the internal
gap in the buffer.

When the program starts, you will be prompted with ```>``` to enter an
initial string. After pressing Enter, the string is loaded into the
gap buffer, and you can begin editing interactively.

### Interactive Controls
 - Printable characters (ASCII 32–126): Insert at the cursor.
 - Backspace / Delete: Remove the character before the cursor.
 - Ctrl-C or Ctrl-D: Exit current string edit, or exit.
 - ```:``` followed by command: Enter a meta-command.

**Meta-Commands**

When you type :, the program will read an additional command:

 - q - Quit the program.
 - b - Move cursor to the beginning.
 - e - Move cursor to the end.
 - ```<number>``` — Move cursor to the given index.
 - ```:``` — Insert a literal colon (:).

If a command is invalid, an error message is shown.

The buffer contents and cursor position are updated on every change.