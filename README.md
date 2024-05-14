# Tests
Test are located in package ```tests```
They divided into unit tests and intagration tests

### Running Tests:

Navigate to the directory containing the test files.
Run the following command in your terminal:

```go test```

### Test Descriptions:

#### Unit tests:

- TestNewAccountCreationRequest - Tests creating a new user account with valid data.
- TestNewAccountCreationRequest2 - Tests creating a new user account with invalid data (invalid email and missing password).
- TestGettingAccountAuthenticationToken - Tests obtaining an authentication token with valid user credentials.
- TestGettingAccountAuthenticationToken2 - Tests obtaining an authentication token with invalid user credentials (non-existent email).
- TestActivatingAccount - Tests activating a user account with a valid activation token.
- TestActivatingAccount2 - Tests activating a user account with an invalid activation token.

#### Integration tests:

- TestInsertingMoviesIntoDatabase - Tests inserting a valid movie entry into the database.
- TestInsertingMoviesIntoDatabaseWithWrongYear - Tests inserting a movie entry with a future year into the database (expected to fail).
- TestInsertingMoviesIntoDatabaseWithWrongRuntime - Tests inserting a movie entry with an invalid runtime format into the database (expected to fail).
- TestMovieDeletionById - Tests deleting a movie by its ID (assuming the ID 5 exists in your database).


