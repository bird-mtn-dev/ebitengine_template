package scenes

import (
	"image/color"
	"log/slog"
	"os"

	"github.com/bird-mtn-dev/ebitengine_template/gamestate"
	fontmanager "github.com/bird-mtn-dev/resource-manager/font-manager"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type TitleScene struct {
	BaseScene
	ui *ebitenui.UI
}

func (s *TitleScene) Update() error {
	_ = s.BaseScene.Update()
	s.ui.Update()
	if gamestate.STATE.InputHandler.ActionIsJustPressed(gamestate.ActionNext) {
		gamestate.STATE.TransitionScene(gamestate.SceneGame)
	}
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	s.ui.Draw(screen)
}

func (s *TitleScene) OnStart() {
	slog.Info("Setting up UI")
	if s.ui == nil {
		titleFace, _ := gamestate.STATE.ResourceMgr.Font.GetFace(fontmanager.STANDARD_NORMAL, 40)
		buttonFace, _ := gamestate.STATE.ResourceMgr.Font.GetFace(fontmanager.STANDARD_NORMAL, 16)
		buttonImage, _ := loadButtonImage()

		rootContainer := widget.NewContainer(
			// the container will use an anchor layout to layout its single child widget
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		)
		// construct a new container that serves as the root of the UI hierarchy
		centeredContainer := widget.NewContainer(
			widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			})),
			// the container will use an anchor layout to layout its single child widget
			widget.ContainerOpts.Layout(widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(20)),
				widget.RowLayoutOpts.Spacing(20),
			)),
		)
		label1 := widget.NewText(
			widget.TextOpts.Text(gamestate.STATE.T("game_title"), titleFace, color.White),
			widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
			),
		)
		centeredContainer.AddChild(label1)

		startbutton := widget.NewButton(
			// set general widget options
			widget.ButtonOpts.WidgetOpts(
				// instruct the container's anchor layout to center the button both horizontally and vertically
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
				widget.WidgetOpts.MinSize(200, 0),
			),

			// specify the images to use
			widget.ButtonOpts.Image(buttonImage),

			// specify the button's text, the font face, and the color
			//widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
			widget.ButtonOpts.Text(gamestate.STATE.T("start_game"), buttonFace, &widget.ButtonTextColor{
				Idle:  color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
				Hover: color.NRGBA{0, 255, 128, 255},
			}),
			// specify that the button's text needs some padding for correct display
			widget.ButtonOpts.TextPadding(widget.Insets{
				Left:   30,
				Right:  30,
				Top:    5,
				Bottom: 5,
			}),

			// add a handler that reacts to clicking the button
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				gamestate.STATE.TransitionScene(gamestate.SceneGame)
			}),
		)

		centeredContainer.AddChild(startbutton)

		exitbutton := widget.NewButton(
			// set general widget options
			widget.ButtonOpts.WidgetOpts(
				// instruct the container's anchor layout to center the button both horizontally and vertically
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				}),
				widget.WidgetOpts.MinSize(200, 0),
			),

			// specify the images to use
			widget.ButtonOpts.Image(buttonImage),

			// specify the button's text, the font face, and the color
			//widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
			widget.ButtonOpts.Text(gamestate.STATE.T("exit_game"), buttonFace, &widget.ButtonTextColor{
				Idle:  color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
				Hover: color.NRGBA{0, 255, 128, 255},
			}),
			// specify that the button's text needs some padding for correct display
			widget.ButtonOpts.TextPadding(widget.Insets{
				Left:   30,
				Right:  30,
				Top:    5,
				Bottom: 5,
			}),

			// add a handler that reacts to clicking the button
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				os.Exit(0)
			}),
		)

		centeredContainer.AddChild(exitbutton)
		rootContainer.AddChild(centeredContainer)
		s.ui = &ebitenui.UI{
			Container: rootContainer,
		}

	}
}
func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewBorderedNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255}, color.NRGBA{90, 90, 90, 255}, 3)

	hover := image.NewBorderedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, color.NRGBA{70, 70, 70, 255}, 3)

	pressed := image.NewAdvancedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, image.NewBorder(3, 2, 2, 2, color.NRGBA{70, 70, 70, 255}))

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
