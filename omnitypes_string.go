// Code generated by "stringer -type=OmniTypes"; DO NOT EDIT.

package reflect

import "fmt"

const _OmniTypes_name = "TableEnumerationStructUnionResourceExternalResourceApplication"

var _OmniTypes_index = [...]uint8{0, 5, 16, 22, 27, 35, 51, 62}

func (i OmniTypes) String() string {
	if i >= OmniTypes(len(_OmniTypes_index)-1) {
		return fmt.Sprintf("OmniTypes(%d)", i)
	}
	return _OmniTypes_name[_OmniTypes_index[i]:_OmniTypes_index[i+1]]
}