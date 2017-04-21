// AUTOGENERATED FILE: easyjson marshaler/unmarshalers.

package api

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"regexp"
	"strings"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
	Reg *regexp.Regexp = regexp.MustCompile(`([<<])\S+([>>])`)
)

func easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi(in *jlexer.Lexer, out *PublicStashTabs) {

	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "next_change_id":
			out.NextChangeId = string(in.String())
		case "stashes":
			if in.IsNull() {
				in.Skip()
				out.Stashes = nil
			} else {
				in.Delim('[')
				if out.Stashes == nil {
					if !in.IsDelim(']') {
						out.Stashes = make([]Stash, 0, 1)
					} else {
						out.Stashes = []Stash{}
					}
				} else {
					out.Stashes = (out.Stashes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Stash
					easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi1(in, &v1)
					out.Stashes = append(out.Stashes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi(out *jwriter.Writer, in PublicStashTabs) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"next_change_id\":")
	out.String(string(in.NextChangeId))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stashes\":")
	if in.Stashes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Stashes {
			if v2 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi1(out, v3)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PublicStashTabs) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PublicStashTabs) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PublicStashTabs) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PublicStashTabs) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi(l, v)
}
func easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi1(in *jlexer.Lexer, out *Stash) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "accountName":
			out.AccountName = string(in.String())
		case "lastCharacterName":
			out.LastCharacterName = string(in.String())
		case "id":
			out.Id = string(in.String())
		case "stash":
			out.Label = string(in.String())
		case "stashType":
			out.Type = string(in.String())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]Item, 0, 1)
					} else {
						out.Items = []Item{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Item
					easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi2(in, &v4)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "public":
			out.IsPublic = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi1(out *jwriter.Writer, in Stash) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"accountName\":")
	out.String(string(in.AccountName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"lastCharacterName\":")
	out.String(string(in.LastCharacterName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stash\":")
	out.String(string(in.Label))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stashType\":")
	out.String(string(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"items\":")
	if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Items {
			if v5 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi2(out, v6)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"public\":")
	out.Bool(bool(in.IsPublic))
	out.RawByte('}')
}
func easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi2(in *jlexer.Lexer, out *Item) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = strings.ToLower(Reg.ReplaceAllString(string(in.String()), ""))
		case "typeLine":
			out.Type = strings.ToLower(string(in.String()))
		case "properties":
			if in.IsNull() {
				in.Skip()
				out.Properties = nil
			} else {
				in.Delim('[')
				if out.Properties == nil {
					if !in.IsDelim(']') {
						out.Properties = make([]ItemProperty, 0, 1)
					} else {
						out.Properties = []ItemProperty{}
					}
				} else {
					out.Properties = (out.Properties)[:0]
				}
				for !in.IsDelim(']') {
					var v7 ItemProperty
					easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi3(in, &v7)
					out.Properties = append(out.Properties, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "requirements":
			if in.IsNull() {
				in.Skip()
				out.Requirements = nil
			} else {
				in.Delim('[')
				if out.Requirements == nil {
					if !in.IsDelim(']') {
						out.Requirements = make([]ItemProperty, 0, 1)
					} else {
						out.Requirements = []ItemProperty{}
					}
				} else {
					out.Requirements = (out.Requirements)[:0]
				}
				for !in.IsDelim(']') {
					var v8 ItemProperty
					easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi3(in, &v8)
					out.Requirements = append(out.Requirements, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sockets":
			if in.IsNull() {
				in.Skip()
				out.Sockets = nil
			} else {
				in.Delim('[')
				if out.Sockets == nil {
					if !in.IsDelim(']') {
						out.Sockets = make([]Socket, 0, 2)
					} else {
						out.Sockets = []Socket{}
					}
				} else {
					out.Sockets = (out.Sockets)[:0]
				}
				for !in.IsDelim(']') {
					var v9 Socket
					easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi4(in, &v9)
					out.Sockets = append(out.Sockets, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "explicitMods":
			if in.IsNull() {
				in.Skip()
				out.ExplicitMods = nil
			} else {
				in.Delim('[')
				if out.ExplicitMods == nil {
					if !in.IsDelim(']') {
						out.ExplicitMods = make([]string, 0, 4)
					} else {
						out.ExplicitMods = []string{}
					}
				} else {
					out.ExplicitMods = (out.ExplicitMods)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.ExplicitMods = append(out.ExplicitMods, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "implicitMods":
			if in.IsNull() {
				in.Skip()
				out.ImplicitMods = nil
			} else {
				in.Delim('[')
				if out.ImplicitMods == nil {
					if !in.IsDelim(']') {
						out.ImplicitMods = make([]string, 0, 4)
					} else {
						out.ImplicitMods = []string{}
					}
				} else {
					out.ImplicitMods = (out.ImplicitMods)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.ImplicitMods = append(out.ImplicitMods, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "utilityMods":
			if in.IsNull() {
				in.Skip()
				out.UtilityMods = nil
			} else {
				in.Delim('[')
				if out.UtilityMods == nil {
					if !in.IsDelim(']') {
						out.UtilityMods = make([]string, 0, 4)
					} else {
						out.UtilityMods = []string{}
					}
				} else {
					out.UtilityMods = (out.UtilityMods)[:0]
				}
				for !in.IsDelim(']') {
					var v12 string
					v12 = string(in.String())
					out.UtilityMods = append(out.UtilityMods, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "enchantMods":
			if in.IsNull() {
				in.Skip()
				out.EnchantMods = nil
			} else {
				in.Delim('[')
				if out.EnchantMods == nil {
					if !in.IsDelim(']') {
						out.EnchantMods = make([]string, 0, 4)
					} else {
						out.EnchantMods = []string{}
					}
				} else {
					out.EnchantMods = (out.EnchantMods)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.EnchantMods = append(out.EnchantMods, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "craftedMods":
			if in.IsNull() {
				in.Skip()
				out.CraftedMods = nil
			} else {
				in.Delim('[')
				if out.CraftedMods == nil {
					if !in.IsDelim(']') {
						out.CraftedMods = make([]string, 0, 4)
					} else {
						out.CraftedMods = []string{}
					}
				} else {
					out.CraftedMods = (out.CraftedMods)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.CraftedMods = append(out.CraftedMods, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "cosmeticMods":
			if in.IsNull() {
				in.Skip()
				out.CosmeticMods = nil
			} else {
				in.Delim('[')
				if out.CosmeticMods == nil {
					if !in.IsDelim(']') {
						out.CosmeticMods = make([]string, 0, 4)
					} else {
						out.CosmeticMods = []string{}
					}
				} else {
					out.CosmeticMods = (out.CosmeticMods)[:0]
				}
				for !in.IsDelim(']') {
					var v15 string
					v15 = string(in.String())
					out.CosmeticMods = append(out.CosmeticMods, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "note":
			out.Note = string(in.String())
		case "verified":
			out.IsVerified = bool(in.Bool())
		case "w":
			out.Width = int(in.Int())
		case "h":
			out.Height = int(in.Int())
		case "ilvl":
			out.ItemLevel = int(in.Int())
		case "icon":
			out.Icon = string(in.String())
		case "league":
			out.League = strings.ToLower(string(in.String()))
		case "id":
			out.Id = string(in.String())
		case "identified":
			out.IsIdentified = bool(in.Bool())
		case "corrupted":
			out.IsCorrupted = bool(in.Bool())
		case "lockedToCharacter":
			out.IsLockedToCharacter = bool(in.Bool())
		case "support":
			out.IsSupport = bool(in.Bool())
		case "descrText":
			out.DescriptionText = string(in.String())
		case "secDescrText":
			out.SecondDescriptionText = string(in.String())
		case "flavourText":
			if in.IsNull() {
				in.Skip()
				out.FlavorText = nil
			} else {
				in.Delim('[')
				if out.FlavorText == nil {
					if !in.IsDelim(']') {
						out.FlavorText = make([]string, 0, 4)
					} else {
						out.FlavorText = []string{}
					}
				} else {
					out.FlavorText = (out.FlavorText)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.FlavorText = append(out.FlavorText, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "artFilename":
			out.ArtFilename = string(in.String())
		case "frameType":
			out.FrameType = FrameType(in.Int())
		case "stackSize":
			out.StackSize = int(in.Int())
		case "maxStackSize":
			out.MaxStackSize = int(in.Int())
		case "x":
			out.X = int(in.Int())
		case "y":
			out.Y = int(in.Int())
		case "inventoryId":
			out.InventoryId = string(in.String())
		case "socketedItems":
			if in.IsNull() {
				in.Skip()
				out.SocketedItems = nil
			} else {
				in.Delim('[')
				if out.SocketedItems == nil {
					if !in.IsDelim(']') {
						out.SocketedItems = make([]Item, 0, 1)
					} else {
						out.SocketedItems = []Item{}
					}
				} else {
					out.SocketedItems = (out.SocketedItems)[:0]
				}
				for !in.IsDelim(']') {
					var v17 Item
					easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi2(in, &v17)
					out.SocketedItems = append(out.SocketedItems, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "isRelic":
			out.IsRelic = bool(in.Bool())
		case "talismanTier":
			out.TalismanTier = int(in.Int())
		case "prophecyText":
			out.ProphecyText = string(in.String())
		case "prophecyDiffText":
			out.ProphecyDifficultyText = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi2(out *jwriter.Writer, in Item) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"typeLine\":")
	out.String(string(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"properties\":")
	if in.Properties == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v18, v19 := range in.Properties {
			if v18 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi3(out, v19)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"requirements\":")
	if in.Requirements == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in.Requirements {
			if v20 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi3(out, v21)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"sockets\":")
	if in.Sockets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v22, v23 := range in.Sockets {
			if v22 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi4(out, v23)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"explicitMods\":")
	if in.ExplicitMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v24, v25 := range in.ExplicitMods {
			if v24 > 0 {
				out.RawByte(',')
			}
			out.String(string(v25))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"implicitMods\":")
	if in.ImplicitMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v26, v27 := range in.ImplicitMods {
			if v26 > 0 {
				out.RawByte(',')
			}
			out.String(string(v27))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"utilityMods\":")
	if in.UtilityMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v28, v29 := range in.UtilityMods {
			if v28 > 0 {
				out.RawByte(',')
			}
			out.String(string(v29))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"enchantMods\":")
	if in.EnchantMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v30, v31 := range in.EnchantMods {
			if v30 > 0 {
				out.RawByte(',')
			}
			out.String(string(v31))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"craftedMods\":")
	if in.CraftedMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v32, v33 := range in.CraftedMods {
			if v32 > 0 {
				out.RawByte(',')
			}
			out.String(string(v33))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"cosmeticMods\":")
	if in.CosmeticMods == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v34, v35 := range in.CosmeticMods {
			if v34 > 0 {
				out.RawByte(',')
			}
			out.String(string(v35))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"note\":")
	out.String(string(in.Note))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"verified\":")
	out.Bool(bool(in.IsVerified))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"w\":")
	out.Int(int(in.Width))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"h\":")
	out.Int(int(in.Height))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ilvl\":")
	out.Int(int(in.ItemLevel))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"icon\":")
	out.String(string(in.Icon))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"league\":")
	out.String(string(in.League))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"identified\":")
	out.Bool(bool(in.IsIdentified))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"corrupted\":")
	out.Bool(bool(in.IsCorrupted))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"lockedToCharacter\":")
	out.Bool(bool(in.IsLockedToCharacter))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"support\":")
	out.Bool(bool(in.IsSupport))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"descrText\":")
	out.String(string(in.DescriptionText))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"secDescrText\":")
	out.String(string(in.SecondDescriptionText))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"flavourText\":")
	if in.FlavorText == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v36, v37 := range in.FlavorText {
			if v36 > 0 {
				out.RawByte(',')
			}
			out.String(string(v37))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"artFilename\":")
	out.String(string(in.ArtFilename))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"frameType\":")
	out.Int(int(in.FrameType))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stackSize\":")
	out.Int(int(in.StackSize))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxStackSize\":")
	out.Int(int(in.MaxStackSize))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"x\":")
	out.Int(int(in.X))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"y\":")
	out.Int(int(in.Y))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"inventoryId\":")
	out.String(string(in.InventoryId))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"socketedItems\":")
	if in.SocketedItems == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v38, v39 := range in.SocketedItems {
			if v38 > 0 {
				out.RawByte(',')
			}
			easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi2(out, v39)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isRelic\":")
	out.Bool(bool(in.IsRelic))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"talismanTier\":")
	out.Int(int(in.TalismanTier))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"prophecyText\":")
	out.String(string(in.ProphecyText))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"prophecyDiffText\":")
	out.String(string(in.ProphecyDifficultyText))
	out.RawByte('}')
}
func easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi4(in *jlexer.Lexer, out *Socket) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "group":
			out.GroupId = int(in.Int())
		case "attr":
			out.Attribute = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi4(out *jwriter.Writer, in Socket) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"group\":")
	out.Int(int(in.GroupId))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"attr\":")
	out.String(string(in.Attribute))
	out.RawByte('}')
}
func easyjson8cf7917eDecodeGithubComAntholordPoeIndexerApi3(in *jlexer.Lexer, out *ItemProperty) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]interface{}, 0, 4)
					} else {
						out.Values = []interface{}{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v40 interface{}
					if m, ok := v40.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v40.(json.Unmarshaler); ok {
						m.UnmarshalJSON(in.Raw())
					} else {
						v40 = in.Interface()
					}
					out.Values = append(out.Values, v40)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "displayMode":
			out.DisplayMode = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf7917eEncodeGithubComAntholordPoeIndexerApi3(out *jwriter.Writer, in ItemProperty) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"values\":")
	if in.Values == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v41, v42 := range in.Values {
			if v41 > 0 {
				out.RawByte(',')
			}
			if m, ok := v42.(easyjson.Marshaler); ok {
				m.MarshalEasyJSON(out)
			} else if m, ok := v42.(json.Marshaler); ok {
				out.Raw(m.MarshalJSON())
			} else {
				out.Raw(json.Marshal(v42))
			}
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"displayMode\":")
	out.Int(int(in.DisplayMode))
	out.RawByte('}')
}
