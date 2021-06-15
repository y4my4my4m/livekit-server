package rtc

import (
	"github.com/pion/ion-sfu/pkg/sfu"
	"github.com/pion/webrtc/v3"

	livekit "github.com/livekit/livekit-server/proto"
	"github.com/livekit/protocol/utils"
)

type SubscribedTrack struct {
	sourceSid   string
	dt          *sfu.DownTrack
	transceiver *webrtc.RTPTransceiver
	subMuted    utils.AtomicFlag
	pubMuted    utils.AtomicFlag
}

func NewSubscribedTrack(sourceSid string, dt *sfu.DownTrack, transceiver *webrtc.RTPTransceiver) *SubscribedTrack {
	return &SubscribedTrack{
		sourceSid:   sourceSid,
		dt:          dt,
		transceiver: transceiver,
	}
}

func (t *SubscribedTrack) ID() string {
	return t.dt.ID()
}

func (t *SubscribedTrack) MID() string {
	return t.transceiver.Mid()
}

func (t *SubscribedTrack) DownTrack() *sfu.DownTrack {
	return t.dt
}

// has subscriber indicated it wants to mute this track
func (t *SubscribedTrack) IsMuted() bool {
	return t.subMuted.Get()
}

// set subscriber mute preference
func (t *SubscribedTrack) SetMuted(muted bool) {
	t.subMuted.TrySet(muted)
	t.updateDownTrackMute()
}

func (t *SubscribedTrack) SetPublisherMuted(muted bool) {
	t.pubMuted.TrySet(muted)
	t.updateDownTrackMute()
}

func (t *SubscribedTrack) SetVideoQuality(quality livekit.VideoQuality) {
	if t.dt.Kind() == webrtc.RTPCodecTypeVideo {
		t.dt.SwitchSpatialLayer(int64(quality), true)
	}
}

func (t *SubscribedTrack) ToProto() *livekit.SubscribedTrack {
	return &livekit.SubscribedTrack{
		Mid:            t.MID(),
		TrackSid:       t.dt.ID(),
		ParticipantSid: t.sourceSid,
	}
}

func (t *SubscribedTrack) updateDownTrackMute() {
	muted := t.subMuted.Get() || t.pubMuted.Get()
	t.dt.Mute(muted)
}
