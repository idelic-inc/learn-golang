# Golang Lunch and Learn Hackathon

Base repository for our Golang Lunch and Learn seminar.

## Format
 
 1. Start with the base project: a webserver with a single route that uploads a file to S3.
 2. Follow the suggested project requirements, or be creative and build something different.
 3. We are working individually on the project, but learning as a group. If you get ahead of others, it would be great to stop and help them catch up.
 4. The intention is to learn how you would normally learn (the way that you are best at learning), which for most of us involves a combination of research (Google), trial and error, and asking others.

## Project Overview

We are going to build a very simplistic Instagram backend (though once again, feel free to use this time in whichever way is best for you).

The idea is to learn Golang the way it is actually used, which is typically to build robust services. 

### Test It

 1. Build the executable: `go build`
 2. Run the executable: `./golang-learn`
 3. Upload a file using your favorite HTTP client. The `upload_file.sh` script in the repo will run a curl command to do this for you.

    ```
    upload_file.sh myfile.txt
    ```

### The Code

#### Resources

[Gin Web Framework](https://github.com/gin-gonic/gin)

[Golang AWS SDK Docs](https://docs.aws.amazon.com/sdk-for-go/api/aws/)


## Suggested Requirements

 1. Only allow `.png` files to be uploaded.
 2. Only allow files under a certain size limit to be uploaded.
 3. Create unit tests for the features from steps 1 and 2 and run them using `go test`
 4. Add a `GET` route to download files by name.
 5. Allow files to be associated with a user when uploaded.
 6. Update or add a `GET` route to get files by user and name.
 7. Add a route to fetch all file names for a particular user.
 8. Add a filter feature that may optionally be applied to uploaded files.
 9. Allow documents to have labels and store them using S3 metadata. 

## Resources
 * [Go By Example](https://gobyexample.com/)
   Often the easiest way to get working code snippets.
 * [How to Write Go Code](https://golang.org/doc/code.html)
   How to organize your code, compile it, and run it.
 * [Effective Go](https://golang.org/doc/effective_go.html)
   Common, idiomatic patterns used when writing Go.
 * [Go Modules](https://blog.golang.org/using-go-modules)
   Understand how the new Go module system works.

