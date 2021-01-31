// Package powerlog is a fast, low allocation and modular JSON logger.
//
// Basic usage:
// 	import "github.com/noelchavezsimbron/powerlog/log"
//
//	log.Info("hello world !") // {"level":"info","message":"hello world !", "time":1494567715}
//
// You can create your own logger:
//	import "github.com/noelchavezsimbron/powerlog"
//
//	var logger = onelog.New(os.Stdout, onelog.ALL)
//
//	func main() {
//		logger.Info("hello world !") // {"level":"info","message":"hello world !"}
//	}
package powerlog