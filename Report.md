# Restaurant Order Processing 

## Introduction
In this task I implemented two ways of handling restaurant orders: one that does everything step by step (sequential) and another that works on multiple orders at the same time (concurrent). To see the difference between each method and figure which is faster.

## How It Works

### 1. Sequential Processing 
- Orders are taken one by one (Like the normal ones).
- Each order is fully prepared before moving to the next.
- This means waiting for each step to finish before starting a new order.

### 2. Concurrent Processing
- Orders are put into a queue (buffered channel) and handled at the same time (parallel).
- One part of the system prepares food while another delivers it at the same time without needing to finish the first one completely.
- This makes better use of time and allows multiple orders to be processed together.

## comparison
- **Sequential processing** was slow because every order had to be completed before starting the next one. (Example: 10.0172 Sec)
- **Concurrent processing** was much faster because it worked on multiple orders at once. (Example: 6.011 Sec)
- By using goroutines and channels, the system could handle many processes at the same time without needing to wait for other flows to finish.


## Conclusion
- The concurrent method is much better for saving time.
- Using goroutines and channels helps dividing the tasks so they don’t slow each other down.
