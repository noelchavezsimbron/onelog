package benchmarks

import (
	"github.com/noelchavezsimbron/powerlog"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	logger := powerlog.New(os.Stdout, powerlog.ALL).
		Hook(func(e powerlog.Entry) {
			e.Int64("time", time.Now().Unix())
		})
	logger.InfoWith().
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		String("test", "test").
		Write()
}

func BenchmarkOnelog(b *testing.B) {
	b.Run("with-fields", func(b *testing.B) {
		logger := powerlog.New(ioutil.Discard, powerlog.ALL).
			Hook(func(e powerlog.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		s := struct {
			i int
		}{i: 0}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoWithFields("message", func(e powerlog.Entry) {
					e.Int("test", s.i)
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
					e.String("test", "test")
				})
			}
		})
	})
	b.Run("message-only", func(b *testing.B) {
		logger := powerlog.New(ioutil.Discard, powerlog.ALL).
			Hook(func(e powerlog.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("message")
			}
		})
	})
	b.Run("entry-message-only", func(b *testing.B) {
		logger := powerlog.New(ioutil.Discard, powerlog.ALL).
			Hook(func(e powerlog.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoWith().Write()
			}
		})
	})
	b.Run("entry-fields", func(b *testing.B) {
		logger := powerlog.New(ioutil.Discard, powerlog.ALL).
			Hook(func(e powerlog.Entry) {
				e.Int64("time", time.Now().Unix())
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoWith().
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					String("test", "test").
					Write()
			}
		})
	})

	b.Run("accumulated context", func(b *testing.B) {
		logger := powerlog.New(ioutil.Discard, powerlog.ALL).
			With(func(e powerlog.Entry) {
				e.Int("int", 1)
			})
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("message")
			}
		})
	})
}
