package cmd

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestPath( t * testing.T){
	// files to test
	input := []string{
		"../test.html",
		"test1",
		"test2",
		"index.html",
		"../../index.html",
	}
	
	// expected result
	expected := []bool{
		true,
		false,
		false,
		true,
		false,
	}

	for i:= 0 ; i< len(input) ; i++{
		assert.Equal(t,expected[i], runT(input [i]) ,"input was: "+input[i])
	}		
}

// method to test
func runT(mpath string) bool {
	
if _, err := os.Stat(mpath); err != nil {
   return false 
} 

return true
}


func TestVersion( t * testing.T){
	
	// expected result
	expected := []string{
		"1.0",
		"1.0",
		"1.0",
		"1.0",
	}

	for i:= 0 ; i< len(expected) ; i++{
		assert.Equal(t,expected[i] ,getVersion(),"Failed at %d time ",i)
	}		
}