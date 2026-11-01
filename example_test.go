package zoltar_test

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/dogmatiq/zoltar"
)

// This example predicts what a function returns when it is called with a
// known argument.
func Example() {
	ctx := context.Background()

	pkg := buildPackage(`
		package greet

		func Greeting(name string) string {
			return "Hello, " + name + "!"
		}
	`)

	z := zoltar.New(pkg.Prog)

	p, err := z.Predict(ctx, pkg.Func("Greeting"), zoltar.Known("Zoltar"))
	if err != nil {
		panic(err)
	}

	results, err := p.Returns(ctx)
	if err != nil {
		panic(err)
	}

	greeting, _ := results[0].As[string]()
	fmt.Println(greeting)

	// Output:
	// Hello, Zoltar!
}

// This example predicts the value passed to a call, just before the call
// happens.
func ExamplePrediction_ValueOf() {
	ctx := context.Background()

	pkg := buildPackage(`
		package server

		func listen(addr string) {}

		func Start(host string) {
			port := "8080"
			listen(host + ":" + port)
		}
	`)

	z := zoltar.New(pkg.Prog)
	start := pkg.Func("Start")

	p, err := z.Predict(ctx, start, zoltar.Known("localhost"))
	if err != nil {
		panic(err)
	}

	call := callTo(start, "listen")

	addr, err := p.ValueOf(ctx, call.Call.Args[0], call)
	if err != nil {
		panic(err)
	}

	s, _ := addr.As[string]()
	fmt.Println(s)

	// Output:
	// localhost:8080
}

// This example shows a value that differs depending on how the program runs.
// The argument isn't known, so the predictor can't tell which way the if
// statement goes, and the result has a timeline for each.
func ExampleValue_Timelines() {
	ctx := context.Background()

	pkg := buildPackage(`
		package greet

		func Greeting(name string) string {
			if name == "" {
				name = "world"
			}
			return "Hello, " + name + "!"
		}
	`)

	z := zoltar.New(pkg.Prog)
	greeting := pkg.Func("Greeting")

	p, err := z.Predict(
		ctx,
		greeting,
		zoltar.Unknown(greeting.Params[0].Type(), "chosen by the caller"),
	)
	if err != nil {
		panic(err)
	}

	results, err := p.Returns(ctx)
	if err != nil {
		panic(err)
	}

	var timelines []string
	for _, t := range results[0].Timelines() {
		if s, ok := t.As[string](); ok {
			timelines = append(timelines, s)
		} else {
			timelines = append(timelines, "(unknown)")
		}
	}

	slices.Sort(timelines)
	fmt.Println(strings.Join(timelines, "\n"))

	// Output:
	// (unknown)
	// Hello, world!
}

// This example shows a blind spot: the predictor doesn't follow calls into
// os.Getenv(), so it can't see what it returns, and the blind spot says why.
func ExampleValue_BlindSpots() {
	ctx := context.Background()

	pkg := buildPackage(`
		package config

		import "os"

		func Port() string {
			return os.Getenv("PORT")
		}
	`)

	z := zoltar.New(pkg.Prog)

	p, err := z.Predict(ctx, pkg.Func("Port"))
	if err != nil {
		panic(err)
	}

	results, err := p.Returns(ctx)
	if err != nil {
		panic(err)
	}

	for _, b := range results[0].BlindSpots() {
		fmt.Println(b.Kind())
	}

	// Output:
	// NotFollowed
}

// This example follows a call into a helper function, to find the name of
// the environment variable that a package reads.
func ExampleTarget_Follow() {
	ctx := context.Background()

	pkg := buildPackage(`
		package config

		import "os"

		func Addr() string {
			return env("ADDR")
		}

		func env(name string) string {
			return os.Getenv("MYAPP_" + name)
		}
	`)

	z := zoltar.New(pkg.Prog)
	addr := pkg.Func("Addr")

	p, err := z.Predict(ctx, addr)
	if err != nil {
		panic(err)
	}

	// Find what the call to env() runs, and continue inside it.
	targets, err := p.Targets(ctx, callTo(addr, "env"))
	if err != nil {
		panic(err)
	}

	env, err := targets[0].Follow(ctx)
	if err != nil {
		panic(err)
	}

	call := callTo(env.Function(), "Getenv")

	name, err := env.ValueOf(ctx, call.Call.Args[0], call)
	if err != nil {
		panic(err)
	}

	s, _ := name.As[string]()
	fmt.Println(s)

	// Output:
	// MYAPP_ADDR
}

// This example tells the predictor how to compute calls to a function in the
// standard library, which it doesn't follow.
func ExampleReturnsResultOf() {
	ctx := context.Background()

	pkg := buildPackage(`
		package naming

		import "strings"

		func Name() string {
			return strings.ToUpper("hello")
		}
	`)

	z := zoltar.New(
		pkg.Prog,
		zoltar.WithAssumptions(
			"strings", "ToUpper",
			zoltar.ReturnsResultOf(strings.ToUpper),
			zoltar.Pure,
		),
	)

	p, err := z.Predict(ctx, pkg.Func("Name"))
	if err != nil {
		panic(err)
	}

	results, err := p.Returns(ctx)
	if err != nil {
		panic(err)
	}

	name, _ := results[0].As[string]()
	fmt.Println(name)

	// Output:
	// HELLO
}
