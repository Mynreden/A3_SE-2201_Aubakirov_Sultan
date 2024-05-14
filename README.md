# Tests
Test are located in package ```tests```
They divided into unit tests and intagration tests

### Running Tests:

Navigate to the directory containing the test files.
Run the following command in your terminal:

```go test```

### Test Results Images
![image](https://github.com/Mynreden/A3_SE-2201_Aubakirov_Sultan/assets/110660562/31e144f9-93d3-4834-bec3-7a6075d3adc6)

![image](https://github.com/Mynreden/A3_SE-2201_Aubakirov_Sultan/assets/110660562/0cf0512b-497b-4dca-a226-0cf5044862d2)



### Test Descriptions:

#### Unit tests:

- TestAccountCreationRequest - Tests the creation of a user account by sending a POST request with valid data and asserts the correctness of the response.
- TestIncorrectAccountCreationRequest - Tests the creation of a user account with intentionally incorrect data to validate the error messages in the response.
- TestLogin - Tests the user login process by sending a POST request with valid credentials and asserts the correctness of the authentication token received in the response.
- TestIncorrectLogin - Tests the user login process with intentionally incorrect credentials to validate the error message in the response.
- TestValidToken - Tests the activation of a user account using a valid token and asserts the correctness of the error message in the response.
- TestInvalidToken - Tests the activation of a user account using an invalid token and asserts the correctness of the error message in the response.

#### Integration tests:

- TestInsertingMoviesIntoDatabase - Tests inserting a valid movie entry into the database.
- TestInsertingMoviesIntoDatabaseWithWrongYear - Tests inserting a movie entry with a future year into the database (expected to fail).
- TestInsertingMoviesIntoDatabaseWithWrongRuntime - Tests inserting a movie entry with an invalid runtime format into the database (expected to fail).
- TestMovieDeletionById - Tests deleting a movie by its ID (assuming the ID 5 exists in your database).


