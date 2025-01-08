# Resumable copy
This utility copies content using byte index of source file. It loops based on 
chunk size.

## How to run
```
./resumable-copy copy \
    --src {source-file} \
    --dest {destination-file} \
    --resume-at {byte-index} \
    --chunk-size {chunk-size} \
    --lag {second-every-copy-iteration}
```
If no --dest is supplied, it will create new "destination.txt" file.

Setting --resume-at 10 will write 10th byte and onwards until EOF. If 10th byte 
exist on destination file, it will be overwritten. Therefore, --resume-at 0 will 
overwrite everything on destination.

If chunk size is 4, the copy will repeat
every 4 bytes. The script will exit whenever chunk size is bigger than
actual source file size.

## Example
Copying contents from `source` to `destination` from 10th byte. We set chunk size 
to 4 and the copy occurs every 3 second.
```
./resumable-copy copy \
    --src source.txt \
    --dest destination.txt \
    --resume-at 5 \
    --chunk-size 4 \
    --lag 3
```

The stdout will look like this,
```
go run . copy --src source.txt --dest destination.txt --resume-at 10
source: source.txt
destination: destination.txt
resume at: 10
chunk size: 4
lag: 3


copied byte index 10 to 13
copied byte index 14 to 17
copied byte index 18 to 21
copied byte index 22 to 25
copied byte index 26 to 29
copied byte index 30 to 33
copied byte index 34 to 37
...
``` 