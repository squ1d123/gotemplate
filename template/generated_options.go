// Code generated by "stringer -type=Options -output generated_options.go"; DO NOT EDIT.

package template

import "strconv"

const _Options_name = "RazorExtensionMathSprigDataLoggingRuntimeUtilsNetOptionOnByDefaultCountOverwriteOutputStdoutRenderingDisabled"

var _Options_index = [...]uint8{0, 5, 14, 18, 23, 27, 34, 41, 46, 49, 71, 80, 92, 109}

func (i Options) String() string {
	if i < 0 || i >= Options(len(_Options_index)-1) {
		return "Options(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Options_name[_Options_index[i]:_Options_index[i+1]]
}
