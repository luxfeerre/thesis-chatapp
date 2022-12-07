/*
* The functions in this program are based on
* https://github.com/libp2p/go-libp2p/blob/master/examples/pubsub/chat/ui.go
* They use the code as a skeleton program to fit it to the rest and get a valid UI for this proof of work.
* It writes when changes have been made which are the authors.
* @Author: Joakim
* @Year: 2022
* @Description:
* File is used to create a user interface which fits this program needs
* in both form of usebility, integration with the smart contracts and testing.
 */

package userInterface

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/blockchain/chatApp/ethereumService"
	"github.com/blockchain/chatApp/pubSubService"
	state "github.com/blockchain/chatApp/state"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/rivo/tview"
)

// Used to keep if a thread is changing the hash map of peer ids
var mapCanBeUpdated bool = true

// Modified ChatUserInterface struct which has added a hash map to keep track of
// peer ids mapped to a username.
type ChatUserInterface struct {
	Chat      *pubSubService.CRoom
	App       *tview.Application
	PeersList *tview.TextView

	MsgW    io.Writer
	InputCh chan string
	DoneCh  chan struct{}
	PeerMap map[peer.ID]string
}

// This function is not modified except for adding the PerrMap to keep track of users in the chat.
// NewChatUI returns a new ChatUI struct that controls the text UI.
// It won't actually do anything until you call Run().
func NewChatUI(chat *pubSubService.CRoom) *ChatUserInterface {
	app := tview.NewApplication()

	// make a text view to contain our chat messages
	msgBox := tview.NewTextView()
	msgBox.SetDynamicColors(true)
	msgBox.SetBorder(true)
	msgBox.SetTitle(fmt.Sprintf("Room: %s", chat.Topic.String()))

	// text views are io.Writers, but they don't automatically refresh.
	// this sets a change handler to force the app to redraw when we get
	// new messages to display.
	msgBox.SetChangedFunc(func() {
		app.Draw()
	})

	// an input field for typing messages into
	inputCh := make(chan string, 64)
	input := tview.NewInputField().
		SetLabel(chat.UserName + " > ").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorBlack)

	// the done func is called when the user hits enter, or tabs out of the field
	input.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			// we don't want to do anything if they just tabbed away
			return
		}
		line := input.GetText()
		if len(line) == 0 {
			// ignore blank lines
			return
		}

		// bail if requested
		if line == "/quit" {
			app.Stop()
			return
		}

		// send the line onto the input chan and reset the field text
		inputCh <- line
		input.SetText("")
	})

	// make a text view to hold the list of peers in the room, updated by ui.refreshPeers()
	peersList := tview.NewTextView()
	peersList.SetBorder(true)
	peersList.SetTitle("Peers")
	peersList.SetChangedFunc(func() { app.Draw() })

	// chatPanel is a horizontal box with messages on the left and peers on the right
	// the peers list takes 20 columns, and the messages take the remaining space
	chatPanel := tview.NewFlex().
		AddItem(msgBox, 0, 1, false).
		AddItem(peersList, 20, 1, false)

	// flex is a vertical box with the chatPanel on top and the input field at the bottom.

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(chatPanel, 0, 1, false).
		AddItem(input, 1, 1, true)

	app.SetRoot(flex, true)

	// Added a PeerMap to keep track of users and set it to default size
	return &ChatUserInterface{
		Chat:      chat,
		App:       app,
		PeersList: peersList,
		MsgW:      msgBox,
		InputCh:   inputCh,
		DoneCh:    make(chan struct{}),
		PeerMap:   make(map[peer.ID]string, 128),
	}
}

// This function is not modified other then for the test which changed it to the testEventHandler
// Run starts the chat event loop in the background, then starts
// the event loop for the text UI.
func (ui *ChatUserInterface) Run(instance *state.State, transactionData *bind.TransactOpts) error {
	go ui.testEventHandler(instance, transactionData)
	defer ui.end()
	return ui.App.Run()
}

// This function is not modified.
// end signals the event loop to exit gracefully
func (ui *ChatUserInterface) end() {
	ui.DoneCh <- struct{}{}
}

/*
* Description: This function first clears the UI of peer names and then adds thiere userName
* to the user interface by iterating over the hashmap and then using ForceDraw to show them to the user.
* @dev Return nothing.
* @return value of 'nothing'
 */
func (ui *ChatUserInterface) refreshPeers() {
	ui.PeersList.Clear()

	for _, userName := range ui.PeerMap {
		fmt.Fprintln(ui.PeersList, userName)
	}

	ui.App.ForceDraw()
}

