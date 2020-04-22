package main

import (
	"testing"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func TestFlatMapExample(t *testing.T) {
	// given
	observerMock := observer.NewObserverMock()

	// and
	primeSequence := observable.Just([]int{2, 3, 5, 7, 11, 13})

	// when
	<-primeSequence.
		FlatMap(func(primes interface{}) observable.Observable {
			return observable.Create(func(emitter *observer.Observer, disposed bool) {
				for _, prime := range primes.([]int) {
					emitter.OnNext(prime)
				}
				emitter.OnDone()
			})
		}, 1).
		Last().
		Subscribe(observerMock.Capture())

	// then
	observerMock.AssertCalled(t, "OnNext", 13)
}
