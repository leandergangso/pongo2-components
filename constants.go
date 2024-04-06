package pongo2components

// Supported tag names.
const (
	componentTagName    = "component"
	endComponentTagName = "endcomponent"
	slotTagName         = "slot"
	endSlotTagName      = "endslot"
)

// Context keys.
const (
	componentCtxKey = "_component"
)

// Helper for commonly used slot names.
const (
	SlotBefore  = "slotBefore"
	SlotDefault = "slotDefault"
	SlotAfter   = "slotAfter"
)