/*
* Description: This function modifies the PerrMap in the ui struct
* to keep track of users allready i nthe hashmap and otherwise query the Ethereum service for thier
* username and mapp it to thier peerid.
* @dev Return nothing.
* @return value of 'nothing'
 */
func (ui *ChatUserInterface) mapPeerToUname(instance *state.State, transactionData *bind.TransactOpts) {

	// Get a slice of active peers in the chat
	peers := ui.Chat.Topic.ListPeers()

	// Iterate over found peers and check if they are in the hash map
	for _, foundPeer := range peers {
		if _, ok := ui.PeerMap[foundPeer]; !ok {
			username, _ := ethereumService.GetUserName(instance, transactionData, foundPeer.Pretty())
			ui.PeerMap[foundPeer] = username
		}
	}
}

// This function is not modified.
// displayChatMessage writes a ChatMessage from the room to the message window,
// with the sender's nick highlighted in green.
func (ui *ChatUserInterface) displayChatMessage(cm *pubSubService.CMessage) {
	prompt := setColor("green", fmt.Sprintf("<%s>:", ui.PeerMap[cm.SenderID]))
	fmt.Fprintf(ui.MsgW, "%s %s\n", prompt, cm.Message)
}

// This function is not modified.
// displaySelfMessage writes a message from ourself to the message window,
// with our nick highlighted in yellow.
func (ui *ChatUserInterface) DisplaySelfMessage(msg string) {
	prompt := setColor("yellow", fmt.Sprintf("<%s>:", ui.Chat.UserName))
	fmt.Fprintf(ui.MsgW, "%s %s\n", prompt, msg)
}

/* This function is modified in the way it publishes messages or updates the UI with new peers.
 * This function runs a loop where it selects input read from channels.
 * Displays received messages from the chat room or sends them.
 * It also refreshes the list of peers in the terminal.
 */
func (ui *ChatUserInterface) eventHandler(instance *state.State, transactionData *bind.TransactOpts) {
	peerRefreshTicker := time.NewTicker(time.Second * 2)
	defer peerRefreshTicker.Stop()

	for {
		select {
		case input := <-ui.InputCh:
			// when the user types in a line, publish it to the chat room and print to the message window
			err := ui.Chat.Topic.Publish(ui.Chat.CtxBackground, []byte(input))
			if err != nil {
				log.Printf("publish error: %s", err)
			}
			ui.DisplaySelfMessage(input)

		case msg := <-ui.Chat.Messages:
			// when we receive a message from the chat room, print it to the message window
			ui.displayChatMessage(msg)

		case <-peerRefreshTicker.C:
			// refresh the list of peers in the chat room periodically
			if mapCanBeUpdated {
				mapCanBeUpdated = false
				ui.mapPeerToUname(instance, transactionData)
				ui.refreshPeers()
				mapCanBeUpdated = true
			}

		case <-ui.DoneCh:
			return

		case <-ui.Chat.CtxBackground.Done():
			return
		}
	}
}

/*
* Description: This function is used to mimick a real user and sending 1 message every second
* of length 64 bytes until 60 has been sent. While also updating the peer ui list and then exit when finshed.
* @dev Return nothing.
* @return value of 'nothing'
 */
func (ui *ChatUserInterface) testEventHandler(instance *state.State, transactionData *bind.TransactOpts) {
	// How often messages gets sent
	messageRefreshTicker := time.NewTicker(time.Second * 1)
	// Message counter
	messageCounter := 0
	// Max message lenght of 64 bytes
	payload := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	peerRefreshTicker := time.NewTicker(time.Second * 2)

	defer peerRefreshTicker.Stop()
	defer messageRefreshTicker.Stop()

	for {
		select {
		case <-messageRefreshTicker.C:
			if messageCounter < 60 {
				err := ui.Chat.Topic.Publish(ui.Chat.CtxBackground, []byte(payload))
				if err != nil {
					log.Printf("publish error: %s", err)
				}
				ui.DisplaySelfMessage(payload)
				messageCounter++
			} else {
				<-ui.DoneCh
				return
			}
		case msg := <-ui.Chat.Messages:
			// when we receive a message from the chat room, print it to the message window
			ui.displayChatMessage(msg)

		case <-peerRefreshTicker.C:
			// refresh the list of peers in the chat room periodically
			if mapCanBeUpdated {
				mapCanBeUpdated = false
				ui.mapPeerToUname(instance, transactionData)
				ui.refreshPeers()
				mapCanBeUpdated = true
			}
		case <-ui.Chat.CtxBackground.Done():
			return
		}
	}
}

// This function sets a color for terminal output.
func setColor(color, msg string) string {
	return fmt.Sprintf("[%s]%s[-]", color, msg)
}
