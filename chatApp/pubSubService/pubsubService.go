/*
* @Author: Joakim
* @Year: 2022
* @Description:
* File is used to create pubsub interactions by using structs for messages and
* functions to handle reciving messages and keeping track of when to stop.
 */

package pubSubService

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
)

// Struct used to keep track of messages and who sends them.
type CMessage struct {
	Message  string
	SenderID peer.ID
}

// Struct used to keep track of messages rescived for a topic and subscription
type CRoom struct {
	// Channle to rescive messages
	Messages      chan *CMessage
	CtxBackground context.Context
	// Users CID
	Self peer.ID

	Topic *pubsub.Topic
	Sub   *pubsub.Subscription

	UserName string
}

/*
* Description: This function creates a topic and a subscription for the pubsub message chat.
* Then it initalizes the CRoom struct which is used to interact with the chat.
* Then before returning it starts a threads with a messageRoom that is used to rescive messages.
* @dev Return a pointer to a CRoom struct and nil if sucessfull or error message on failure.
* @return value of 'CRoom struct and error'
 */
func JoinChatRoom(ctxBackground context.Context, pubService *pubsub.PubSub, selfID peer.ID, username string, roomName string) (*CRoom, error) {

	// join the pubsub topic
	topic, _ := pubService.Join(roomName)
	// and subscribe to it
	sub, _ := topic.Subscribe()

	cRoom := &CRoom{
		CtxBackground: ctxBackground,
		Topic:         topic,
		Sub:           sub,
		Self:          selfID,
		UserName:      username,
		Messages:      make(chan *CMessage),
	}

	// start reading messages from the subscription in a loop
	go cRoom.messageLoop()
	return cRoom, nil
}

/*
* Description: This function readLoop pulls messages from the pubsub topic
* and pushes them onto the Messages channel.
* @dev Return Nothing.
* @return value of 'Nothing'
 */
// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (cRoom *CRoom) messageLoop() {
	for {
		msg, err := cRoom.Sub.Next(cRoom.CtxBackground)
		if err != nil {
			close(cRoom.Messages)
			return
		}

		// Read messages thats not sent from self
		if msg.ReceivedFrom != cRoom.Self {
			cRoom.Messages <- &CMessage{
				Message:  string(msg.Data),
				SenderID: msg.ReceivedFrom,
			}
		}
	}
}
