// GENERATED BY PROTOCOL_BUILDER
// DO NOT TOUCH

package protocol

import (
	"bytes"
	"io"
)

func (k *KeepAliveClientbound) id() int { return 0 }
func (k *KeepAliveClientbound) write(w *bytes.Buffer) error {
	var err error
	if err = writeVarInt(w, k.ID); err != nil {
		return err
	}
	return nil
}
func (k *KeepAliveClientbound) read(r *bytes.Reader) error {
	var err error
	if k.ID, err = readVarInt(r); err != nil {
		return err
	}
	return nil
}

func (j *JoinGame) id() int { return 1 }
func (j *JoinGame) write(w *bytes.Buffer) error {
	var err error
	var tmp [4]byte
	tmp[0] = byte(j.EntityID >> 24)
	tmp[1] = byte(j.EntityID >> 16)
	tmp[2] = byte(j.EntityID >> 8)
	tmp[3] = byte(j.EntityID >> 0)
	if _, err = w.Write(tmp[:4]); err != nil {
		return err
	}
	tmp[0] = byte(j.Gamemode >> 0)
	if _, err = w.Write(tmp[:1]); err != nil {
		return err
	}
	tmp[0] = byte(j.Dimension >> 0)
	if _, err = w.Write(tmp[:1]); err != nil {
		return err
	}
	tmp[0] = byte(j.Difficulty >> 0)
	if _, err = w.Write(tmp[:1]); err != nil {
		return err
	}
	tmp[0] = byte(j.MaxPlayers >> 0)
	if _, err = w.Write(tmp[:1]); err != nil {
		return err
	}
	if err = writeString(w, j.LevelType); err != nil {
		return err
	}
	if err = writeBool(w, j.ReducedDebugInfo); err != nil {
		return err
	}
	return nil
}
func (j *JoinGame) read(r *bytes.Reader) error {
	var err error
	var tmp [4]byte
	if _, err = io.ReadFull(r, tmp[:4]); err != nil {
		return err
	}
	j.EntityID = int32((uint32(tmp[3]) << 0) | (uint32(tmp[2]) << 8) | (uint32(tmp[1]) << 16) | (uint32(tmp[0]) << 24))
	if _, err = io.ReadFull(r, tmp[:1]); err != nil {
		return err
	}
	j.Gamemode = (byte(tmp[0]) << 0)
	if _, err = io.ReadFull(r, tmp[:1]); err != nil {
		return err
	}
	j.Dimension = int8((uint8(tmp[0]) << 0))
	if _, err = io.ReadFull(r, tmp[:1]); err != nil {
		return err
	}
	j.Difficulty = (byte(tmp[0]) << 0)
	if _, err = io.ReadFull(r, tmp[:1]); err != nil {
		return err
	}
	j.MaxPlayers = (byte(tmp[0]) << 0)
	if j.LevelType, err = readString(r); err != nil {
		return err
	}
	if j.ReducedDebugInfo, err = readBool(r); err != nil {
		return err
	}
	return nil
}

func init() {
	packetCreator[Play][clientbound][0] = func() Packet { return &KeepAliveClientbound{} }
	packetCreator[Play][clientbound][1] = func() Packet { return &JoinGame{} }
}
