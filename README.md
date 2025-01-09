# Resumable copy
This utility copies content using byte index of source file. It loops based on 
chunk size.

## How to run
```bash
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
```bash
./resumable-copy copy \
    --src source.txt \
    --dest destination.txt \
    --resume-at 5 \
    --chunk-size 4 \
    --lag 3
```

The stdout will look like this,
```
❯ ./resumable-copy copy \
    --src source.txt \
    --dest destination.txt \
    --resume-at 5 \
    --chunk-size 4 \
    --lag 3
source: source.txt
destination: destination.txt
resume at: 5
chunk size: 4
lag: 3


copied byte index 5 to 8
copied byte index 9 to 12
copied byte index 13 to 16
copied byte index 17 to 20
copied byte index 21 to 24
copied byte index 25 to 28
copied byte index 29 to 32
copied byte index 33 to 36
...
``` 

## Make syntax
Build binary,
```bash
make build
```

Delete default destination (`destination.txt`)
```bash
make clean 
```

Delete customized destination (`my_own_destination`)
```bash
make clean my_own_destination
```
