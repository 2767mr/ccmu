package pkg

import "fmt"

import "net/http"

import "errors"

import "encoding/json"

//ErrorReason that is included in Error.
type ErrorReason int

//Reasons for an installation, uninstallation or update to fail
const (
	ReasonUnknown ErrorReason = iota
	ReasonNoInternet
	ReasonAlreadyInstalled
	ReasonInvalidFormat
)

//Mode of the installation.
type Mode int

//Modes of operations
const (
	ModeInstall Mode = iota
	ModeUninstall
	ModeUpdate
)

//Error contains details about the errors that occured while installing or uninstalling a package.
type Error struct {
	Reason ErrorReason
	Mode   Mode
	Pkg    Package
	Err    error
}

//NewError with given mode and package.
func NewError(mode Mode, pkg Package, err error) Error {
	var reason = ReasonUnknown

	if errors.Is(err, &http.ProtocolError{}) {
		reason = ReasonNoInternet
	} else if errors.Is(err, &json.SyntaxError{}) {
		reason = ReasonInvalidFormat
	}

	return Error{reason, mode, pkg, err}
}

//NewErrorReason with given reason, mode and package.
func NewErrorReason(reason ErrorReason, mode Mode, pkg Package, err error) Error {
	return Error{reason, mode, pkg, err}
}

func (p Error) Error() string {
	return "pkg: " + p.String()
}

//Unwrap the error underneath.
func (p Error) Unwrap() error {
	return p.Err
}

func (p Error) String() string {
	info, _ := p.Pkg.Info()

	switch p.Reason {
	case ReasonUnknown:
		fallthrough
	default:
		return fmt.Sprintf("Could not " + p.Mode.String() + " " + info.NiceName + " because an unknown error occured")
	}
}

func (m Mode) String() string {
	switch m {
	case ModeInstall:
		return "install"
	case ModeUninstall:
		return "uninstall"
	case ModeUpdate:
		return "update"
	default:
		return "manipulate"
	}
}