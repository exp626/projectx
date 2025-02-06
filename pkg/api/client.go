// GENERATED CODE
// DO NOT EDIT

package game_protocol

import (
	"context"
	"errors"
	"io"
)

type Client struct {
	wr io.Writer
}

// 21
func (c *Client) player_move(ctx context.Context, body player_move) (err error) {
	rawCommandBody, err := Newplayer_moveBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, Sizeplayer_move+1)

	rawBody = append(rawBody, CommandCodeplayer_move)
	rawBody = append(rawBody, rawCommandBody[:]...)

	n, err := c.wr.Write(rawBody)
	if err != nil {
		return err
	}

	if n != len(rawBody) {
		return errors.New("all information was not writen")
	}

	return nil
}

// 19
func (c *Client) create_player(ctx context.Context, body create_player) (err error) {
	rawCommandBody, err := Newcreate_playerBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, Sizecreate_player+1)

	rawBody = append(rawBody, CommandCodecreate_player)
	rawBody = append(rawBody, rawCommandBody[:]...)

	n, err := c.wr.Write(rawBody)
	if err != nil {
		return err
	}

	if n != len(rawBody) {
		return errors.New("all information was not writen")
	}

	return nil
}

// 20
func (c *Client) input(ctx context.Context, body input) (err error) {
	rawCommandBody, err := NewinputBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, Sizeinput+1)

	rawBody = append(rawBody, CommandCodeinput)
	rawBody = append(rawBody, rawCommandBody[:]...)

	n, err := c.wr.Write(rawBody)
	if err != nil {
		return err
	}

	if n != len(rawBody) {
		return errors.New("all information was not writen")
	}

	return nil
}
