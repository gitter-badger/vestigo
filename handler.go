package vestigo

import "net/http"

// Resource - internal structure for specifying which handlers belong to a particular route
type Resource struct {
	Cors           *CorsAccessControl
	Connect        http.HandlerFunc
	Delete         http.HandlerFunc
	Get            http.HandlerFunc
	Head           http.HandlerFunc
	Patch          http.HandlerFunc
	Post           http.HandlerFunc
	Put            http.HandlerFunc
	Trace          http.HandlerFunc
	allowedMethods string
}

func NewResource() *Resource {
	r := new(Resource)
	r.Cors = new(CorsAccessControl)
	return r
}

// CopyTo - Copy the Resource to another Resource passed in by reference
func (h *Resource) CopyTo(v *Resource) {
	*v.Cors = *h.Cors
	v.Get = h.Get
	v.Connect = h.Connect
	v.Delete = h.Delete
	v.Get = h.Get
	v.Head = h.Head
	v.Patch = h.Patch
	v.Post = h.Post
	v.Put = h.Put
	v.Trace = h.Trace
	v.allowedMethods = h.allowedMethods
}

// addToAllowedMethods - Add a method to the allowed methods for this route
func (h *Resource) addToAllowedMethods(method string) {
	if h.allowedMethods == "" {
		h.allowedMethods = method
	} else {
		h.allowedMethods = h.allowedMethods + ", " + method
	}
}

// AddMethodResource - Add a method/Resource pair to the Resource structure
func (h *Resource) AddMethodHandler(method string, handler http.HandlerFunc) {
	l := len(method)
	firstChar := method[0]
	secondChar := method[1]
	if h != nil {
		if l == 3 {
			if uint16(firstChar)<<8|uint16(secondChar) == 0x4745 {
				h.addToAllowedMethods(method)
				h.Get = handler
			}
			if uint16(firstChar)<<8|uint16(secondChar) == 0x5055 {
				h.addToAllowedMethods(method)
				h.Put = handler
			}
		} else if l == 4 {
			if uint16(firstChar)<<8|uint16(secondChar) == 0x504f {
				h.addToAllowedMethods(method)
				h.Post = handler
			}
			if uint16(firstChar)<<8|uint16(secondChar) == 0x4845 {
				h.addToAllowedMethods(method)
				h.Head = handler
			}
		} else if l == 5 {
			if uint16(firstChar)<<8|uint16(secondChar) == 0x5452 {
				h.addToAllowedMethods(method)
				h.Trace = handler
			}
			if uint16(firstChar)<<8|uint16(secondChar) == 0x5041 {
				h.addToAllowedMethods(method)
				h.Patch = handler
			}
		} else if l >= 6 {
			if uint16(firstChar)<<8|uint16(secondChar) == 0x4445 {
				h.addToAllowedMethods(method)
				h.Delete = handler
			}
			if uint16(firstChar)<<8|uint16(secondChar) == 0x434f {
				h.addToAllowedMethods(method)
				h.Connect = handler
			}
		}
	}
}

// GetMethodResource - Get a method/Resource pair from the Resource structure
func (h *Resource) GetMethodHandler(method string) (http.HandlerFunc, string) {
	l := len(method)
	firstChar := method[0]
	secondChar := method[1]
	if l == 3 {
		if uint16(firstChar)<<8|uint16(secondChar) == 0x4745 {
			return h.Get, h.allowedMethods
		}
		if uint16(firstChar)<<8|uint16(secondChar) == 0x5055 {
			return h.Put, h.allowedMethods
		}
	} else if l == 4 {
		if uint16(firstChar)<<8|uint16(secondChar) == 0x504f {
			return h.Post, h.allowedMethods
		}
		if uint16(firstChar)<<8|uint16(secondChar) == 0x4845 {
			return h.Head, h.allowedMethods
		}
	} else if l == 5 {
		if uint16(firstChar)<<8|uint16(secondChar) == 0x5452 {
			return h.Trace, h.allowedMethods
		}
		if uint16(firstChar)<<8|uint16(secondChar) == 0x5041 {
			return h.Patch, h.allowedMethods
		}
	} else if l >= 6 {
		if uint16(firstChar)<<8|uint16(secondChar) == 0x4445 {
			return h.Delete, h.allowedMethods
		}
		if uint16(firstChar)<<8|uint16(secondChar) == 0x434f {
			return h.Connect, h.allowedMethods
		}
	}
	return nil, h.allowedMethods
}
