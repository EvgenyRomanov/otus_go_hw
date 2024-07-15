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
		in = func(stage Stage, in, done In) In {
			ch := make(chan interface{})

			go func() {
				defer close(ch)

				for data := range stage(in) {
					select {
					case ch <- data:
					case <-done:
						return
					}
				}
			}()

			return ch
		}(stage, in, done)
	}

	return in
}
