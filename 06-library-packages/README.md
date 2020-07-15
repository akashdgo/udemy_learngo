### Librart Packages

-  It's a good practice to name your package as the same as his folder's name.
- The package cannot be run since it does not contain the 'main' package but it can be compiled using `'go build'`
- The package need not be compiled but needs to be imported into another package.
- install the package using 'go install'

![image](https://user-images.githubusercontent.com/28204484/87497537-55b9cc80-c673-11ea-8de3-fe63f21f5020.png)

![image](https://user-images.githubusercontent.com/28204484/87497550-5eaa9e00-c673-11ea-9f8d-35aab182a7b6.png)


- the install file is a compressed file and can be uncompressed
- the install file is a compiled file and so when a package imports this library package and compiles go will get this compiled package from here and inject it into the final compile file.

