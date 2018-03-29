# diskusage 
diskusage is a command line utility for getting information about usage of disk(s) or folder(s) space.

## Features
- A primitive tool for getting usage of disk(s) / folder(s) space
- Command line environment only
- Supports both folders and disks in arguments
- Recursive pass through folders tree on defined disk(s) / folders(s)
- Calculate size of each folder
- Print list of (sub)folders with a biggest sizes
- Set limit to number folders in printing

## Start on Windows

```cmd
diskusage.exe -path "C:/Temp; D:/" -limit 20
```
where:
```cmd
-path "C:/Temp; D:/"
``` 
is set of disk(s) / folder(s) separated by semicolon (required)
```cmd 
-limit 20
```
is how much max-sized folders you want to see in the results (optional)



