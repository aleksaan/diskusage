package config

var ConfigTemplate = `
# How does the program work?
# This one scans directory tree from given path to deeper folders and files and calculates sizes of each folder & files in tree.
# Then, it prints results to text file uses some filters (by type, by depth, by number)

# -----------------------------------------------------------------------------------------------
# analyzerOptions - group of options for managing behavior of analyzer (where, how deep, what's included)
analyzerOptions:

  # path - (optional) name of disk or folder for analyzing inside them 
  # for example: abc/xyz (Unix-style) or C:\ or C:\temp (Windows-style) etc
  # The default value is the working directory
  path: C:\

  # sizeCalculatingMethod - (optional) ndefines how the program calculates folder size
  # Possible values: 
  #   - plain - size of folder includes sizes of nested files only (without nested folders)
  #   - cumulative - (default) size of folder includes sizes of nested files & folders
  sizeCalculatingMethod: cumulative

# -----------------------------------------------------------------------------------------------
# filterOptions - group of options for filtering results
filterOptions:

  # depth - (optional) cuts all objects whose depth (calulates from path in option "path") is more than given "depth"
  # Default value is 5
  depth: 5

  # limit - (optional) defines how much objects will be finally printed
  # If you set limit = 15 then you see only top-15 largest objects in the results
  # Default value is 20
  limit: 20

  # filterByObjectType - (optional) defines type of objects should be filtered
  # Possible values:
  # files - keep only files
  # folders - keep only folders
  # <empty> or folders&files - (default) keep both of them
  filterByObjectType: folders&files

# -----------------------------------------------------------------------------------------------
# printerOptions - group of options for managing formatting results (how much, units, filter by objects, names of results files)
printerOptions:

  # units - (optional) defines unit style for representing folder sizes. It can be fixed or dynamic-scaled.
  # Default value is empty what means dynamic-scaled units style.
  # Fixed scale values: b, Kb, Mb, Gb, Tb, Pb.
  units: 
 
  # toTextFile - (optional) defines name of file for results of program in text format
  # Possible values:
  #   <empty> - results will be outputed to console window
  #   <some file name> - results will be outputed to file
  toTextFile: diskusage_out.txt

  # toYamlFile - (optional) defines name of file for results of program in YAML format
  # Possible values:
  #   <empty> - no YAML file will be generated
  #   <some file name> - results will be outputed to file
  toYamlFile: 
`
