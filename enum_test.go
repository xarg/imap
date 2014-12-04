package imap

import "testing"

func TestEnumString(t *testing.T) {
	type dataIn struct {
		v        uint32
		names    []enumName
		goSyntax bool
	}

	type dataTest struct {
		title string
		in    dataIn
		out   string
	}
	d := []dataTest{
		dataTest{
			"A",
			dataIn{
				1,
				[]enumName{
					enumName{1, "gmail"},
					enumName{2, "yahoo"},
				},
				true,
			},
			"imap.gmail",
		},
		dataTest{
			"B",
			dataIn{
				1,
				[]enumName{
					enumName{1, "gmail"},
					enumName{2, "yahoo"},
				},
				false,
			},
			"gmail",
		},
		dataTest{
			"C",
			dataIn{
				2,
				[]enumName{
					enumName{1, "gmail"},
					enumName{2, "yahoo"},
				},
				false,
			},
			"yahoo",
		},
		dataTest{
			"D",
			dataIn{
				1,
				[]enumName{
					enumName{1, ""},
					enumName{2, "yahoo"},
				},
				false,
			},
			"",
		},
		dataTest{
			"E",
			dataIn{
				3,
				[]enumName{
					enumName{1, "gmail"},
					enumName{2, "yahoo"},
				},
				false,
			},
			"gmail+yahoo",
		},
		dataTest{
			"F",
			dataIn{
				3,
				[]enumName{
					enumName{1, "gmail"},
					enumName{2, "yahoo"},
				},
				true,
			},
			"imap.gmail+imap.yahoo",
		},
		dataTest{
			"G",
			dataIn{
				3,
				[]enumName{},
				true,
			},
			"0x3",
		},
		dataTest{
			"H",
			dataIn{
				2,
				[]enumName{enumName{1, ""}},
				true,
			},
			"0x2",
		},
		dataTest{
			"I",
			dataIn{
				4,
				[]enumName{
					enumName{0, "gmail"},
					enumName{1, "yahoo"},
					enumName{2, "kjds"},
					enumName{3, "6sqd54f"},
				},
				true,
			},
			"0x4",
		},
	}
	for i := 0; i < len(d); i++ {
		dataTitle := d[i].title
		dataIn := d[i].in
		dataOut := d[i].out
		res := enumString(dataIn.v, dataIn.names, dataIn.goSyntax)
		if res != dataOut {
			t.Error(
				"For", dataTitle,
				"Expected", dataOut,
				"Got", res,
			)
		}
	}
}
