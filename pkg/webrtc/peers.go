type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}