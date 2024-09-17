a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
    Packages are structs containing two int fields (representing seq and ack). Channels are used to send data between server and client.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
    We use threads. This is not realistic because the threads are running locally instead of over a network.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
    If our messages happend to arrive in a wrong order, we would attempt to store this noncompatible data and wait for the
    next message until we recieve a compatible message, then we can use the stored message checking if any of them are now 
    compatible before recieving the next message.

d) In case messages can be delayed or lost, how does your implementation handle message loss?
    our implementation would wait for an amount of time and if that time has elapsed resend the previous request

e) Why is the 3-way handshake important?
    It assures that a connection has been properly established without any noticable problems.