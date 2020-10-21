package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	// In is the syntactic sugar for readonly chan.
	In = <-chan interface{}
	// Out is the syntactic sugar for writeonly chan.
	Out = In
	// Bi is the syntactic sugar for chan.
	Bi = chan interface{}
)

// Stage is the syntactic sugar for func, representing pipeline stage.
type Stage func(in In) (out Out)

// ExecutePipeline is the func for executing stages.
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var out Out = in
	for _, s := range stages {
		tmp := make(Bi)
		go func(in In) {
			defer close(tmp)
			for {
				select {
				case val, ok := <-in:
					if !ok {
						return
					}
					tmp <- val
				case <-done:
					return
				}
			}
		}(out)

		out = s(tmp)
	}

	return out
}
