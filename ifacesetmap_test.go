package setmap

import "testing"

func TestIfaceSetMap(t *testing.T) {
	sm := NewIfaceSetMap[string, int]()
	sm.Add(SimpleIfaceSetMapItem[string]{"1"}, SimpleIfaceSetMapItem[int]{1})
	sm.Add(SimpleIfaceSetMapItem[string]{"1"}, SimpleIfaceSetMapItem[int]{2})
	sm.Add(SimpleIfaceSetMapItem[string]{"1"}, SimpleIfaceSetMapItem[int]{3})
	sm.Add(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{1})
	sm.Add(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2})
	sm.Add(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{3})
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{1}))
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"1"}, SimpleIfaceSetMapItem[int]{3}))
	t.Log(sm.GetSet(SimpleIfaceSetMapItem[string]{"1"}))
	t.Log(sm.GetSet(SimpleIfaceSetMapItem[string]{"2"}))
	sm.Remove(SimpleIfaceSetMapItem[string]{"1"}, SimpleIfaceSetMapItem[int]{3})
	sm.Remove(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{1})
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{1}))
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"1"}, SimpleIfaceSetMapItem[int]{3}))
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2}))
	sm.Remove(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{1})
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2}))
	sm.Remove(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2})
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2}))
	sm.Remove(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{3})
	t.Log(sm.GetSet(SimpleIfaceSetMapItem[string]{"1"}))
	t.Log(sm.GetSet(SimpleIfaceSetMapItem[string]{"2"}))
	sm.Add(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{1})
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2}))
	sm.Add(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2})
	t.Log(sm.Exists(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{2}))
	sm.Add(SimpleIfaceSetMapItem[string]{"2"}, SimpleIfaceSetMapItem[int]{3})
}
