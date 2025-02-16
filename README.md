Restaurant Order Processing - Lab Task

Overview

This is a lab task to explore different ways of handling restaurant orders using programming concepts in Go. The task involves implementing two approaches:

Sequential Processing - Orders are handled one by one.

Concurrent Processing - Orders are processed in parallel using goroutines and channels.

Objective

The goal of this lab is to compare the efficiency of these two methods and understand how concurrency improves processing time. By using Go's goroutines and channels, we can optimize the system to handle multiple orders simultaneously.

Files

main.go - Contains the implementation of both sequential and concurrent order processing.

Report.md - A detailed explanation of the observations, efficiency differences, and conclusions drawn from the experiment.

How to Run

To execute the program, use the following command:

``` go run main.go ```

This will first run the sequential method and then the concurrent method, displaying the time taken for each.

Expected Outcome

Sequential processing will take longer as each order is handled one at a time.

Concurrent processing will be significantly faster since multiple orders are processed simultaneously.

Refer to [Task Report](Report.md)
 for a more detailed analysis of the results.
