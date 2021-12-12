package node

import (
	"github.com/fe3dback/galaxy/galx"
)

type ObjectQuery struct {
	sceneManager galx.SceneManager
	camera       galx.Camera
}

func NewObjectQuery(sceneManager galx.SceneManager, camera galx.Camera) *ObjectQuery {
	return &ObjectQuery{
		sceneManager: sceneManager,
		camera:       camera,
	}
}

func (o *ObjectQuery) All() []galx.GameObject {
	return o.sceneManager.Current().Entities()
}

func (o *ObjectQuery) AllIn(flag galx.QueryFlag) []galx.GameObject {
	list := make([]galx.GameObject, 0)

	for _, object := range o.sceneManager.Current().Entities() {
		if flag&galx.QueryFlagExcludeLocked != 0 && o.isLocked(object) {
			continue
		}
		if flag&galx.QueryFlagExcludeRoots != 0 && o.isRoot(object) {
			continue
		}
		if flag&galx.QueryFlagExcludeLeaf != 0 && o.isLeaf(object) {
			continue
		}
		if flag&galx.QueryFlagOnlyOnScreen != 0 && o.isOutsideScreen(object) {
			continue
		}
		if flag&galx.QueryFlagOnlySelected != 0 && !o.isSelected(object) {
			continue
		}

		list = append(list, object)
	}

	return list
}

func (o *ObjectQuery) isLocked(object galx.GameObject) bool {
	return object.IsLocked()
}

func (o *ObjectQuery) isRoot(object galx.GameObject) bool {
	return object.IsRoot()
}

func (o *ObjectQuery) isLeaf(object galx.GameObject) bool {
	return object.IsLeaf()
}

func (o *ObjectQuery) isSelected(object galx.GameObject) bool {
	return object.IsSelected()
}

func (o *ObjectQuery) isOutsideScreen(object galx.GameObject) bool {
	pos := object.AbsPosition()
	camPos := o.camera.Position()

	if pos.X < camPos.X {
		return true
	}
	if pos.Y < camPos.Y {
		return true
	}
	if pos.X > camPos.X+float64(o.camera.Width()) {
		return true
	}
	if pos.Y > camPos.Y+float64(o.camera.Height()) {
		return true
	}

	return false
}
