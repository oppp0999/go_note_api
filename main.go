/* package main

import "fmt"

func main(){
	fmt.Println("go_note_api")
}
 */

 package main

 import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
 )

 func HandleRequest(ctx context.Context) (string, error){
	return "hello YourName", nil
 }

 func main(){
	lambda.Start(HandleRequest)
 }