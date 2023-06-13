package g2d

type Actor interface {
    Move(arena *Arena)
    //Collide(other Actor)
    Pos() Point
    Size() Point
    Sprite() Point
}

type Arena struct {
    w, h   int
    actors []Actor
    currKeys, prevKeys map[string]bool
    collisions [][]Actor
    turn   int
}

func NewArena(size Point) *Arena {
    a := &Arena{size.X, size.Y, []Actor{}, make(map[string]bool), make(map[string]bool), [][]Actor{}, -1}
    return a
}

func (a *Arena) Spawn(actor Actor) {
    if !a.Contains(actor) {
        a.actors = append(a.actors, actor)
    }
}

func (a *Arena) Kill(actor Actor) {
    for i, v := range a.actors {
        if v == actor {
            a.actors = append(a.actors[:i], a.actors[i+1:]...)
            return
        }
    }
}

func (a *Arena) Contains(actor Actor) bool {
    for _, v := range a.actors {
        if v == actor {
            return true
        }
    }
    return false
}

func (a *Arena) Collisions() []Actor {
  if a.turn < 0 || a.turn >= len(a.collisions) {
    return []Actor{}
  }
    return a.collisions[a.turn]
}

func (a *Arena) Tick(currKeys, prevKeys map[string]bool) {
    a.currKeys = currKeys
    a.prevKeys = prevKeys
    a.collisions = make([][]Actor, len(a.actors))
    actors := a.ReversedActors()
    for i, v := range actors {
        a.collisions[i] = []Actor{}
        for _, w := range actors {
          if v != w && CheckCollision(v, w) {
            a.collisions[i] = append(a.collisions[i], w)
          }
        }
    }

    for t, actor := range actors {
        a.turn = t
        actor.Move(a)
    }
}

func CheckCollision(a1, a2 Actor) bool {
    p1 := a1.Pos()
    s1 := a1.Size()
    p2 := a2.Pos()
    s2 := a2.Size()
    return (p2.X <= p1.X+s1.X && p1.X <= p2.X+s2.X &&
        p2.Y <= p1.Y+s1.Y && p1.Y <= p2.Y+s2.Y)
}

func (a *Arena) Actors() []Actor {
    actors := make([]Actor, len(a.actors))
    for i, v := range a.actors {
        actors[i] = v
    }
    return actors
}

func (a *Arena) ReversedActors() []Actor {
    actors := make([]Actor, len(a.actors))
    for i, v := range a.actors {
        actors[len(a.actors)-i-1] = v
    }
    return actors
}

func (a *Arena) Size() Point {
    return Point{a.w, a.h}
}

func (a *Arena) CurrentKeys() map[string]bool {
    return a.currKeys
}

func (a *Arena) PreviousKeys() map[string]bool {
    return a.prevKeys
}
