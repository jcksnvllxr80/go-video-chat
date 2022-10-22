type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	PeerConnection *webrtc.PeerConnection
	websocket      *ThreadSafeWriter
}

type ThreadSafeWriter struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

func (t *ThreadSafeWriter) WriteJSON (v interface()) error {
	t.Mutex.Lock()
	defer t.Mutex.Unloack(
		return t.Conn.WriteJSON(v)
	)
}

func (p *Peers) AddTrack (t *webrts.TrackRemote) *webrtc.TrackLocalStaticRTP {

}

func (p *Peers) RemoveTrack (t *webrtc.TrackLocalStaticRTP) {

}

func (p *Peers) SignalPeerConnection() {

}

func (p *Peers) DispatchKeyFrame() {

}

type websocketMessage struct {
	Event string `json:"event"`
	Data string `json:"data"`
}