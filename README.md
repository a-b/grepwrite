# Grep Write
Write grep -n output back to files

# Why?

Make mass editing command line friendly.

Programs like `grep`, `ag` or `ripgrep` made search across the project
trivial. Vim users enjoy populating search results to the quickfix list, while
mass editing such files require additional effort.

In order to populate vim quickfix list there is a [special format](https://vimhelp.org/options.txt.html#%27grepformat%27):

```
path_to_file:line_number:column content of the line
```

Commands like `grep -n pattern .` or `rg pattern --vimgrep` outputting results
that way.

General idea:
`grep -n pattern . | sed 's/pattern/new_pattern' | grepwrite`
This will write all your changes back to files.

# Options

-d  Dry run.
-i  Interactive mode. Confirm each change, or all changes per file.

# vim:tw=78
