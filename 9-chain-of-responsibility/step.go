package main

type step interface {
	run(*customer)
	setNextStep(step)
}
