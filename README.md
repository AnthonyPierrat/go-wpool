# WPOOL - A simple WorkerPool 👀
###### Worker Pool implementation in Go for learning purposes based on http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

New in the golang world, I've made this package to train go-routines, channels and other stuff. It has no aim for production use.

## How it works 🛠
> specifify number of workers to the worker pool

```
wp := NewWorkerPool(runtime.NumCPU())
	wp.Start()
	defer wp.Stop()

	for i := 0; i < b.N; i++ {
		wp.Submit(func() {
			Fib(10)
		})
	}
```

## Todos 👨🏻‍💻
- Handle jobs error
- Add better logs
- Send back jobs result if needed
- Add capacity to the pool
- Add CI
