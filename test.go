//Package test to display package
package test

//TestString is a global variable
var TestString = "Package created"

//PrintTestString returns the package name created
func PrintTestString(name string) string{
	return TestString + "with name" + name
}