// Copyright 2020 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enginetest

import (
	"math"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/parse"
)

var InsertQueries = []WriteQueryTest{
	{
		"INSERT INTO mytable (s, i) VALUES ('x', 999);",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT i FROM mytable WHERE s = 'x';",
		[]sql.Row{{int64(999)}},
	},
	{
		"INSERT INTO niltable (i, f) VALUES (10, 10.0), (12, 12.0);",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT i,f FROM niltable WHERE f IN (10.0, 12.0) ORDER BY f;",
		[]sql.Row{{int64(10), 10.0}, {int64(12), 12.0}},
	},
	{
		"INSERT INTO mytable SET s = 'x', i = 999;",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT i FROM mytable WHERE s = 'x';",
		[]sql.Row{{int64(999)}},
	},
	{
		"INSERT INTO mytable VALUES (999, 'x');",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT i FROM mytable WHERE s = 'x';",
		[]sql.Row{{int64(999)}},
	},
	{
		"INSERT INTO mytable SET i = 999, s = 'x';",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT i FROM mytable WHERE s = 'x';",
		[]sql.Row{{int64(999)}},
	},
	{
		`INSERT INTO typestable VALUES (
			999, 127, 32767, 2147483647, 9223372036854775807,
			255, 65535, 4294967295, 18446744073709551615,
			3.40282346638528859811704183484516925440e+38, 1.797693134862315708145274237317043567981e+308,
			'2037-04-05 12:51:36', '2231-11-07',
			'random text', true, '{"key":"value"}', 'blobdata'
			);`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM typestable WHERE id = 999;",
		[]sql.Row{{
			int64(999), int8(math.MaxInt8), int16(math.MaxInt16), int32(math.MaxInt32), int64(math.MaxInt64),
			uint8(math.MaxUint8), uint16(math.MaxUint16), uint32(math.MaxUint32), uint64(math.MaxUint64),
			float32(math.MaxFloat32), float64(math.MaxFloat64),
			sql.Timestamp.MustConvert("2037-04-05 12:51:36"), sql.Date.MustConvert("2231-11-07"),
			"random text", sql.True, ([]byte)(`{"key":"value"}`), "blobdata",
		}},
	},
	{
		`INSERT INTO typestable SET
			id = 999, i8 = 127, i16 = 32767, i32 = 2147483647, i64 = 9223372036854775807,
			u8 = 255, u16 = 65535, u32 = 4294967295, u64 = 18446744073709551615,
			f32 = 3.40282346638528859811704183484516925440e+38, f64 = 1.797693134862315708145274237317043567981e+308,
			ti = '2037-04-05 12:51:36', da = '2231-11-07',
			te = 'random text', bo = true, js = '{"key":"value"}', bl = 'blobdata'
			;`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM typestable WHERE id = 999;",
		[]sql.Row{{
			int64(999), int8(math.MaxInt8), int16(math.MaxInt16), int32(math.MaxInt32), int64(math.MaxInt64),
			uint8(math.MaxUint8), uint16(math.MaxUint16), uint32(math.MaxUint32), uint64(math.MaxUint64),
			float32(math.MaxFloat32), float64(math.MaxFloat64),
			sql.Timestamp.MustConvert("2037-04-05 12:51:36"), sql.Date.MustConvert("2231-11-07"),
			"random text", sql.True, ([]byte)(`{"key":"value"}`), "blobdata",
		}},
	},
	{
		`INSERT INTO typestable VALUES (
			999, -128, -32768, -2147483648, -9223372036854775808,
			0, 0, 0, 0,
			1.401298464324817070923729583289916131280e-45, 4.940656458412465441765687928682213723651e-324,
			'0000-00-00 00:00:00', '0000-00-00',
			'', false, '', ''
			);`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM typestable WHERE id = 999;",
		[]sql.Row{{
			int64(999), int8(-math.MaxInt8 - 1), int16(-math.MaxInt16 - 1), int32(-math.MaxInt32 - 1), int64(-math.MaxInt64 - 1),
			uint8(0), uint16(0), uint32(0), uint64(0),
			float32(math.SmallestNonzeroFloat32), float64(math.SmallestNonzeroFloat64),
			sql.Timestamp.Zero(), sql.Date.Zero(),
			"", sql.False, ([]byte)(`""`), "",
		}},
	},
	{
		`INSERT INTO typestable SET
			id = 999, i8 = -128, i16 = -32768, i32 = -2147483648, i64 = -9223372036854775808,
			u8 = 0, u16 = 0, u32 = 0, u64 = 0,
			f32 = 1.401298464324817070923729583289916131280e-45, f64 = 4.940656458412465441765687928682213723651e-324,
			ti = '0000-00-00 00:00:00', da = '0000-00-00',
			te = '', bo = false, js = '', bl = ''
			;`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM typestable WHERE id = 999;",
		[]sql.Row{{
			int64(999), int8(-math.MaxInt8 - 1), int16(-math.MaxInt16 - 1), int32(-math.MaxInt32 - 1), int64(-math.MaxInt64 - 1),
			uint8(0), uint16(0), uint32(0), uint64(0),
			float32(math.SmallestNonzeroFloat32), float64(math.SmallestNonzeroFloat64),
			sql.Timestamp.Zero(), sql.Date.Zero(),
			"", sql.False, ([]byte)(`""`), "",
		}},
	},
	{
		`INSERT INTO mytable (i,s) VALUES (10, 'NULL')`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM mytable WHERE i = 10;",
		[]sql.Row{{int64(10), "NULL"}},
	},
	{
		`INSERT INTO typestable VALUES (999, null, null, null, null, null, null, null, null,
			null, null, null, null, null, null, null, null);`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM typestable WHERE id = 999;",
		[]sql.Row{{int64(999), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
	},
	{
		`INSERT INTO typestable SET id=999, i8=null, i16=null, i32=null, i64=null, u8=null, u16=null, u32=null, u64=null,
			f32=null, f64=null, ti=null, da=null, te=null, bo=null, js=null, bl=null;`,
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM typestable WHERE id = 999;",
		[]sql.Row{{int64(999), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
	},
	{
		"INSERT INTO mytable SELECT i+100,s FROM mytable",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM mytable ORDER BY i",
		[]sql.Row{
			{int64(1), "first row"},
			{int64(2), "second row"},
			{int64(3), "third row"},
			{int64(101), "first row"},
			{int64(102), "second row"},
			{int64(103), "third row"},
		},
	},
	{
		"INSERT INTO emptytable SELECT * FROM mytable",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM emptytable ORDER BY i",
		[]sql.Row{
			{int64(1), "first row"},
			{int64(2), "second row"},
			{int64(3), "third row"},
		},
	},
	{
		"INSERT INTO emptytable SELECT * FROM mytable where mytable.i > 2",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM emptytable ORDER BY i",
		[]sql.Row{
			{int64(3), "third row"},
		},
	},
	{
		"INSERT INTO mytable (i,s) SELECT i+10, 'new' FROM mytable",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM mytable ORDER BY i",
		[]sql.Row{
			{int64(1), "first row"},
			{int64(2), "second row"},
			{int64(3), "third row"},
			{int64(11), "new"},
			{int64(12), "new"},
			{int64(13), "new"},
		},
	},
	{
		"INSERT INTO mytable SELECT i2+100, s2 FROM othertable",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM mytable ORDER BY i,s",
		[]sql.Row{
			{int64(1), "first row"},
			{int64(2), "second row"},
			{int64(3), "third row"},
			{int64(101), "third"},
			{int64(102), "second"},
			{int64(103), "first"},
		},
	},
	{
		"INSERT INTO emptytable (s,i) SELECT * FROM othertable",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM emptytable ORDER BY i,s",
		[]sql.Row{
			{int64(1), "third"},
			{int64(2), "second"},
			{int64(3), "first"},
		},
	},
	{
		"INSERT INTO emptytable (s,i) SELECT concat(m.s, o.s2), m.i FROM othertable o JOIN mytable m ON m.i=o.i2",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM emptytable ORDER BY i,s",
		[]sql.Row{
			{int64(1), "first rowthird"},
			{int64(2), "second rowsecond"},
			{int64(3), "third rowfirst"},
		},
	},
	{
		"INSERT INTO mytable (i,s) SELECT (i + 10.0) / 10.0 + 10 + i, concat(s, ' new') FROM mytable",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM mytable ORDER BY i, s",
		[]sql.Row{
			{int64(1), "first row"},
			{int64(2), "second row"},
			{int64(3), "third row"},
			{int64(12), "first row new"},
			{int64(13), "second row new"},
			{int64(14), "third row new"},
		},
	},
	{
		"INSERT INTO mytable (i,s) SELECT CHAR_LENGTH(s), concat('numrows: ', count(*)) from mytable group by 1",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT * FROM mytable ORDER BY i, s",
		[]sql.Row{
			{1, "first row"},
			{2, "second row"},
			{3, "third row"},
			{9, "numrows: 2"},
			{10, "numrows: 1"},
		},
	},
	// TODO: this doesn't match MySQL. MySQL requires giving an alias to the expression to use it in a HAVING clause,
	//  but that causes an error in our engine. Needs work
	{
		"INSERT INTO mytable (i,s) SELECT CHAR_LENGTH(s), concat('numrows: ', count(*)) from mytable group by 1 HAVING CHAR_LENGTH(s)  > 9",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM mytable ORDER BY i, s",
		[]sql.Row{
			{1, "first row"},
			{2, "second row"},
			{3, "third row"},
			{10, "numrows: 1"},
		},
	},
	{
		"INSERT INTO mytable (i,s) SELECT i * 2, concat(s,s) from mytable order by 1 desc limit 1",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM mytable ORDER BY i, s",
		[]sql.Row{
			{1, "first row"},
			{2, "second row"},
			{3, "third row"},
			{6, "third rowthird row"},
		},
	},
	{
		"INSERT INTO mytable (i,s) SELECT i + 3, concat(s,s) from mytable order by 1 desc",
		[]sql.Row{{sql.NewOkResult(3)}},
		"SELECT * FROM mytable ORDER BY i, s",
		[]sql.Row{
			{1, "first row"},
			{2, "second row"},
			{3, "third row"},
			{4, "first rowfirst row"},
			{5, "second rowsecond row"},
			{6, "third rowthird row"},
		},
	},
	{
		"INSERT INTO mytable (i,s) values (1, 'hello') ON DUPLICATE KEY UPDATE s='hello'",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT * FROM mytable WHERE i = 1",
		[]sql.Row{{int64(1), "hello"}},
	},
	{
		"INSERT INTO mytable (i,s) values (1, 'hello2') ON DUPLICATE KEY UPDATE s='hello3'",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT * FROM mytable WHERE i = 1",
		[]sql.Row{{int64(1), "hello3"}},
	},
	{
		"INSERT INTO mytable (i,s) values (1, 'hello') ON DUPLICATE KEY UPDATE i=10",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT * FROM mytable WHERE i = 10",
		[]sql.Row{{int64(10), "first row"}},
	},
	{
		"INSERT INTO mytable (i,s) values (1, 'hello2') ON DUPLICATE KEY UPDATE s='hello3'",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT * FROM mytable WHERE i = 1",
		[]sql.Row{{int64(1), "hello3"}},
	},
	{
		"INSERT INTO mytable (i,s) values (1, 'hello2'), (2, 'hello3'), (4, 'no conflict') ON DUPLICATE KEY UPDATE s='hello4'",
		[]sql.Row{{sql.NewOkResult(5)}},
		"SELECT * FROM mytable ORDER BY 1",
		[]sql.Row{
			{1, "hello4"},
			{2, "hello4"},
			{3, "third row"},
			{4, "no conflict"},
		},
	},
	{
		"INSERT INTO mytable (i,s) values (10, 'hello') ON DUPLICATE KEY UPDATE s='hello'",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM mytable ORDER BY 1",
		[]sql.Row{
			{1, "first row"},
			{2, "second row"},
			{3, "third row"},
			{10, "hello"},
		},
	},
	{
		"INSERT INTO auto_increment_tbl (c0) values (44)",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM auto_increment_tbl ORDER BY pk",
		[]sql.Row{
			{1, 11},
			{2, 22},
			{3, 33},
			{4, 44},
		},
	},
	{
		"INSERT INTO auto_increment_tbl (c0) values (44),(55)",
		[]sql.Row{{sql.NewOkResult(2)}},
		"SELECT * FROM auto_increment_tbl ORDER BY pk",
		[]sql.Row{
			{1, 11},
			{2, 22},
			{3, 33},
			{4, 44},
			{5, 55},
		},
	},
	{
		"INSERT INTO auto_increment_tbl values (NULL, 44)",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM auto_increment_tbl ORDER BY pk",
		[]sql.Row{
			{1, 11},
			{2, 22},
			{3, 33},
			{4, 44},
		},
	},
	{
		"INSERT INTO auto_increment_tbl values (0, 44)",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM auto_increment_tbl ORDER BY pk",
		[]sql.Row{
			{1, 11},
			{2, 22},
			{3, 33},
			{4, 44},
		},
	},
	{
		"INSERT INTO auto_increment_tbl values (5, 44)",
		[]sql.Row{{sql.NewOkResult(1)}},
		"SELECT * FROM auto_increment_tbl ORDER BY pk",
		[]sql.Row{
			{1, 11},
			{2, 22},
			{3, 33},
			{5, 44},
		},
	},
	{
		"INSERT INTO auto_increment_tbl values " +
			"(NULL, 44), (NULL, 55), (9, 99), (NULL, 110), (NULL, 121)",
		[]sql.Row{{sql.NewOkResult(5)}},
		"SELECT * FROM auto_increment_tbl ORDER BY pk",
		[]sql.Row{
			{1, 11},
			{2, 22},
			{3, 33},
			{4, 44},
			{5, 55},
			{9, 99},
			{10, 110},
			{11, 121},
		},
	},
}

var InsertScripts = []ScriptTest{
	{
		Name: "insert into sparse auto_increment table",
		SetUpScript: []string{
			"create table auto (pk int primary key auto_increment)",
			"insert into auto values (10), (20), (30)",
			"insert into auto values (NULL)",
			"insert into auto values (40)",
			"insert into auto values (0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{10}, {20}, {30}, {31}, {40}, {41},
				},
			},
		},
	},
	{
		Name: "auto increment table handles deletes",
		SetUpScript: []string{
			"create table auto (pk int primary key auto_increment)",
			"insert into auto values (10)",
			"delete from auto where pk = 10",
			"insert into auto values (NULL)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{11},
				},
			},
		},
	},
	{
		Name: "create auto_increment table with out-of-line primary key def",
		SetUpScript: []string{
			`create table auto (
				pk int auto_increment,
				c0 int,
				primary key(pk)
			);`,
			"insert into auto values (NULL,10), (NULL,20), (NULL,30)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1, 10}, {2, 20}, {3, 30},
				},
			},
		},
	},
	{
		Name: "alter auto_increment value",
		SetUpScript: []string{
			`create table auto (
				pk int auto_increment,
				c0 int,
				primary key(pk)
			);`,
			"insert into auto values (NULL,10), (NULL,20), (NULL,30)",
			"alter table auto auto_increment 9;",
			"insert into auto values (NULL,90)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1,10}, {2,20}, {3,30}, {9,90},
				},
			},
		},
	},
	{
		Name: "alter auto_increment value to float",
		SetUpScript: []string{
			`create table auto (
				pk int auto_increment,
				c0 int,
				primary key(pk)
			);`,
			"insert into auto values (NULL,10), (NULL,20), (NULL,30)",
			"alter table auto auto_increment = 19.9;",
			"insert into auto values (NULL,190)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1,10}, {2,20}, {3,30}, {19,190},
				},
			},
		},
	},
	{
		Name: "auto increment on tinyint",
		SetUpScript: []string{
			"create table auto (pk tinyint primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1},{10},{11},
				},
			},
		},
	},
	{
		Name: "auto increment on smallint",
		SetUpScript: []string{
			"create table auto (pk smallint primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1},{10},{11},
				},
			},
		},
	},
	{
		Name: "auto increment on mediumint",
		SetUpScript: []string{
			"create table auto (pk mediumint primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1},{10},{11},
				},
			},
		},
	},
	{
		Name: "auto increment on int",
		SetUpScript: []string{
			"create table auto (pk int primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1},{10},{11},
				},
			},
		},
	},
	{
		Name: "auto increment on bigint",
		SetUpScript: []string{
			"create table auto (pk bigint primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{1},{10},{11},
				},
			},
		},
	},
	{
		Name: "auto increment on tinyint unsigned",
		SetUpScript: []string{
			"create table auto (pk tinyint unsigned primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{uint64(1)},{uint64(10)},{uint64(11)},
				},
			},
		},
	},
	{
		Name: "auto increment on smallint unsigned",
		SetUpScript: []string{
			"create table auto (pk smallint unsigned primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{uint64(1)},{uint64(10)},{uint64(11)},
				},
			},
		},
	},
	{
		Name: "auto increment on mediumint unsigned",
		SetUpScript: []string{
			"create table auto (pk mediumint unsigned primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{uint64(1)},{uint64(10)},{uint64(11)},
				},
			},
		},
	},
	{
		Name: "auto increment on int unsigned",
		SetUpScript: []string{
			"create table auto (pk int unsigned primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{uint64(1)},{uint64(10)},{uint64(11)},
				},
			},
		},
	},
	{
		Name: "auto increment on bigint unsigned",
		SetUpScript: []string{
			"create table auto (pk bigint unsigned primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{uint64(1)},{uint64(10)},{uint64(11)},
				},
			},
		},
	},
	{
		Name: "auto increment on float",
		SetUpScript: []string{
			"create table auto (pk float primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{float64(1)},{float64(10)},{float64(11)},
				},
			},
		},
	},
	{
		Name: "auto increment on double",
		SetUpScript: []string{
			"create table auto (pk double primary key auto_increment)",
			"insert into auto values (NULL),(10),(0)",
		},
		Assertions: []ScriptTestAssertion{
			{
				Query: "select * from auto order by 1",
				Expected: []sql.Row{
					{float64(1)},{float64(10)},{float64(11)},
				},
			},
		},
	},
}

