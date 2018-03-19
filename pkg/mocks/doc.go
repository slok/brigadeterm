/*
Package mocks will have all the mocks of the application, we'll try to use mocking using blackbox
testing and integration tests whenever is possible.
*/
package mocks // import "github.com/slok/brigadeterm/pkg/mocks"

// Service mocks.
//go:generate mockery -output ./service/repository/brigade -outpkg brigade -dir ../service/repository/brigade -name Repository
