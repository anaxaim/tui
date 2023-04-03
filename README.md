# TUI - Terraform User Interface

TUI is a web-based UI for Terraform that makes it easy to manage Terraform modules.

Prerequisites
-------------
To run this application, you need to have the following software installed on your system:
* Go 1.20+
* Node.js 12+
* NPM 6+
* Docker with docker-compose

Getting Started
---------------
1. To start the application, run the following command from the project root directory:
    ```
    make up
    ```
2. Once all containers are running, the application can be opened in a browser at: 
   ```
   http://localhost:8080/
   ```

Usage
---------------
1. Create a new user or log in using the provided credentials. 
2. After signing in, you can sign out of the site. 
3. Users can create, modify, and delete modules. 
4. Users can import module content.
