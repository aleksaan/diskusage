 [![cover.run](https://cover.run/go/github.com/aleksaan/diskusage.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Faleksaan%2Fdiskusage) 
 
# diskusage 
diskusage is a command line utility for calculating folders sizes.
```cmd
Arguments:
   path: c:\go\
   limit: 10
   fixunit: 
   depth: 5

Start scanning
  1.| PATH:   pkg                                   | SIZE:   212.95 Mb   | DEPTH: 1 
  2.| PATH:   pkg\tool                              | SIZE:   123.65 Mb   | DEPTH: 2 
  3.| PATH:   pkg\tool\windows_amd64                | SIZE:   123.65 Mb   | DEPTH: 3 
  4.| PATH:   src                                   | SIZE:    62.58 Mb   | DEPTH: 1 
  5.| PATH:   pkg\windows_amd64_race                | SIZE:    45.89 Mb   | DEPTH: 2 
  6.| PATH:   pkg\windows_amd64                     | SIZE:    38.95 Mb   | DEPTH: 2 
  7.| PATH:   bin                                   | SIZE:    30.45 Mb   | DEPTH: 1 
  8.| PATH:   src\cmd                               | SIZE:    30.11 Mb   | DEPTH: 2 
  9.| PATH:   pkg\tool\windows_amd64\compile.exe    | SIZE:    19.84 Mb   | DEPTH: 4 
 10.| PATH:   bin\godoc.exe                         | SIZE:    14.99 Mb   | DEPTH: 2 
Finish scanning

Overall info:
   Total time: 4.2919743s
   Total dirs: 1129
   Total files: 8690
   Total links: 9819
   Total size: 325.81 Mb
   Total size (bytes): 341640673
   Unaccessible dirs & files: 0

```
## Features
- A primitive tool for getting folder(s) sizes
- Command line environment only
- Supports both folders and disks as arguments
- Recursive passes through subfolders
- Calculates size of each folder
- Analyzes on defined depth of subfolders
- Sets limit how much folders will be printed in a results
- Fast

## Main cons
- No any dummies protection (also pros)
- No any intelligents features (also pros)

## Releases

Releases available as single executable files – just [download latest release](https://github.com/aleksaan/diskusage/releases) for your platform, unpack and run.

## Start on Windows - simple usage

```cmd
diskusage.exe -path "c:/somedir"
```
if you want to get 20 biggest directories in c:/somedir

```cmd
diskusage.exe -path "c:/somedir" -depth 1
```
if you want to get only c:/somedir first level subfolders sizes


## Start on Windows - advanced usage

```cmd
diskusage.exe -path "c:/somedir" -limit 20 -fixunit "Gb" -depth 3
```
if you want to get 20 biggest directories across c:/somedir with a three subfolder's levels depth. All results will be represented in Gb.


where:
```cmd
-path "c:/somedir"
``` 
is a folder name (required)
```cmd 
-limit 20
```
is how much biggest folders will be printed in the results (optional)
```cmd 
-fixunit "Gb"
```
you can choose unit style to representing folder sizes. It can be fixed or dynamic-scaled.
You can use "fixunit" in case you want to compare sizes afterward (optional).
```cmd 
-depth 3
```
is depth of subfolders to analyze (optional)


For integration with a other systems I recommend create a batch file like this:
```cmd
del results.txt
diskusage.exe -path "c:/somedir" > results.txt
```



