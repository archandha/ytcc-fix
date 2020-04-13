This is a very simple CLI tool to fix my closed caption files downloaded from
YouTube to import them into my favorite CC editor, 
which is https://github.com/SubtitleEdit/subtitleedit without pain.

The originally downloaded files have overlapping cc times, i. e. the next
line starts before the former ends. This can be annoying to be edited.

There is no sophisticated algorithm to determine the run length. It just takes
the start of the next cc line minus one millisecond. If next millisecond is
already "000", it keeps "000", so there might arise an overlap.

This does not take into account that there might be gaps between following
cc lines. This will be something to determine in later releases ;)

... and yes, this lines are my first steps with golang, so be patient with me.
