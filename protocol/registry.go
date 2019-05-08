package protocol

import (
	"reflect"
	"sync"
)

// Version represents our version.
var (
	Version = ""

	usedRegistry *Registry
	// TODO We might not need the mutex, because we only really need to initiate the registry at the start, and not anywhere inbetween.
	usedRegistryLock = &sync.Mutex{}
)

// GetPacketType will return the packet, if it exists.
func GetPacketType(direction Direction, state State, id int) (reflect.Type, error) {
	usedRegistryLock.Lock()
	defer usedRegistryLock.Unlock()

	if usedRegistry != nil {
		if _, ok := usedRegistry.packets[direction]; ok {
			if _, ok := usedRegistry.packets[direction][state]; ok {
				if typ, ok := usedRegistry.packets[direction][state][id]; ok {
					return typ, nil
				}
			}
		}
	}

	return nil, ErrUnknownPacketType
}

// Registry will help manage the registration of packets.
type Registry struct {
	checker func(string) bool
	packets map[Direction]map[State]map[int]reflect.Type
}

func (registry *Registry) validatePacketMap(direction Direction, state State) {
	if registry.packets == nil {
		registry.packets = make(map[Direction]map[State]map[int]reflect.Type)
	}
	if _, ok := registry.packets[direction]; !ok {
		registry.packets[direction] = make(map[State]map[int]reflect.Type)
	}
	if _, ok := registry.packets[direction][state]; !ok {
		registry.packets[direction][state] = make(map[int]reflect.Type)
	}
}

func (registry *Registry) validatePacketType(packet reflect.Type) bool {
	// TODO verify the packet type.
	return true
}

// RegisterPacket will register the packet to the registry
func (registry *Registry) RegisterPacket(direction Direction, state State, id int, packet reflect.Type) {
	registry.validatePacketMap(direction, state)
	registry.packets[direction][state][id] = packet
}

// Submit will submit the registry for the server to use.
func (registry *Registry) Submit() {
	usedRegistryLock.Lock()
	defer usedRegistryLock.Unlock()

	if registry.checker(Version) {
		// If the registry is not nil, then a packet list is already selected and in use.
		if usedRegistry == nil {
			usedRegistry = registry
		}
	}
}

// NewRegistry will create a new registry to work with.
func NewRegistry(checker func(version string) (coverage bool)) *Registry {
	return &Registry{checker: checker}
}
