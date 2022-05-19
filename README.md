# DUCK - (D)isk(U)sage(CK) 
Duck is a very fast utility to find largest directories or files
<br><img src="./img/duck.jpg" width="200"><br>
Illustrator: Ekaterina [[t.me/@kateUV](https://t.me/kateUV)]

## To Elon Musk
I hope you will see it and I have couple of quite well ideas to chat them with you

## Features
- Gathers directories/files sizes
- Finds top of largest directories/files
- Very fast
- JSON compatible (for a service use)
- Human readable mode of output (for a console use)
- More accuracy than FAR manager

## Releases
Releases available as single executable files – just [download latest release](https://github.com/aleksaan/diskusage/releases) for your platform, unpack and run.

## How it works
Since version **2.8.0** utility has a two parts:
- ```Duck_a``` duck says **analyse**
- ```Duck_f``` duck says **find**

```Duck_a``` gathers sizes of all directories and files under specific path.
```Duck_f``` takes results of ```Duck_a``` and looking for top of largest objects among them.

You can scanning 1Tb disk only once by ```Duck_a``` (*some minutes*) and then many times looking for largest objects by ```Duck_f``` with different parameters (*some milliseconds*).

## ```Duck_a``` utility

```Duck_a``` calculates sizes of directories and files. 

**Parameters**:
- ```-path=c:\temp``` - starting point to analyse
- ```-depth=2``` - depth of analysis (it's a filter to exclude directories which level is more than 'depth'. It does not reduce time of analysis!)
- ```-hr``` - human readable results representation (text format), if omit that means JSON format

By default program outputs results to ```console```.

**Example a.1**. Scanning without limit by depth from c:\temp and saving results in JSON format to ```results_a.txt```
   
```duck_a.exe -depth=0 -path c:\temp\ > .\results_a.txt```

**Example a.2**. Scanning with limit by depth=2 from c:\temp and saving results in human readable format to ```console```
   
```duck_a.exe -depth=2 -path c:\temp\ -hr``` 

## ```Duck_f``` utility

```Duck_f``` takes results of ```Duck_a```, iterates over them and finds top largest directories or files

**Parameters**:
- ```-top=20``` - how much directories or files will be founded
- ```-depth=2``` - depth of analysis inside of results ```Duck_a```
- ```-filter=df``` - filter by objects types (```f``` - files only, ```d``` - directories only, ```df``` - both of them)
- ```-size=c``` - method of calculating directories size (```c``` clean size (excludes sizes of subdirectories) or ```f``` - full size (inludes subdirectories))
- ```-path=abc``` - not the same what this parameters means in ```duck_a```. It's a filter by part of the path (will be outputed all rows which path includes this one)
- ```-hr``` - human readable results representation (text format), if omit that will be JSON format

By default program outputs results to ```console```.

**Example f.1**. Searching top-10 largest directories or files on depth 2 and outputing results as JSON to file
   
```duck_f.exe -depth=2 -size=c -top=10 -filter=df < .\results_a.txt > .\results_f.txt```

**Example f.2**. Searching top-12 largest directories or files on depth 3 and outputing results in human readable format to ```console```
   
```duck_f.exe -depth=3 -size=c -top=12 -filter=d -hr < .\results_a.txt```

**Example f.3**. Like as ```Example f.2``` but with filtering by path of file (for example, print only dir or files contains `.git` in their path & names)
                                                                         
```duck_f.exe -depth=2 -size=c -top=12 -filter=d -path=.git -hr < .\results_a.txt```

**So there are results of Example f.3**
```-------------------
Arguments:
   filter:    d
   depth:     3
   top:       12
   hr:        true
   size:      c
-------------------
Results:
     1.| PATH:   diskusage\.git\hooks     | FULL SIZE:    22.89 Kb   | CLEAN SIZE:    22.89 Kb   | DEPTH: 3
     2.| PATH:   statusek\.git\hooks      | FULL SIZE:    22.89 Kb   | CLEAN SIZE:    22.89 Kb   | DEPTH: 3
     3.| PATH:   statusek\.git            | FULL SIZE:   114.22 Mb   | CLEAN SIZE:     5.22 Kb   | DEPTH: 2
     4.| PATH:   diskusage\.git           | FULL SIZE:    22.67 Mb   | CLEAN SIZE:     5.08 Kb   | DEPTH: 2
     5.| PATH:   statusek\.git\logs       | FULL SIZE:     5.30 Kb   | CLEAN SIZE:     2.17 Kb   | DEPTH: 3
     6.| PATH:   diskusage\.git\logs      | FULL SIZE:     1.79 Kb   | CLEAN SIZE:   741.00 b    | DEPTH: 3
     7.| PATH:   diskusage\.git\info      | FULL SIZE:   240.00 b    | CLEAN SIZE:   240.00 b    | DEPTH: 3
     8.| PATH:   statusek\.git\info       | FULL SIZE:   240.00 b    | CLEAN SIZE:   240.00 b    | DEPTH: 3
     9.| PATH:   diskusage\.git\objects   | FULL SIZE:    22.64 Mb   | CLEAN SIZE:     0.00 b    | DEPTH: 3
    10.| PATH:   diskusage\.git\refs      | FULL SIZE:   155.00 b    | CLEAN SIZE:     0.00 b    | DEPTH: 3
    11.| PATH:   statusek\.git\objects    | FULL SIZE:   114.19 Mb   | CLEAN SIZE:     0.00 b    | DEPTH: 3
    12.| PATH:   statusek\.git\refs       | FULL SIZE:   196.00 b    | CLEAN SIZE:     0.00 b    | DEPTH: 3
  ```
* How you can see results are sorted by ```CLEAN SIZE``` (not included sizes of subdirectories). ```FULL SIZE``` is not sorted and not the same as ```CLEAN SIZE```.

**Note about ```FULL SIZE``` and ```CLEAN SIZE```**
   
For example, if you have directories:
- ```A (100Mb)\B (70Mb)\C (60Mb)```

then ```CLEAN SIZE``` of this dirs will be:
- ```A``` - ```30Mb``` (exclude size of ```B```)
- ```B``` - ```10Mb``` (exclude size of ```C```)
- ```C``` - ```60Mb``` (the same as ```FULL SIZE``` because no any subdirs inside)

