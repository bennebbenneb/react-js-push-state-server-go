# Go Server for React Projects Using Push State
This server is designed to handle sites built with [BrowserRouter](https://reacttraining.com/react-router/web/api/BrowserRouter).

## Instructions
1. Build your React project `react-scripts build` and copy it to `/web/`
1. Test on localhost 
    1. Run `dev_appserver.py app.yaml` in the same directory as the `app.yaml` file
    1. Open browser and navigate to http://localhost:8080
1. Deploy to Google App Engine
    1. Run `gcloud app deploy` in the same directory as the `app.yaml` file
    
### Adding more static directories 
During your react build, any folder in `/public/` will create it's own static directory. These folders need 
to be added to the staticPaths array.   
1. Open `main.go` in a text editor
1. Add elements to the `staticPaths` string array

Example
```go
staticPaths := [...]string{"/static/", "/img/", "/media/"}
```