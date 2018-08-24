 [![cover.run](https://cover.run/go/github.com/aleksaan/diskusage.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Faleksaan%2Fdiskusage) 
 
# diskusage 
diskusage is a command line utility for getting information about usage of disk(s) or folder(s) space.
```cmd
Input arguments:
   path: C:/Windows; D:/Programs
   limit: 10
   fixunit: 
Start scanning
  1.| DIR: C:/Windows                                                               | SIZE: 51.13 Gb
  2.| DIR: C:/Windows/winsxs                                                        | SIZE: 26.74 Gb
  3.| DIR: C:/Windows/System32                                                      | SIZE: 12.25 Gb
  4.| DIR: C:/Windows/System32/DriverStore                                          | SIZE: 8.55 Gb
  5.| DIR: C:/Windows/System32/DriverStore/FileRepository                           | SIZE: 8.55 Gb
  6.| DIR: D:/Programs                                                              | SIZE: 5.06 Gb
  7.| DIR: D:/Programs/DVD9_Office_Select_Edition_2016_W32_RuEn_MLF_SPecialiST      | SIZE: 4.90 Gb
  8.| DIR: C:/Windows/Installer                                                     | SIZE: 3.62 Gb
  9.| DIR: C:/Windows/SysWOW64                                                      | SIZE: 2.28 Gb
 10.| DIR: C:/Windows/assembly                                                      | SIZE: 1.92 Gb
Finish scanning
Total time: 34.0898503s
``` 
## Features
- A primitive tool for getting usage of disk(s) / folder(s) space
- Command line environment only
- Supports both folders and disks in arguments
- Recursive pass through folders tree on defined disk(s) / folders(s)
- Calculate size of each folder
- Print list of (sub)folders with a biggest sizes
- Set limit to number folders in printing
- Fast

## Main cons
- No any dummies protection (also pros)
- No any intelligents features (also pros)

## Releases

Releases available as single executable files – just [download latest release](https://github.com/aleksaan/diskusage/releases) for your platform, unpack and run.

## Start on Windows

```cmd
diskusage.exe -path "C:/Temp; D:/" -limit 20 -fixunit "Gb"
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
```cmd 
-fixunit "Gb"
```
is a fixed unit of dir-size for a results (optional). If this parameter doesn't set then you get dynamic-scaled results in a more comfort units for each folder. You can use "fixunit" in case you want to compare sizes afterward.

For integration with a other systems I recommend create a batch file like this or more complex if you want:
```cmd
del results.txt
diskusage.exe -path "C:/" -limit 20 >> results.txt
rem pause
rem see to results.txt
```