var InsertErrorTests = []GenericErrorQueryTest{
	{
		"too few values",
		"INSERT INTO mytable (s, i) VALUES ('x');",
	},
	{
		"too many values one column",
		"INSERT INTO mytable (s) VALUES ('x', 999);",
	},
	{
		"too many values two columns",
		"INSERT INTO mytable (i, s) VALUES (999, 'x', 'y');",
	},
	{
		"too few values no columns specified",
		"INSERT INTO mytable VALUES (999);",
	},
	{
		"too many values no columns specified",
		"INSERT INTO mytable VALUES (999, 'x', 'y');",
	},
	{
		"non-existent column values",
		"INSERT INTO mytable (i, s, z) VALUES (999, 'x', 999);",
	},
	{
		"non-existent column set",
		"INSERT INTO mytable SET i = 999, s = 'x', z = 999;",
	},
	{
		"duplicate column",
		"INSERT INTO mytable (i, s, s) VALUES (999, 'x', 'x');",
	},
	{
		"duplicate column set",
		"INSERT INTO mytable SET i = 999, s = 'y', s = 'y';",
	},
	{
		"null given to non-nullable",
		"INSERT INTO mytable (i, s) VALUES (null, 'y');",
	},
	{
		"incompatible types",
		"INSERT INTO mytable (i, s) select * FROM othertable",
	},
	{
		"column count mismatch in select",
		"INSERT INTO mytable (i) select * FROM othertable",
	},
	{
		"column count mismatch in select",
		"INSERT INTO mytable select s FROM othertable",
	},
	{
		"column count mismatch in join select",
		"INSERT INTO mytable (s,i) SELECT * FROM othertable o JOIN mytable m ON m.i=o.i2",
	},
	{
		"duplicate key",
		"INSERT INTO mytable (i,s) values (1, 'hello')",
	},
	{
		"duplicate keys",
		"INSERT INTO mytable SELECT * from mytable",
	},
	{
		"bad column in on duplicate key update clause",
		"INSERT INTO mytable values (10, 'b') ON DUPLICATE KEY UPDATE notExist = 1",
	},
}

var InsertErrorScripts = []ScriptTest{
	{
		Name:        "create table with non-pk auto_increment column",
		Query:       "create table bad (pk int primary key, c0 int auto_increment);",
		ExpectedErr: parse.ErrInvalidAutoIncCols,
	},
	{
		Name:        "create multiple auto_increment columns",
		Query:       "create table bad (pk1 int auto_increment, pk2 int auto_increment, primary key (pk1,pk2));",
		ExpectedErr: parse.ErrInvalidAutoIncCols,
	},
	{
		Name:        "create auto_increment column with default",
		Query:       "create table bad (pk1 int auto_increment default 10, c0 int);",
		ExpectedErr: parse.ErrInvalidAutoIncCols,
	},
}
