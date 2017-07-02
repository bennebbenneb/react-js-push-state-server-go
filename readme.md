# Go Server for React Projects Using Push State
This server is designed to handle sites built with [BrowserRouter](https://reacttraining.com/react-router/web/api/BrowserRouter).

## Instructions
1. Build your React project and copy it to `/web/`
1. Test on localhost 
    1. Run `dev_appserver.py app.yaml` in the same directory as the `app.yaml` file
1. Deploy to Google App Engine
    1. Run `gcloud app deploy` in the same directory as the `app.yaml` file
    
### Adding more static directories 
During your react build, any folder in `/public/` will create it's own static directory. These folders need 
to be added to the staticPaths array.   
1. Open `main.go`
1. Increase the size of the staticPaths array to account for your static directories
    1. `var staticPaths [3]string` 
1. Add an element to the `staticPaths` string array
    1. `staticPaths[1] = "/img/"`
    1. `staticPaths[2] = "/media/"`

       
Result
```go
var staticPaths [3]string
staticPaths[0] = "/static/"
staticPaths[1] = "/img/"
staticPaths[2] = "/media/"
```