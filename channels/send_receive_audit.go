package channels

/*
	Create audit object to track sends vs receives
	Consists of 2 channels:
	- sendQ		: every send, writes message to this Q
	- receiveQ	: every receive, writes message to this Q

	Internally compare sends to receives
*/
