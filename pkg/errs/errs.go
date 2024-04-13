package errs

import "errors"

// Budget
var Budget404Err = errors.New("Budget not found")

// Generic
var Generic400Err = errors.New("Invalid request")
var Generic401Err = errors.New("Unauthorized")
var Generic422Err = errors.New("Invalid or missing request params")

// Auth
var AuthInvalidHeader = errors.New("Invalid Authentaction Header")
var AuthIvalidPass = errors.New("Invalid Password or Email")
var AuthTokenNotFound = errors.New("Token not found in db")

// User
var User404Err = errors.New("User not found")
var UserInvalidEmail = errors.New("Invalid User email")
var UserInvalidUsername = errors.New("Invalid User username")
var UserInvalidPassword = errors.New("Invalid User Password")


// Transaction
var Transaction404Err = errors.New("Transaction not fount")
var TransactionInvalidAmount = errors.New("Invalid Transaction amount")
var TransactionInvalidBudgetId = errors.New("Invalid Budget id")

