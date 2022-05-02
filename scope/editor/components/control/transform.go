package control

import (
	"math"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/scope/editor/components/gui"
)

const surroundingBoxSize = 3
const arrowLength = 40
const circleRadius = 9
const axisColorX = galx.ColorRed
const axisColorY = galx.ColorGreen
const axisAngleX = galx.Angle0
const axisAngleY = galx.Angle90

type Transform struct {
	selectedObjects []galx.GameObject

	// options
	settings    *gui.Settings
	snapSize    int32
	snapOn      bool
	snapForceOn bool

	// ctl
	axisX  galx.Circle
	axisY  galx.Circle
	axisXY galx.Circle

	// state
	activeX              bool
	activeY              bool
	startMousePosition   galx.Vec2d
	startObjectsPosition []galx.Vec2d
	keyXDown             bool
	keyYDown             bool
	keyCtrlDown          bool

	// visual
	camera            galx.Camera
	attachAnchor      galx.Vec2d
	surroundingSelect galx.Rect
}

func (c Transform) Id() string {
	return "37901d67-ffd9-4ee3-af0a-e98a5f6230cf"
}

func (c *Transform) OnCreated(require galx.EditorComponentResolver) {
	c.settings = require(c.settings).(*gui.Settings)
	c.snapSize = 32
}

func (c *Transform) OnUpdate(state galx.State) error {
	c.displaySettingsWindow()
	c.camera = state.Camera()
	c.keyXDown = state.Keyboard().IsDown('x')
	c.keyYDown = state.Keyboard().IsDown('y')
	c.snapForceOn = state.Movement().Shift()

	if !state.Mouse().IsButtonsAvailable(galx.MousePropagationPriorityEditorGizmos) {
		return nil
	}

	// disable move mode, when mouse released
	if state.Mouse().LeftReleased() && (c.activeX || c.activeY) {
		state.Mouse().StopPropagation(galx.MousePropagationPriorityEditorGizmos)
		c.resetState()
	}

	mousePos := state.Mouse().MouseCoords()
	worldMouse := state.Camera().Screen2World(mousePos)

	// if move mode active
	// disable mouse actions
	// and move objects
	if (c.activeX || c.activeY) && state.Mouse().LeftDown() {
		state.Mouse().StopPropagation(galx.MousePropagationPriorityEditorGizmos)
		c.move(worldMouse)
		return nil
	}

	// find selected objects for draw gizmos
	c.selectedObjects = state.ObjectQueries().AllIn(galx.QueryFlagOnlySelected)
	if len(c.selectedObjects) == 0 {
		return nil
	}

	// calculate anchor/axis
	c.calculateAnchor(worldMouse)

	// on press tick, update widget state
	// and save it on local memory
	// each next frame, we will work
	// with saved state, until mouse is released
	if state.Mouse().LeftPressed() {
		c.saveState(worldMouse)
	}

	return nil
}

func (c *Transform) saveState(worldMouse galx.Vec2d) {
	if c.axisXY.Contains(worldMouse) {
		c.activeX = true
		c.activeY = true
	} else {
		c.activeX = c.axisX.Contains(worldMouse)
		c.activeY = c.axisY.Contains(worldMouse)
	}

	if c.activeX || c.activeY {
		c.startMousePosition = worldMouse
		c.startObjectsPosition = make([]galx.Vec2d, len(c.selectedObjects))
		for idx, object := range c.selectedObjects {
			c.startObjectsPosition[idx] = object.AbsPosition()
		}
	}
}

func (c *Transform) resetState() {
	c.activeX = false
	c.activeY = false
	c.startMousePosition = galx.Vec2d{}
	c.startObjectsPosition = nil
}

func (c *Transform) move(worldMouse galx.Vec2d) {
	lockedX := !c.activeX
	lockedY := !c.activeY

	if c.keyXDown {
		lockedY = true
	}
	if c.keyYDown {
		lockedX = true
	}

	if lockedX && lockedY {
		return
	}

	diff := worldMouse.Sub(c.startMousePosition)

	for idx, object := range c.selectedObjects {
		newPosition := c.startObjectsPosition[idx].Add(diff)

		if lockedX {
			newPosition.X = c.startObjectsPosition[idx].X
		}
		if lockedY {
			newPosition.Y = c.startObjectsPosition[idx].Y
		}

		object.SetPosition(c.snapPosition(newPosition))
	}
}

