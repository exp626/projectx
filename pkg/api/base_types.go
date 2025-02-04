// GENERATED CODE
// DO NOT EDIT

package game_protocol

type vector struct {
	x int32
	y int32
}

type position struct {
	x int32
	y int32
}

type entity_code byte

const (
	entity_codeplayer entity_code = 0
	entity_codeenemy  entity_code = 1
)
