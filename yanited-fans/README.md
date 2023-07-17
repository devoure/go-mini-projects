# 🌐 Yanited Web Server
> :bulb: **Project** 2 / 12
## 💬 Description
> This is a simple web server created by Go that handles GET and POST requests. It listens to this request at port 8080
![desc](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/Yanited-server.png)
## 📜 More
> 4 different URL Paths are handled by different http handlers written in Go
  1. / path
     > This path "localhost:8080/" is handled by the fileServer handle object and processes the index.html file
     > which provides an overview of our simple website and links to other pages.
     <details>
      <summary>Index.html Screenshots</summary>
      
      ![index.html screenshot1](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/index1.png)
      ![index.html screenshot2](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/index2.png)
     </details>
  1. /fpl path
     > This path "localhost:8080/fpl" is handled by the fplPathHandler and processes the fpl.html page thats shows table of members' points in the clubs fantasy league.

     <details>
      <summary>
       fpl.html screenshots
      </summary>

      ![fpl.html](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/fpl.png)
     </details>
  1. /transfers path
     > This path "localhost:8080/transfers" is handled by the transfersPathHandler and processes the transfers.html page thats shows transfer news

     <details>
      <summary>
       transfers.html screenshots
      </summary>

      ![fpl.html](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/transfer.png)
     </details>
  1. /join path
     > This path "localhost:8080/join" is handled by the joinPathHandler and processes the join.html page thats a form that takes a user's details. This handler also processes the POST request to the server.

     <details>
      <summary>
       join.html screenshots
      </summary>

      ![join.html Screenshot 1](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/formdetails.png)
           ![join.html Screenshot 2](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/response.png)
     </details>

## 🔧 Code Setup
   1. Clone the repo
        ```
            git clone 

        ```
   1. Move into the project file
        ```bash
            cd yanited-fans

        ```
   1. Run the code
        ```bash
            go run main.go

        ```
## 💻🏃‍♂️ Running Code Snippet
![code](https://github.com/devoure/go-mini-projects/blob/main/yanited-fans/static/images/coderun.png)





