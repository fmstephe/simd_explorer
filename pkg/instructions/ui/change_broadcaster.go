package ui

import "github.com/google/uuid"

type dataReceiver interface {
	receiverId() uuid.UUID
	dstDataChanged([]byte)
}

type changeBroadcaster struct {
	active    bool
	receivers []dataReceiver
}

func newChangeBroadcaster() *changeBroadcaster {
	return &changeBroadcaster{
		active:    true,
		receivers: []dataReceiver{},
	}
}

func (b *changeBroadcaster) addReceiver(l dataReceiver) {
	b.receivers = append(b.receivers, l)
}

func (b *changeBroadcaster) broadcastChange(bytes []byte, senderId uuid.UUID) {
	if !b.active {
		return
	}

	for _, r := range b.receivers {
		// Avoid updating the receiver that generated this change
		// Updating the sender produces an infinite loop of update events
		r.dstDataChanged(append(make([]byte, 0, len(bytes)), bytes...))
	}
}

func (b *changeBroadcaster) activate() {
	b.active = true
}

func (b *changeBroadcaster) deactivate() {
	b.active = false
}
