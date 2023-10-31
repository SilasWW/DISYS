# TCP/IP Simulator in Go


Implement the TCP/IP protocol in Go. Your implementation has to be a simulation of the protocol seen in class (see slides).

There are different levels that you can work on. In order to pass, you need to implement at least (1) or (2).



(1)[Easy] Implement the TCP/IP Handshake using threads. This is not realistic (since the protocol should run across a network) but your implementation needs to show that you have a good understanding of the protocol. 

(2)[Hard] Implement a TCP/IP Handshake using the net package.

(3)[Medium] Implement a forwarder process/thread that simulates the middleware, where messages can be delayed or lost. All messages must go through the forwarder.	    



Attach to your submission, a *README* file answering the following questions:

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

d) In case messages can be delayed or lost, how does your implementation handle message loss?

e) Why is the 3-way handshake important?
