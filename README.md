# Sample Requests for POSTMAN

```
POST /submit

In headers, add
Content-Type as application/json
Body -> raw -> JSON

Requst:
{
    "UserID": "Test",
    "Code": "#include <iostream>\nint main(){\n    int n = 0;\n    std::cin >> n;\n    std::cout << n*(n-1);\n    return 0;\n} ",
    "QuestionID": 123
}

Response:
{
    "message": "submitted",
    "submissionID": "37015e171abb11ebb205809133c1ccfd"
}

POST /status

In headers, add
Content-Type as application/json
Body -> raw -> JSON

Request: 
{
    "SubmissionID": "37015e171abb11ebb205809133c1ccfd" #ID which you get on previous POST request
}
Response:
{
    "status": "WRONG ANSWER"
}