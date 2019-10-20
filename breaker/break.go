package breaker

import (
	"fmt"

	"github.com/Liberatys/Sanctuary/communication"
)

//TODO: implement a circuit breaker

//RequestBreak, handles the backing off from a request by a given number of times it should try to execute the request and a delay time, gv
func RequestBreak(request communication.Request, repeat int, delay int) (bool, string) {
	max := repeat
	current := 0
	for current < max {
		err, returned_message := request.SendRequest()
		if err != nil {
			fmt.Println(returned_message)
			current++
		} else {
			return true, returned_message
		}
	}
	return false, "Due to an error, the request was not able to find the server or had a server issue."
}
