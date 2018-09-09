 [![cover.run](https://cover.run/go/github.com/aleksaan/diskusage.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Faleksaan%2Fdiskusage) 
 
# diskusage 
diskusage is a command line utility for getting information about usage of disk(s) or folder(s) space.
```cmd
Parsing input arguments
Arguments:
   path: d:/go; d:/Books
   limit: 20
   fixunit: 
   depth: 2
Start scanning
  1.| DIR: d:/go            | SIZE: 325.72 Mb   | DEPTH: 1 
  2.| DIR: d:/go/pkg        | SIZE: 212.88 Mb   | DEPTH: 2 
  3.| DIR: d:/go/src        | SIZE:  62.57 Mb   | DEPTH: 2 
  4.| DIR: d:/go/bin        | SIZE:  30.44 Mb   | DEPTH: 2 
  5.| DIR: d:/Books/Chess   | SIZE:  14.01 Mb   | DEPTH: 2 
  6.| DIR: d:/Books         | SIZE:  14.01 Mb   | DEPTH: 1 
  7.| DIR: d:/go/api        | SIZE:   6.41 Mb   | DEPTH: 2 
  8.| DIR: d:/go/test       | SIZE:   5.11 Mb   | DEPTH: 2 
  9.| DIR: d:/go/doc        | SIZE:   4.00 Mb   | DEPTH: 2 
 10.| DIR: d:/go/misc       | SIZE:   3.82 Mb   | DEPTH: 2 
 11.| DIR: d:/go/lib        | SIZE: 358.25 Kb   | DEPTH: 2 
Finish scanning
Total time: 272.0156ms
```
## Features
- A primitive tool for getting folder(s) sizes
- Command line environment only
- Supports both folders and disks as arguments
- Recursive pass through subfolders
- Calculate size of each folder
- Analyze for a nedeed depth of subfolders
- Set limit how much folders will be printed in a results
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
if you want to get 20 biggest directories in a d:/somedir

```cmd
diskusage.exe -path "c:/somedir" -depth 1
```
if you want to get only d:/somedir size

```cmd
diskusage.exe -path "c:/somedir; d:/otherdir"
```
if you want to calculate size each of them


## Start on Windows - advanced usage

```cmd
diskusage.exe -path "c:/somedir; d:/otherdir" -limit 20 -fixunit "Gb" -depth 3
```
if you want to get 20 biggest directories across c:/somedir and d:/otherdir with a three subfolder's levels depth. All results will be represented in Gb.


where:
```cmd
-path "C:/Temp; D:/"
``` 
is list of disk(s) / folder(s) separated by semicolon (required)
```cmd 
-limit 20
```
is how much biggest folders will be printed in the results (optional)
```cmd 
-fixunit "Gb"
```
you can shoose unit style to representing folder sizes. It can be fixed or dynamic-scaled.
You can use "fixunit" in case you want to compare sizes afterward (optional).
```cmd 
-depth "2"
```
is depth of subfolders to analyze (optional)


For integration with a other systems I recommend create a batch file like this:
```cmd
del results.txt
diskusage.exe -path "c:/somedir" > results.txt
```



