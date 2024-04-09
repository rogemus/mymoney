package utils

import "errors"

// Budget
var Budget404Err = errors.New("Budget not found")

// Generic
var Generic400Err = errors.New("Invalid request")
var Generic401Err = errors.New("Unauthorized")

// Auth
var AuthInvalidHeader = errors.New("Invalid Authentaction Header")
var AuthIvalidPass = errors.New("Invalid Password or Email")

// User
var User404Err = errors.New("User not found")
