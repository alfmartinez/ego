package physics

import "image"

type Body interface {
	IsMobile() bool
	IsHit(image.Rectangle) bool
}

func CreateBody(mobile bool, hitbox image.Rectangle) Body {
	return &body{mobile, hitbox}
}

type body struct {
	mobile bool
	hitbox image.Rectangle
}

func (b *body) IsMobile() bool {
	return b.mobile
}

func (b *body) IsHit(o image.Rectangle) bool {
	return b.hitbox.Intersect(o) != image.Rect(0, 0, 0, 0)
}
