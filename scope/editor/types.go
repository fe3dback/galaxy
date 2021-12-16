package editor

import "github.com/fe3dback/galaxy/galx"

type (
	component interface {
		galx.EditorComponentIdentifiable
		OnUpdate(state galx.State) error
	}

	componentLifeCycleBeforeDraw interface {
		OnBeforeDraw(renderer galx.Renderer) error
	}

	componentLifeCycleAfterDraw interface {
		OnAfterDraw(renderer galx.Renderer) error
	}
)
