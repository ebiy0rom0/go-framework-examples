// Code generated by "stringer -type LogLevel logger.go"; DO NOT EDIT.

package log

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Info-1]
	_ = x[Warn-2]
	_ = x[Error-3]
	_ = x[Fatal-4]
	_ = x[Debug-5]
}

const _LogLevel_name = "InfoWarnErrorFatalDebug"

var _LogLevel_index = [...]uint8{0, 4, 8, 13, 18, 23}

func (i LogLevel) String() string {
	i -= 1
	if i < 0 || i >= LogLevel(len(_LogLevel_index)-1) {
		return "LogLevel(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _LogLevel_name[_LogLevel_index[i]:_LogLevel_index[i+1]]
}