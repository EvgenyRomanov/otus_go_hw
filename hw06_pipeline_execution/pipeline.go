package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	for _, stage := range stages {
		tmpChan := make(Bi)

		go func(in, done In) {
			defer close(tmpChan)

			for {
				select {
				case data, ok := <-in:
					if !ok {
						return
					}
					tmpChan <- data
				case <-done:
					return
				}
			}
		}(in, done)

		in = stage(tmpChan)
	}

	return in
}
