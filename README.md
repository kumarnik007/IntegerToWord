# IntegerToWord

Description:
    1.	The web server is designed to handle HTTP GET /identity and HTTP POST /convert requests.
    2.	HTTP POST /convert request body is expected to be a JSON object like below :
    {"value":978
    } 
    3.	Go programming language has been used to develop the web server.
    4.	The server is designed to listen on port 8081.

Pre-requisites for deploying the web server:
    1.	Ensure that Go version (at least) go1.14.6 is installed on the workstation where the server needs to be deployed. The version of Go installed can be checked by executing the below command in terminal :
        •	go version
    2.	The workstation where the server is to be deployed should have port 8081 available.

Steps to Deploy:
    1.	Unzip the “IntegerToWord.zip” file provided.
    2.	Open a terminal and navigate to the unzipped directory “IntegerToWord”
    3.	Compile all the source code files (*.go) by providing the below command
        •	go build .
    4.	Step (3) creates an executable file in the same directory.
    5.	The web server can be run by executing the executable file as below:
        •	./IntegerToWord (for linux OS and macOS)
        •	IntegerToWord.exe (for Windows OS)

Unit Testing:
I have included a test file “Handlers_test.go” containing the unit test cases which were covered a part of the web server implementation. The unit test cases present in this test file can be executed by following the below steps:
    1.	Open a terminal and navigate to the unzipped directory “IntegerToWord”
    2.	Run the unit test cases by providing either of the below commands 
        •	Execute the tests without printing any logs from the testing package
            	go test  
        •	Execute and log all tests as they are run 
            	go test –v
        •	Execute tests and do not start new tests after the first test failure
            	go test -v -failfast
    3.	Check for the status of each of the unit test case in the terminal output, all the test cases should pass. A test status of ‘ok’ or ‘FAIL’ along with the package name and elapsed time is printed at the end.
    4.	Verify the result received as part of the test case against the input value

Sample HTTP Requests:
    1.	GET /identity
        •	GET /identity HTTP/1.1
    Host: localhost:8081
        •	Sample URL : http://localhost:8081/identity
    2.	POST /convert
        •	POST /convert HTTP/1.1
        Host: localhost:8081
        Content-Type: application/json

        {"value":978
        }

Assumptions/Limitations:
    1.	For POST /convert API, the server returns a 400 BAD REQUEST with appropriate error message in the following scenarios :
        a.	All positive integer greater than 999999999 (any 10 digit integer).
        b.	All negative integers (values less than 0).
        c.	Any decimal, character or string values.
    2.	For POST /convert API, the input JSON is not expected to have any preceding 0’s in the contained positive integer value.
