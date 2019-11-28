 [![cover.run](https://cover.run/go/github.com/aleksaan/diskusage.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Faleksaan%2Fdiskusage) 
 
# diskusage 
diskusage is a utility to find top largest directories on the disk.

## !!! 2019.11.28: Version 2.0.3 is now avaliable  !!!


```cmd
About:
   github/aleksaan/diskusage, 2.0.3, Alexander Anufriev, 2019

Arguments:
   path:      d:\_appl\go\src\
   limit:     20
   units:     <dynamic>
   depth:     5
   sort:      size_desc
   tofile:    diskusage_out.txt

Results:
     1.| PATH:   github.com                                | SIZE:   316.65 Mb   | DEPTH: 1 
     2.| PATH:   github.com\aws                            | SIZE:   140.36 Mb   | DEPTH: 2 
     3.| PATH:   github.com\aws\aws-sdk-go                 | SIZE:   140.36 Mb   | DEPTH: 3 
     4.| PATH:   golang.org                                | SIZE:    73.65 Mb   | DEPTH: 1 
     5.| PATH:   golang.org\x                              | SIZE:    73.65 Mb   | DEPTH: 2 
     6.| PATH:   github.com\aws\aws-sdk-go\.git            | SIZE:    66.13 Mb   | DEPTH: 4 
     7.| PATH:   github.com\aws\aws-sdk-go\.git\objects    | SIZE:    65.83 Mb   | DEPTH: 5 
     8.| PATH:   github.com\aleksaan                       | SIZE:    63.05 Mb   | DEPTH: 2 
     9.| PATH:   github.com\aleksaan\diskusage             | SIZE:    60.76 Mb   | DEPTH: 3 
    10.| PATH:   github.com\aws\aws-sdk-go\service         | SIZE:    48.31 Mb   | DEPTH: 4 
    11.| PATH:   golang.org\x\tools                        | SIZE:    32.83 Mb   | DEPTH: 3 
    12.| PATH:   github.com\derekparker                    | SIZE:    32.60 Mb   | DEPTH: 2 
    13.| PATH:   github.com\derekparker\delve              | SIZE:    32.60 Mb   | DEPTH: 3 
    14.| PATH:   github.com\aleksaan\diskusage\dist        | SIZE:    28.30 Mb   | DEPTH: 4 
    15.| PATH:   golang.org\x\sys                          | SIZE:    23.44 Mb   | DEPTH: 3 
    16.| PATH:   golang.org\x\tools\.git                   | SIZE:    23.07 Mb   | DEPTH: 4 
    17.| PATH:   golang.org\x\tools\.git\objects           | SIZE:    22.94 Mb   | DEPTH: 5 
    18.| PATH:   github.com\hajimehoshi                    | SIZE:    22.04 Mb   | DEPTH: 2 
    19.| PATH:   github.com\aws\aws-sdk-go\models          | SIZE:    21.92 Mb   | DEPTH: 4 
    20.| PATH:   github.com\hajimehoshi\go-mp3             | SIZE:    21.81 Mb   | DEPTH: 3 

Overall info:
   Total time: 6.3016798s
   Total dirs: 3674
   Total files: 9646
   Total links: 0
   Total size: 414.98 Mb
   Total size (bytes): 435138161
   Unaccessible dirs & files: 0
```
## Features
- A primitive tool for getting folder(s) sizes
- Comfortable setup (yaml config)
- Supports both folders and disks as arguments
- Recursive passes through subfolders
- Calculates size of each folder
- Analyzes on defined depth of subfolders
- Sets limit how much folders will be printed in a results
- Fast
- Saves results to csv-file

## Main cons
- No any dummies protection (also pros)
- No any intelligents features (also pros)

## Releases

Releases available as single executable files – just [download latest release](https://github.com/aleksaan/diskusage/releases) for your platform, unpack and run.

## Simple usage (Windows example)

Put ```diskusage.exe``` into analyzed directory, run it and get results in ```diskusage_out.txt```

* diskusage_config.yaml will be created with a default settings

## Advanced usage (Windows example)

(Optional) Download, create or save ```diskusage_config.yaml``` near ```diskusage.exe```.

Open ```diskusage_config.yaml``` in text editor to setup diskusage

You will see:
```yaml
# Analyzer options
analyzer: 
  path: 'D:\_docs'
  depth: 5
# Results options
printer:
  limit: 20
  fixunit: Gb
  tofile: diskusage_out.txt
  ```
where:
```yaml
   path: D:\_docs
``` 
is a folder or disk name (required)

```yaml
   depth: 5
```
is depth of subfolders to analyze (optional)

```yaml
   limit: 20
```
is how much biggest folders will be printed in the results (optional)
if you set -limit to 0 it means limitless (no one row be cuted from results). Be warned it might be a huge list of files!
```yaml
   fixunit: Gb
```
you can choose unit style to representing folder sizes. It can be fixed or dynamic-scaled.

If you omit 'fixunit' it means dynamic-scaled units style.

Fixed scale values: b, Kb, Mb, Gb, Tb, Pb.

You can use "fixunit" in case you want to compare sizes afterward (optional).

```yaml
   tofile: diskusage_out.txt
```

File name to save results. If value is empty file will not be created and you will see results in console window with prompt to exit at the end.

Run ```diskusage.exe```
