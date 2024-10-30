package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	if in == nil {
		return nil
	}

	out := in
	for _, stage := range stages {
		out = executeStage(stage(out), done)
	}

	return out
}

func executeStage(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out)
			for len(in) > 0 {
				<-in
			}
		}()

		for data := range in {
			if isTerminated(done) {
				return
			}

			out <- data
		}
	}()

	return out
}

func isTerminated(done In) bool {
	if done != nil {
		if _, ok := <-done; !ok {
			return true
		}
	}

	return false
}
