### Exporting

A lib-pkg can have func the name of which can start with an uppercase or a lower case.
- Upper case mean : the func is being exported and any other package importing the lib-pkg can use this func
- lower case means : the func is not being exported and is restricted to use by the package alone [package scope]

![image](https://user-images.githubusercontent.com/28204484/87498971-beef0f00-c676-11ea-92d0-e09e4d90c932.png)

 - New(text string) : This func can be used by any package importing "error" package
    - Similary 'Error()' which is a **method** (a function but attached to a type).
- But the struct 'errorString' can only be used by package error and not by any other package (package scope).
