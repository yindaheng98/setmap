package setmap

import "testing"

func TestHalfIfaceSetMap(t *testing.T) {
	sm := NewHalfIfaceSetMap[string, int]()
	sm.Add("1", SimpleIfaceSetMapItem[int]{1})
	sm.Add("1", SimpleIfaceSetMapItem[int]{2})
	sm.Add("1", SimpleIfaceSetMapItem[int]{3})
	sm.Add("2", SimpleIfaceSetMapItem[int]{1})
	sm.Add("2", SimpleIfaceSetMapItem[int]{2})
	sm.Add("2", SimpleIfaceSetMapItem[int]{3})
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{1}))
	t.Log(sm.Exists("1", SimpleIfaceSetMapItem[int]{3}))
	t.Log(sm.GetSet("1"))
	t.Log(sm.GetSet("2"))
	sm.Remove("1", SimpleIfaceSetMapItem[int]{3})
	sm.Remove("2", SimpleIfaceSetMapItem[int]{1})
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{1}))
	t.Log(sm.Exists("1", SimpleIfaceSetMapItem[int]{3}))
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{2}))
	sm.Remove("2", SimpleIfaceSetMapItem[int]{1})
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{2}))
	sm.Remove("2", SimpleIfaceSetMapItem[int]{2})
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{2}))
	sm.Remove("2", SimpleIfaceSetMapItem[int]{3})
	t.Log(sm.GetSet("1"))
	t.Log(sm.GetSet("2"))
	sm.Add("2", SimpleIfaceSetMapItem[int]{1})
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{2}))
	sm.Add("2", SimpleIfaceSetMapItem[int]{2})
	t.Log(sm.Exists("2", SimpleIfaceSetMapItem[int]{2}))
	sm.Add("2", SimpleIfaceSetMapItem[int]{3})
}
