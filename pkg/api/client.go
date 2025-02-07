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

func NewClient(wr io.Writer) *Client {
	return &Client{
		wr: wr,
	}
}

// 21
func (c *Client) PlayerMove(ctx context.Context, body PlayerMoveBody) (err error) {
	rawCommandBody, err := NewPlayerMoveBodyBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, SizePlayerMoveBody+1)

	rawBody = append(rawBody, CommandCodePlayerMove)
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
func (c *Client) CreatePlayer(ctx context.Context, body CreatePlayerBody) (err error) {
	rawCommandBody, err := NewCreatePlayerBodyBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, SizeCreatePlayerBody+1)

	rawBody = append(rawBody, CommandCodeCreatePlayer)
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
func (c *Client) Input(ctx context.Context, body InputBody) (err error) {
	rawCommandBody, err := NewInputBodyBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, SizeInputBody+1)

	rawBody = append(rawBody, CommandCodeInput)
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
