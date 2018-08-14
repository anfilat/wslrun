# wslrun
Helper for run linux utilities in windows. It converts windows paths to wsl paths

For example
```
wslrun diff --color c:\projects\hello\index.js c:\projects\oldHello\index.js
```
will launched in wsl as
```
diff --color /mnt/c/projects/hello/index.js /mnt/c/projects/oldHello/index.js
```