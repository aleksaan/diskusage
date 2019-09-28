 [![cover.run](https://cover.run/go/github.com/aleksaan/diskusage.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Faleksaan%2Fdiskusage) 
 
# diskusage 
diskusage is a utility for calculating folders sizes.
```cmd
About:
   github/aleksaan/diskusage, 2.0.2, Anufriev Alexander, 2019

Arguments:
   path:      c:\Windows
   limit:     20
   units:     <dynamic>
   depth:     5
   sort:      size_desc
   tofile:    out.txt

Results:
     1.| PATH:   WinSxS                                                    | SIZE:    11.16 Gb   | DEPTH: 1 
     2.| PATH:   System32                                                  | SIZE:     3.98 Gb   | DEPTH: 1 
     3.| PATH:   SysWOW64                                                  | SIZE:     1.21 Gb   | DEPTH: 1 
     4.| PATH:   Installer                                                 | SIZE:  1007.59 Mb   | DEPTH: 1 
     5.| PATH:   servicing                                                 | SIZE:   963.02 Mb   | DEPTH: 1 
     6.| PATH:   System32\DriverStore                                      | SIZE:   948.36 Mb   | DEPTH: 2 
     7.| PATH:   System32\DriverStore\FileRepository                       | SIZE:   944.60 Mb   | DEPTH: 3 
     8.| PATH:   MEMORY.DMP                                                | SIZE:   848.83 Mb   | DEPTH: 1 
     9.| PATH:   servicing\LCU                                             | SIZE:   845.45 Mb   | DEPTH: 2 
    10.| PATH:   assembly                                                  | SIZE:   833.59 Mb   | DEPTH: 1 
    11.| PATH:   Microsoft.NET                                             | SIZE:   802.01 Mb   | DEPTH: 1 
    12.| PATH:   Panther                                                   | SIZE:   624.18 Mb   | DEPTH: 1 
    13.| PATH:   System32\DriverStore\FileRepository\nv_ref_pubwu.inf      | SIZE:   558.38 Mb   | DEPTH: 4 
    14.| PATH:   servicing\LCU\Package_for_RollupFix~31bf3856ad364e35      | SIZE:   445.44 Mb   | DEPTH: 3 
    15.| PATH:   SoftwareDistribution                                      | SIZE:   421.71 Mb   | DEPTH: 1 
    16.| PATH:   WinSxS\Backup                                             | SIZE:   402.11 Mb   | DEPTH: 2 
    17.| PATH:   servicing\LCU\Package_for_RollupFix~31bf3856ad364e35      | SIZE:   400.01 Mb   | DEPTH: 3 
    18.| PATH:   SoftwareDistribution\Download                             | SIZE:   391.21 Mb   | DEPTH: 2 
    19.| PATH:   Fonts                                                     | SIZE:   361.27 Mb   | DEPTH: 1 
    20.| PATH:   assembly\NativeImages_v4.0.30319_64                       | SIZE:   359.89 Mb   | DEPTH: 2 

Overall info:
   Total time: 1m51.4143664s
   Total dirs: 45275
   Total files: 189145
   Total links: 0
   Total size: 23.33 Gb
   Total size (bytes): 25046502289
   Unaccessible dirs & files: 336

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

Put ```diskusage.exe``` into analyzed directory, run it and get results in ```out.txt```

* config.yaml will be created with a default settings

## Advanced usage (Windows example)

(Optional) Download, create or save ```config.yaml``` near ```diskusage.exe```.

Open ```config.yaml``` in text editor to setup diskusage

You will see:
```yaml
# Analyzer options
analyzer: 
  path: 'D:\_docs'
  depth: 5
# Result's options
printer:
  limit: 20
  fixunit: Gb
  tofile: out.txt
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
   tofile: out.txt
```

File name to save results. If value is empty file will not be created and you will see results in console window with prompt to exit at the end.

Run ```diskusage.exe```
