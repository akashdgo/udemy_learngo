### Exporting

A lib-pkg can have func the name of which can start with an uppercase or a lower case.
- Upper case mean : the func is being exported and any other package importing the lib-pkg can use this func
- lower case means : the func is not being exported and is restricted to use by the package alone [package scope]

![image](https://user-images.githubusercontent.com/28204484/87498971-beef0f00-c676-11ea-92d0-e09e4d90c932.png)

 - New(text string) : This func can be used by any package importing "error" package
    - Similary 'Error()' which is a **method** (a function but attached to a type).
- But the struct 'errorString' can only be used by package error and not by any other package (package scope).

#### Example (pkg. errors used in pkg mail)

![image](https://user-images.githubusercontent.com/28204484/87499996-227a3c00-c679-11ea-92d4-7ba9c49e267b.png)

![image](https://user-images.githubusercontent.com/28204484/87499913-ee068000-c678-11ea-9918-e3f6a382ce81.png)

`Recommendation` 
```
"mail" package is also another library package. Yet it uses the errors package, which is also a library package. A single library can't contain everything. One shouldn't create packages for everything. Each package should be specific on its own. Each one should be responsible for one or a few things at most. "mail" (package) is dealing with "mailing". "errors" package is dealing with "errors", and "fmt" package is dealing with "formatting", and "printing", and so on, so on. So, when creating a package, its always recommended to keep it concise, and to the point.
```
