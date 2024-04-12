// Copyright 2024 Kozmoai Authors.
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

package parser_sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipIgnored(t *testing.T) {
	sql_1 := `
	/* select syntax for client query */
	SELECT * FROM tab1 where id=12;
	`
	lexer := NewLexer(sql_1)
	lexer.skipIgnored()

	_, _, token, err := lexer.GetNextToken()
	assert.Nil(t, err)

	assert.Equal(t, tokenNameMap[TOKEN_SELECT], token, "the token should be equal")
}

func TestSkipIgnored2(t *testing.T) {

	sql_2 := `
	# comment
	select * FROM tab1 where id=12; -- comment
	`

	lexer := NewLexer(sql_2)
	lexer.skipIgnored()

	_, _, token, err := lexer.GetNextToken()
	assert.Nil(t, err)

	assert.Equal(t, tokenNameMap[TOKEN_SELECT], token, "the token should be equal")
}

func TestSkipIgnored3(t *testing.T) {

	sql_3 := `
	# comment /* also comment */
	SelecT * FROM tab1 where id=12; -- comment
	`

	lexer := NewLexer(sql_3)
	lexer.skipIgnored()

	_, _, token, err := lexer.GetNextToken()
	assert.Nil(t, err)

	assert.Equal(t, tokenNameMap[TOKEN_SELECT], token, "the token should be equal")
}

func TestSkipIgnored4(t *testing.T) {

	sql_4 := `
	/* comment  */
	DELETE  FROM tab1 where id=12; -- comment
	`

	lexer := NewLexer(sql_4)
	lexer.skipIgnored()

	_, _, token, err := lexer.GetNextToken()
	assert.Nil(t, err)

	assert.Equal(t, tokenNameMap[TOKEN_DELETE], token, "the token should be equal")
}

func TestNextTokenIs1(t *testing.T) {
	sql_1 := `
	/* comment  */
	DELETE  FROM tab1 where id=12; -- comment
	`

	lexer := NewLexer(sql_1)
	lexer.skipIgnored()

	_, _, err := lexer.NextTokenIs(TOKEN_DELETE)
	assert.Nil(t, err)
}

func TestNextTokenIs2(t *testing.T) {
	sql_2 := `
	/* comment  */
	DELETE  FROM tab1 where id=12; -- comment
	`

	lexer := NewLexer(sql_2)
	lexer.skipIgnored()

	_, _, err := lexer.NextTokenIs(TOKEN_UPDATE)
	assert.NotNil(t, err)
}

func TestLookAhead(t *testing.T) {
	sql_1 := `
	/* select syntax for client query */
	SELECT id FROM tab1 where id=12;
	`
	lexer := NewLexer(sql_1)
	lexer.skipIgnored()
	token, err := lexer.LookAhead()

	assert.Nil(t, err)
	assert.Equal(t, TOKEN_SELECT, token, "it should be a TOKEN_SELECT")
}

func TestNextSQLIs(t *testing.T) {
	sql_1 := `
	/* select syntax for client query */
	SELECT id FROM tab1 where id=12;
	`
	lexer := NewLexer(sql_1)
	lexer.skipIgnored()
	doestItIs := lexer.nextSQLIs("SELECT")

	assert.Equal(t, true, doestItIs, "it should be a SELECT token")
}

func TestGetNextToken(t *testing.T) {
	sql_1 := `
	/* select syntax for client query */
	SELECT id FROM tab1 where id=12;
	`
	lexer := NewLexer(sql_1)
	lexer.skipIgnored()
	_, _, token, err := lexer.GetNextToken()

	assert.Nil(t, err)
	assert.Equal(t, "select", token, "it should be a select token")
}
