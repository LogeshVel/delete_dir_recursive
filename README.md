## Delete the Directory recursively from the given path

To delete the particular Directory from the given path recursively, by searching in all the directories hierarchy for the given directory from the given path.

  Ex:
    Consider the below folder structure in the given path.
    
    .
    |__ __pycache__
    |__ dir1
    |   |__ somefile.txt
    |__ dir2
    |   |__ __pycache__
    |   |__ child_dir
    |__ somedir
        |__ dir
        |__ __pycache__


Here \_\_pycache__ folder present in almost three places but in the different folders. This Golang program will remove all the \_\_pycache__ folder if the root path is given from where it will look for that dir name in all the folders.

It will recursively search for that foldername in each folder hierarchy level from the given path and deletes that Dir.


##### Usage of main.go:

  -d string
  
    Dir name to find recursively in the given path and delete   
        
  -p string
  
    path to find the dir and delete recursively
        
  -v   
  
    Prints the dir name in all dir while searching for the given dir


##### To delete the Directory named dir1 recursively from the current path

```
go run main.go -p . -d dir1
            or
executable_bin -p . -d dir1
```

##### To enable the verbose

```
go run main.go -p . -d dir1 -v
            or
executable_bin -p . -d dir1 -v
```

# Installation

Download the zip file corresponding to your OS in the release page. As of now the **deldirrecur** app supports only for Linux and Windows.

After downloading, extract the zip file using the below cmd

```
tar -zxvf filename.tar.gz
```

After unzipping execute the following cmd in the same dir where the files have been extracted.

In Windows, open the command prompt with the administrative privilege (Run as Administrator)

```
move deldirrecur.exe "C:/WINDOWS"
```

For Windows after executing the above cmds, close the existing cmd prompt and open new one to use the **deldirrecur** app

In Linux

```
sudo mv deldirrecur /usr/local/bin
```

## after installation

We can use this app by the downloaded bin.

Ex: To delete dir1 from the given path in all the hierarchy level.

```
deldirrecur -p . -d dir1 -v
```

