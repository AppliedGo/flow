/*
<!--
Copyright (c) 2017 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = ""
description = ""
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2017-02-24"
publishdate = "2017-02-24"
draft = "true"
domains = ["Concurrent Programming"]
tags = ["fbp", "flow-based programming", "dataflow"]
categories = ["Tutorial"]
+++

### Summary goes here

<!--more-->

## Intro goes here

## The code
*/

// ## Imports and globals

package main

import (
	"fmt"
	"strings"

	"github.com/trustmaster/goflow"
)

// A component that generates greetings
type Counter struct {
	flow.Component               // component "superclass" embedded
	Sentence       <-chan string // input port
	Count          chan<- int    // output port
}

// Reaction to a new sentence input
func (c *Counter) OnSentence(sentence string) {
	c.Count <- len(strings.Split(sentence, " "))
}

// A component that prints its input on screen
type Printer struct {
	flow.Component
	Line <-chan int // inport
}

// Prints a line when it gets it
func (p *Printer) OnCount(count int) {
	fmt.Printf("%d\n", count)
}

// Our counter network
type CounterApp struct {
	flow.Graph // graph "superclass" embedded
}

// Graph constructor and structure definition
func NewCounterApp() *CounterApp {
	n := new(CounterApp) // creates the object in heap
	n.InitGraphState()   // allocates memory for the graph
	// Add processes to the network
	n.Add(new(Counter), "counter")
	n.Add(new(Printer), "printer")
	// Connect them with a channel
	n.Connect("counter", "Count", "printer", "Line")
	// Our net has 1 inport mapped to counter.Sentence
	n.MapInPort("In", "Counter", "Sentence")
	return n
}

func main() {
	// Create the network
	net := NewGreetingApp()
	// We need a channel to talk to it
	in := make(chan string)
	net.SetInPort("In", in)
	// Run the net
	flow.RunNet(net)
	// Now we can send some names and see what happens
	in <- "John"
	in <- "Boris"
	in <- "Hanna"
	// Close the input to shut the network down
	close(in)
	// Wait until the app has done its job
	<-net.Wait()
}

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/github.com/appliedgo/TODO:

Step 3. Run the binary.

    ./TODO:


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
