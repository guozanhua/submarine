package context

import (
	"app/typhenapi/type/submarine/battle"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContainerTest(t *testing.T) {
	Convey("Container", t, func() {
		battleContext := NewContext()
		container := battleContext.Container

		Convey("when an actor is created", func() {
			actor := newSubmarine(battleContext)

			Convey("should add the actor", func() {
				So(container.HasActor(actor.ID()), ShouldBeTrue)
			})

			Convey("should call the actor's Start method", func() {
				So(actor.isCalledStart, ShouldBeTrue)
			})
		})

		Convey("when an actor is destroyed", func() {
			actor := newSubmarine(battleContext)
			actor.Destroy()

			Convey("should remove the actor", func() {
				So(container.HasActor(actor.ID()), ShouldBeFalse)
			})

			Convey("should call the actor's OnDestroy method", func() {
				So(actor.isCalledOnDestroy, ShouldBeTrue)
			})
		})

		Convey("#UpdateActors", func() {
			actors := make([]*actor, 3)
			for i := 0; i < 3; i++ {
				actors[i] = newSubmarine(battleContext)
			}
			container.UpdateActors()

			Convey("should update all actors", func() {
				for _, actor := range actors {
					So(actor.isCalledUpdate, ShouldBeTrue)
				}
			})
		})

		Convey("#Actor", func() {
			actorID := newSubmarine(battleContext).ID()

			Convey("with valid actor id", func() {
				Convey("should return the actor", func() {
					actor := container.Actor(actorID)
					So(actor.ID(), ShouldEqual, actorID)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					actor := container.Actor(actorID + 1)
					So(actor, ShouldBeNil)
				})
			})
		})

		Convey("#SubmarineByUserID", func() {
			userID := newSubmarine(battleContext).UserID()

			Convey("with valid user id", func() {
				Convey("should return the user's submarine", func() {
					submarine := container.SubmarineByUserID(userID)
					So(submarine.UserID(), ShouldEqual, userID)
					So(submarine.ActorType(), ShouldEqual, battle.ActorType_Submarine)
				})
			})

			Convey("with invalid user id", func() {
				Convey("should return nil", func() {
					submarine := container.SubmarineByUserID(userID + 1)
					So(submarine, ShouldBeNil)
				})
			})
		})
	})
}
