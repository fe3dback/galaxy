package event

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

//goland:noinspection GoUnusedConst
const (
	KeyboardPressTypeReleased KeyboardPressType = iota
	KeyboardPressTypePressed
	KeyboardPressTypeDoublePressed
)

//goland:noinspection GoUnusedConst
const (
	Key0                  = sdl.K_0
	Key1                  = sdl.K_1
	Key2                  = sdl.K_2
	Key3                  = sdl.K_3
	Key4                  = sdl.K_4
	Key5                  = sdl.K_5
	Key6                  = sdl.K_6
	Key7                  = sdl.K_7
	Key8                  = sdl.K_8
	Key9                  = sdl.K_9
	KeyA                  = sdl.K_a
	KeyAcBack             = sdl.K_AC_BACK
	KeyAcBookmarks        = sdl.K_AC_BOOKMARKS
	KeyAcForward          = sdl.K_AC_FORWARD
	KeyAcHome             = sdl.K_AC_HOME
	KeyAcRefresh          = sdl.K_AC_REFRESH
	KeyAcSearch           = sdl.K_AC_SEARCH
	KeyAcStop             = sdl.K_AC_STOP
	KeyAgain              = sdl.K_AGAIN
	KeyAlterase           = sdl.K_ALTERASE
	KeyQuote              = sdl.K_QUOTE
	KeyApplication        = sdl.K_APPLICATION
	KeyAudiomute          = sdl.K_AUDIOMUTE
	KeyAudionext          = sdl.K_AUDIONEXT
	KeyAudioplay          = sdl.K_AUDIOPLAY
	KeyAudioprev          = sdl.K_AUDIOPREV
	KeyAudiostop          = sdl.K_AUDIOSTOP
	KeyB                  = sdl.K_b
	KeyBackslash          = sdl.K_BACKSLASH
	KeyBackspace          = sdl.K_BACKSPACE
	KeyBrightnessdown     = sdl.K_BRIGHTNESSDOWN
	KeyBrightnessup       = sdl.K_BRIGHTNESSUP
	KeyC                  = sdl.K_c
	KeyCalculator         = sdl.K_CALCULATOR
	KeyCancel             = sdl.K_CANCEL
	KeyCapslock           = sdl.K_CAPSLOCK
	KeyClear              = sdl.K_CLEAR
	KeyClearagain         = sdl.K_CLEARAGAIN
	KeyComma              = sdl.K_COMMA
	KeyComputer           = sdl.K_COMPUTER
	KeyCopy               = sdl.K_COPY
	KeyCrsel              = sdl.K_CRSEL
	KeyCurrencysubunit    = sdl.K_CURRENCYSUBUNIT
	KeyCurrencyunit       = sdl.K_CURRENCYUNIT
	KeyCut                = sdl.K_CUT
	KeyD                  = sdl.K_d
	KeyDecimalseparator   = sdl.K_DECIMALSEPARATOR
	KeyDelete             = sdl.K_DELETE
	KeyDisplayswitch      = sdl.K_DISPLAYSWITCH
	KeyDown               = sdl.K_DOWN
	KeyE                  = sdl.K_e
	KeyEject              = sdl.K_EJECT
	KeyEnd                = sdl.K_END
	KeyEquals             = sdl.K_EQUALS
	KeyEscape             = sdl.K_ESCAPE
	KeyExecute            = sdl.K_EXECUTE
	KeyExsel              = sdl.K_EXSEL
	KeyF                  = sdl.K_f
	KeyF1                 = sdl.K_F1
	KeyF10                = sdl.K_F10
	KeyF11                = sdl.K_F11
	KeyF12                = sdl.K_F12
	KeyF13                = sdl.K_F13
	KeyF14                = sdl.K_F14
	KeyF15                = sdl.K_F15
	KeyF16                = sdl.K_F16
	KeyF17                = sdl.K_F17
	KeyF18                = sdl.K_F18
	KeyF19                = sdl.K_F19
	KeyF2                 = sdl.K_F2
	KeyF20                = sdl.K_F20
	KeyF21                = sdl.K_F21
	KeyF22                = sdl.K_F22
	KeyF23                = sdl.K_F23
	KeyF24                = sdl.K_F24
	KeyF3                 = sdl.K_F3
	KeyF4                 = sdl.K_F4
	KeyF5                 = sdl.K_F5
	KeyF6                 = sdl.K_F6
	KeyF7                 = sdl.K_F7
	KeyF8                 = sdl.K_F8
	KeyF9                 = sdl.K_F9
	KeyFind               = sdl.K_FIND
	KeyG                  = sdl.K_g
	KeyBackquote          = sdl.K_BACKQUOTE
	KeyH                  = sdl.K_h
	KeyHelp               = sdl.K_HELP
	KeyHome               = sdl.K_HOME
	KeyI                  = sdl.K_i
	KeyInsert             = sdl.K_INSERT
	KeyJ                  = sdl.K_j
	KeyK                  = sdl.K_k
	KeyKbdillumdown       = sdl.K_KBDILLUMDOWN
	KeyKbdillumtoggle     = sdl.K_KBDILLUMTOGGLE
	KeyKbdillumup         = sdl.K_KBDILLUMUP
	KeyKp0                = sdl.K_KP_0
	KeyKp00               = sdl.K_KP_00
	KeyKp000              = sdl.K_KP_000
	KeyKp1                = sdl.K_KP_1
	KeyKp2                = sdl.K_KP_2
	KeyKp3                = sdl.K_KP_3
	KeyKp4                = sdl.K_KP_4
	KeyKp5                = sdl.K_KP_5
	KeyKp6                = sdl.K_KP_6
	KeyKp7                = sdl.K_KP_7
	KeyKp8                = sdl.K_KP_8
	KeyKp9                = sdl.K_KP_9
	KeyKpA                = sdl.K_KP_A
	KeyKpAmpersand        = sdl.K_KP_AMPERSAND
	KeyKpAt               = sdl.K_KP_AT
	KeyKpB                = sdl.K_KP_B
	KeyKpBackspace        = sdl.K_KP_BACKSPACE
	KeyKpBinary           = sdl.K_KP_BINARY
	KeyKpC                = sdl.K_KP_C
	KeyKpClear            = sdl.K_KP_CLEAR
	KeyKpClearEntry       = sdl.K_KP_CLEARENTRY
	KeyKpColon            = sdl.K_KP_COLON
	KeyKpComma            = sdl.K_KP_COMMA
	KeyKpD                = sdl.K_KP_D
	KeyKpDbLAmpersand     = sdl.K_KP_DBLAMPERSAND
	KeyKpDbLVerticalBar   = sdl.K_KP_DBLVERTICALBAR
	KeyKpDecimal          = sdl.K_KP_DECIMAL
	KeyKpDivide           = sdl.K_KP_DIVIDE
	KeyKpE                = sdl.K_KP_E
	KeyKpEnter            = sdl.K_KP_ENTER
	KeyKpEquals           = sdl.K_KP_EQUALS
	KeyKpExclam           = sdl.K_KP_EXCLAM
	KeyKpF                = sdl.K_KP_F
	KeyKpGreater          = sdl.K_KP_GREATER
	KeyKpHash             = sdl.K_KP_HASH
	KeyKpHexadecimal      = sdl.K_KP_HEXADECIMAL
	KeyKpLeftBrace        = sdl.K_KP_LEFTBRACE
	KeyKpLeftParen        = sdl.K_KP_LEFTPAREN
	KeyKpLess             = sdl.K_KP_LESS
	KeyKpMemadd           = sdl.K_KP_MEMADD
	KeyKpMemclear         = sdl.K_KP_MEMCLEAR
	KeyKpMemdivide        = sdl.K_KP_MEMDIVIDE
	KeyKpMemmultiply      = sdl.K_KP_MEMMULTIPLY
	KeyKpMemrecall        = sdl.K_KP_MEMRECALL
	KeyKpMemstore         = sdl.K_KP_MEMSTORE
	KeyKpMemsubtract      = sdl.K_KP_MEMSUBTRACT
	KeyKpMinus            = sdl.K_KP_MINUS
	KeyKpMultiply         = sdl.K_KP_MULTIPLY
	KeyKpOctal            = sdl.K_KP_OCTAL
	KeyKpPercent          = sdl.K_KP_PERCENT
	KeyKpPeriod           = sdl.K_KP_PERIOD
	KeyKpPlus             = sdl.K_KP_PLUS
	KeyKpPlusminus        = sdl.K_KP_PLUSMINUS
	KeyKpPower            = sdl.K_KP_POWER
	KeyKpRightbrace       = sdl.K_KP_RIGHTBRACE
	KeyKpRightparen       = sdl.K_KP_RIGHTPAREN
	KeyKpSpace            = sdl.K_KP_SPACE
	KeyKpTab              = sdl.K_KP_TAB
	KeyKpVerticalbar      = sdl.K_KP_VERTICALBAR
	KeyKpXor              = sdl.K_KP_XOR
	KeyL                  = sdl.K_l
	KeyLalt               = sdl.K_LALT
	KeyLctrl              = sdl.K_LCTRL
	KeyLeft               = sdl.K_LEFT
	KeyLeftbracket        = sdl.K_LEFTBRACKET
	KeyLgui               = sdl.K_LGUI
	KeyLshift             = sdl.K_LSHIFT
	KeyM                  = sdl.K_m
	KeyMail               = sdl.K_MAIL
	KeyMediaselect        = sdl.K_MEDIASELECT
	KeyMenu               = sdl.K_MENU
	KeyMinus              = sdl.K_MINUS
	KeyMode               = sdl.K_MODE
	KeyMute               = sdl.K_MUTE
	KeyN                  = sdl.K_n
	KeyNumlockclear       = sdl.K_NUMLOCKCLEAR
	KeyO                  = sdl.K_o
	KeyOper               = sdl.K_OPER
	KeyOut                = sdl.K_OUT
	KeyP                  = sdl.K_p
	KeyPagedown           = sdl.K_PAGEDOWN
	KeyPageup             = sdl.K_PAGEUP
	KeyPaste              = sdl.K_PASTE
	KeyPause              = sdl.K_PAUSE
	KeyPeriod             = sdl.K_PERIOD
	KeyPower              = sdl.K_POWER
	KeyPrintscreen        = sdl.K_PRINTSCREEN
	KeyPrior              = sdl.K_PRIOR
	KeyQ                  = sdl.K_q
	KeyR                  = sdl.K_r
	KeyRalt               = sdl.K_RALT
	KeyRctrl              = sdl.K_RCTRL
	KeyReturn             = sdl.K_RETURN
	KeyReturn2            = sdl.K_RETURN2
	KeyRgui               = sdl.K_RGUI
	KeyRight              = sdl.K_RIGHT
	KeyRightbracket       = sdl.K_RIGHTBRACKET
	KeyRshift             = sdl.K_RSHIFT
	KeyS                  = sdl.K_s
	KeyScrolllock         = sdl.K_SCROLLLOCK
	KeySelect             = sdl.K_SELECT
	KeySemicolon          = sdl.K_SEMICOLON
	KeySeparator          = sdl.K_SEPARATOR
	KeySlash              = sdl.K_SLASH
	KeySleep              = sdl.K_SLEEP
	KeySpace              = sdl.K_SPACE
	KeyStop               = sdl.K_STOP
	KeySysreq             = sdl.K_SYSREQ
	KeyT                  = sdl.K_t
	KeyTab                = sdl.K_TAB
	KeyThousandsseparator = sdl.K_THOUSANDSSEPARATOR
	KeyU                  = sdl.K_u
	KeyUndo               = sdl.K_UNDO
	KeyUnknown            = sdl.K_UNKNOWN
	KeyUp                 = sdl.K_UP
	KeyV                  = sdl.K_v
	KeyVolumedown         = sdl.K_VOLUMEDOWN
	KeyVolumeup           = sdl.K_VOLUMEUP
	KeyW                  = sdl.K_w
	KeyWww                = sdl.K_WWW
	KeyX                  = sdl.K_x
	KeyY                  = sdl.K_y
	KeyZ                  = sdl.K_z
)

type (
	KeyboardPressType = uint8
	KeyboardKey       = sdl.Keycode

	EvKeyboard struct {
		PressType KeyboardPressType
		Key       KeyboardKey
	}

	// todo: codegen
	HandlerKeyboard func(keyboard EvKeyboard) error
)

// todo: codegen
func (d *Dispatcher) OnKeyBoard(h HandlerKeyboard) {
	d.registryHandler(typeKeyboard, func(e sdl.Event) error {
		evKeyboard, ok := e.(*sdl.KeyboardEvent)
		if !ok {
			panic(fmt.Sprintf("can`t handle `OnKeyboard` unexpected event type `%s`", reflect.TypeOf(e)))
		}

		return h(d.assembleKeyboard(evKeyboard))
	})
}

func (d *Dispatcher) assembleKeyboard(ev *sdl.KeyboardEvent) EvKeyboard {
	var pressType KeyboardPressType

	if ev.Type == sdl.KEYDOWN {
		pressType = KeyboardPressTypePressed
	} else {
		pressType = KeyboardPressTypeReleased
	}

	return EvKeyboard{
		PressType: pressType,
		Key:       ev.Keysym.Sym,
	}
}
