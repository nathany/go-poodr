## go-poodr

A [Go](http://golang.org/) translation of the [example code](https://github.com/skmetz/poodr) from [Practical Object-Oriented Design in Ruby](http://www.poodr.info/) by [Sandi Metz](http://sandimetz.com/).

Use the [Go Playground](http://play.golang.org/) or `go run` to try these examples, eg:

    chapter2> go run gear1.go

### 1. Object-Oriented Design

(intro)

### 2. Designing Classes with a Single Responsibility

* `gear1.go` defines a basic Gear with getters
* `gear2.go` introduces a new feature (and responsibility)
* `gear3.go` hide instance variables (behavior in one place)
* `obscure.go` depending on complicated data structures is bad
* `revealing.go` isolating the incoming array
* `gear4.go` extracting wheel as an internal structure
* `gear5.go` a real Wheel with dependency injection

### 3. Managing Dependencies

* `1-dependencies.go` Gear knows too much about Wheel (actually a step back from gear5.go)
* `2-duck-type.go` We don't need a Wheel specifically, just an object that responds to Diameter()
* `3-isolate-new.go` Isolate instance creation (if you can't inject the dependency for some reason)
* `4-isolate-messages.go` Isolate external messages that could be vulnerable to change
* `5-map-init.go` Remove argument order dependencies (probably not the best way to accomplish this)
* Skipped a factory method to work with an unwieldy constructor (gear-wrapper).
* `7-reverse-dependencies.go` What if Wheel depends on Gear? (which is more stable?)

### 4. Creating Flexible Interfaces

(It's all UML! :-)

### 5. Reducing Costs With Duck Typing

(structural typing in Go)

* `trip1.go` A Trip that knows it needs the bicycles prepared.
* `trip2.go` Trip preparation becomes more complicated. It knows too much.
* `trip3.go` A Preparer interface, more abstract but easier to extend.

### 6. Acquiring Behavior Through Inheritance

(which Go doesn't have)

* `bikes1.go` Starting with a road bike.
* `bikes2.go` We need mountain bikes too. Switching on the type.
* Skipped misapplying inheritance.
* `bikes4.go` Implicit delegation and type embedding instead of subclasses.
* `bikes5.go` Specializing the Spares method.

There really is no way to do the template method pattern here.

### 7. Sharing Role Behavior With Modules

### 8. Combining Objects With Composition

### 9. Designing Cost-Effective Tests