func (c *Transform) snapPosition(pos galx.Vec2d) galx.Vec2d {
	if !c.snapOn && !c.snapForceOn {
		// both off
		return pos
	}

	if c.snapOn && c.snapForceOn {
		// snap on, but forced to off
		return pos
	}

	return galx.Vec2d{
		X: math.Floor(pos.X/float64(c.snapSize)) * float64(c.snapSize),
		Y: math.Floor(pos.Y/float64(c.snapSize)) * float64(c.snapSize),
	}
}

func (c *Transform) calculateAnchor(worldMouse galx.Vec2d) {
	bbox := make([]galx.Rect, 0, len(c.selectedObjects))
	closest := float64(math.MaxInt16)

	for _, object := range c.selectedObjects {
		objectBox := object.BoundingBox(0)
		bbox = append(bbox, objectBox)

		if distance := worldMouse.DistanceTo(objectBox.Center()); distance < closest {
			// attach controls to the closest element (better UX)
			c.attachAnchor = objectBox.Center()
			closest = distance
		}
	}

	c.surroundingSelect = galx.SurroundRect(bbox...)
	c.axisX = c.axisOn(arrowLength, circleRadius, axisAngleX)
	c.axisY = c.axisOn(arrowLength, circleRadius, axisAngleY)
	c.axisXY = c.axisOn(0, arrowLength-circleRadius, axisAngleY-axisAngleX)
}

func (c *Transform) displaySettingsWindow() {
	c.settings.Extend("Transform", 100, func() {
		// snap to grid
		imgui.Checkbox("Snap to grid", &c.snapOn)

		// snap size
		imgui.InputInt("Snap Size", &c.snapSize)
		if c.snapSize < 1 {
			c.snapSize = 1
		}
		if c.snapSize > 1024 {
			c.snapSize = 1024
		}
	})
}

func (c *Transform) axisOn(len float64, size float64, towards galx.Angle) galx.Circle {
	return galx.Circle{
		Pos:    c.attachAnchor.PolarOffset(len, towards),
		Radius: size,
	}
}

func (c *Transform) OnAfterDraw(r galx.Renderer) error {
	if len(c.selectedObjects) == 0 {
		return nil
	}

	// todo: rotation square with polar coordinate line towards center

	if c.activeX || c.activeY {
		c.drawLocks(r)

		// in move mode, disable controls
		return nil
	}

	// axis
	c.drawAxis(r, c.axisX, axisColorX, axisAngleX, c.activeX)
	c.drawAxis(r, c.axisY, axisColorY, axisAngleY, c.activeY)

	// surrounding bbox
	r.DrawSquare(galx.ColorSelection, c.surroundingSelect.Increase(surroundingBoxSize))

	return nil
}

func (c *Transform) drawLocks(r galx.Renderer) {
	lockedX := !c.activeX
	lockedY := !c.activeY

	if c.keyXDown {
		lockedY = true
	}
	if c.keyYDown {
		lockedX = true
	}

	if lockedY {
		r.DrawLine(axisColorX, galx.Line{
			A: galx.Vec2d{
				X: c.camera.Position().X,
				Y: c.attachAnchor.Y,
			},
			B: galx.Vec2d{
				X: c.camera.Position().X + float64(c.camera.Width()),
				Y: c.attachAnchor.Y,
			},
		})
	}

	if lockedX {
		r.DrawLine(axisColorY, galx.Line{
			A: galx.Vec2d{
				X: c.attachAnchor.X,
				Y: c.camera.Position().Y,
			},
			B: galx.Vec2d{
				X: c.attachAnchor.X,
				Y: c.camera.Position().Y + float64(c.camera.Height()),
			},
		})
	}
}

func (c *Transform) drawAxis(r galx.Renderer, axis galx.Circle, color galx.Color, angle galx.Angle, isActive bool) {
	// draw control circle
	r.DrawCircle(color, axis)
	if isActive {
		r.DrawCircle(color, axis.IncreaseRadius(1))
	}

	// draw line
	r.DrawLine(color, galx.Line{
		A: c.attachAnchor,
		B: c.attachAnchor.PolarOffset(arrowLength-circleRadius, angle),
	})
}
