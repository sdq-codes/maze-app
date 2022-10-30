package main

import "log"
import "github.com/sdq-codes/maze-app/server"

var (
	//Db *gorm.DB,
	AllAnswers [][]string
)

func main() {
	/*err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://2a8c18dc646c44fa8698c7b8d18502e6@o396822.ingest.sentry.io/5887099",
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "",
		Release:     "my-project-name@1.0.0",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}*/
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	//defer sentry.Flush(2 * time.Second)
	srv := server.Server()

	Db = config.InitialiseDb()

	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
	//maze1 := map[string]interface{}{
	//	"forward": "tiger",
	//	"left": map[string]interface{}{
	//		"forward": map[string]interface{}{
	//			"upstairs": "exit",
	//		},
	//		"left": "exit",
	//	},
	//	"right": map[string]interface{}{
	//		"forward": "dead end",
	//	},
	//}
	//maze2 := map[string]interface{}{
	//	"forward": "tiger",
	//	"left": map[string]interface{}{
	//		"forward": map[string]interface{}{
	//			"upstairs": "exit",
	//		},
	//		"left": "dragon",
	//	},
	//	"right": map[string]interface{}{
	//		"forward": "dead end",
	//	},
	//}
	//maze3 := map[string]interface{}{
	//	"forward": "exit",
	//}
	//maze4 := map[string]interface{}{
	//	"forward": "tiger",
	//	"left":    "ogre",
	//	"right":   "demon",
	//}
	//_ = findMaze(maze1, []string{})
	//fmt.Println("The final answer for maze1 is :- \n ", AllAnswers)
	//AllAnswers = [][]string{}
	//_ = findMaze(maze2, []string{})
	//fmt.Println("The final answer for maze2 is:- \n ", AllAnswers)
	//AllAnswers = [][]string{}
	//_ = findMaze(maze3, []string{})
	//fmt.Println("The final answer for maze3 is:- \n ", AllAnswers)
	//AllAnswers = [][]string{}
	//_ = findMaze(maze4, []string{})
	//fmt.Println("The final answer for maze4 is:- \n ", AllAnswers)
}
